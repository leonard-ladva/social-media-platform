package request

import (
	"fmt"
	"forum-test/database"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func NewPost(w http.ResponseWriter, r *http.Request, user database.User) {
	tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))
	var pagedata database.PageData
	pagedata.User = user
	var err error
	postid := -1
	keys, ok := r.URL.Query()["id"]
	if ok {
		postid, err = strconv.Atoi(keys[0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.ExecuteTemplate(w, "err500", pagedata)
			log.Println("Cannot convert post id")
			return
		}
		pagedata.Post, err = database.GetPostByPostAndUsedID(postid, user.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.ExecuteTemplate(w, "err500", pagedata)
			log.Println("Cannot get post with id")
			return
		}
		pagedata.Categories, err = database.GetCategoriesForPostEdit(postid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.ExecuteTemplate(w, "err500", pagedata)
			log.Println("Cannot get categories")
			return
		}
	} else {
		pagedata.Categories, err = database.GetCategories()
	}

	if postid > 0 && pagedata.User.Username != pagedata.Post.Author && pagedata.User.RoleID != 3 {
		w.WriteHeader(http.StatusUnauthorized)
		tmpl.ExecuteTemplate(w, "err401", pagedata)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "err500", pagedata)
		log.Println("Cannot get categories from DB - category.go")
		return
	}

	if user.RoleID < 1 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	switch r.Method {

	case "GET":

		tmpl.ExecuteTemplate(w, "new-post", pagedata)

	case "POST":
		r.ParseMultipartForm(21 << 20)
		r.ParseForm()

		title := r.FormValue("title")
		content := r.FormValue("content")
		category := r.Form["category"]

		if category == nil || title == "" || content == "" {
			pagedata.Message.Msg1 = "Please insert post title, write the content and select at least one category for your post"
			pagedata.Message.Msg2 = title
			pagedata.Message.Msg3 = content
			tmpl.ExecuteTemplate(w, "new-post", pagedata)
			return
		}

		//check image size if attached
		file, handler, _ := r.FormFile("image")
		if file != nil {

			//check for uploaded file size
			if handler.Size > 10<<21 {
				w.WriteHeader(http.StatusRequestEntityTooLarge)
				fmt.Fprintf(w, "Upload size limit exceeded! Please keep your upload below 20mb")
				break
			}

		}

		if postid < 0 { //if there was no post_id submitted as parameter (adding new post)
			//add the post to database
			statement, err := database.Db.Prepare("INSERT INTO post ( title, content, user_id, status, datetime) VALUES ( ?, ?, ?, ?, ?);")
			if err != nil {
				log.Println("Database error, cannot insert new post.")
			}
			defer statement.Close()
			var status string
			if pagedata.User.RoleID >= 2 {
				status = "approved" //posts added by moderators and admins are approved automatically
			} else {
				status = "review" //posts added by users should be reviewed by moderators
			}
			statement.Exec(title, content, user.ID, status, time.Now())

			//get the post_id of the post just added to the database
			rows, err := database.Db.Query(`SELECT post_id FROM post WHERE title = ? AND content = ? AND user_id = ? ;`, title, content, user.ID)
			if err != nil {
				log.Println("ERROR | No postid found")
				log.Println(err)
				return
			}
			defer rows.Close()
			for rows.Next() {
				rows.Scan(&postid)
			}

			if status == "review" { //if post was submitted by user, notify moderator about new post which needs to be reviewed
				notification := "Please review new post submitted by " + pagedata.User.Username + ":"
				link := "/post?id=" + strconv.Itoa(postid)
				_ = database.AddNotificationForRole(notification, link, 7, "moderator")
			}

		} else { //if post_id was sumitted as parameter (editing existing post)
			//update the post in database and delete all rows from post_category table
			statement, _ := database.Db.Prepare("UPDATE post SET title = ?, content = ? WHERE post_id = ? ;")
			defer statement.Close()
			t := time.Now()
			content = content + "\n\nEdited by " + user.Username + " on " + t.Format(time.Stamp)
			statement.Exec(title, content, postid)
			statement, _ = database.Db.Prepare("DELETE FROM post_category WHERE post_id = ? ; ")
			defer statement.Close()
			statement.Exec(postid)

			//If post was edited by admin, notify post author about it
			if pagedata.User.Username != pagedata.Post.Author {
				notification := `Your post was edited by ` + pagedata.User.Username + ". Check it out: "
				link := "/post?id=" + strconv.Itoa(postid)
				usr := database.GetUserByUserName(pagedata.Post.Author)
				_ = database.AddNotificationForUser(notification, link, 5, usr.ID)
			}
		}

		if postid == -1 {
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.ExecuteTemplate(w, "err500", pagedata)
			log.Println("Error while processing new or edited post")
			return
		}

		//add the post categories to database
		for _, v := range category {

			statement, err := database.Db.Prepare("INSERT INTO post_category (post_id, category_id) VALUES ( ?, ?);")
			if err != nil {
				log.Println("Database error, cannot insert new post_category relation.")
				defer statement.Close()
			}
			statement.Exec(postid, v)
		}
		//get image data and add to database
		if file != nil {
			defer file.Close()
			fmt.Printf("Uploaded File: %+v\n", handler.Filename)
			fmt.Printf("File Size: %+v\n", handler.Size)
			fmt.Printf("MIME Header: %+v\n", handler.Header)

			//check for uploaded file size
			if handler.Size > 10<<21 {
				w.WriteHeader(http.StatusRequestEntityTooLarge)
				fmt.Fprintf(w, "Upload size limit exceeded! Please keep your upload below 20mb")
				break
			}
			//check if there is already a folder by username. if not, create it
			if _, err := os.Stat("public/upload/" + user.Username); os.IsNotExist(err) {
				err := os.Mkdir("public/upload/"+user.Username, os.ModePerm)
				if err != nil {
					log.Println(err)
				}
			}

			//create file in the folder, add postid to file name
			dst, err := os.Create("public/upload/" + user.Username + "/" + strconv.Itoa(postid) + "-" + handler.Filename)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			// Copy the uploaded file to the created file on the filesystem
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			statement, err := database.Db.Prepare("UPDATE post SET postimg = ? WHERE post_id = ? ;")
			if err != nil {
				log.Println(err)
			}
			defer statement.Close()
			imgurl := user.Username + "/" + strconv.Itoa(postid) + "-" + handler.Filename
			statement.Exec(imgurl, postid)

		}

		//redirect the user to the post page they just made. is it logical? maybe direct to all posts?
		redirect := strconv.Itoa(postid)

		http.Redirect(w, r, "post?id="+redirect, http.StatusSeeOther)

	}

}
