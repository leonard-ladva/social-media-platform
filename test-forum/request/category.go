package request

import (
	"forum-test/database"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Category(w http.ResponseWriter, r *http.Request, user database.User) {
	tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))
	user = database.GetUserByCookie(w, r)
	var pagedata database.PageData
	pagedata.User = user
	var err error
	pagedata.Categories, err = database.GetCategories()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "err500", pagedata)
		log.Println("Cannot get categories from DB - category.go -> getCategories.go")
		return
	}

	switch r.Method {

	case "GET":
		if r.URL.Path == "/category/all" {
			pagedata.Posts, err = database.GetAllPosts()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				log.Println("Cannot get posts from DB - category.go -> getPosts.go")
				return
			}

		} else {

			category_id, _ := database.GetCategoryIdByTitle(r.URL.Path[10:])

			pagedata.Posts, err = database.GetPostsByCategoryId(category_id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
			if category_id < 1 {
				w.WriteHeader(http.StatusNotFound)
				tmpl.ExecuteTemplate(w, "err404", pagedata)
				return
			}
		}

		for i := range pagedata.Posts {
			if len(pagedata.Posts[i].Content) > 75 {
				pagedata.Posts[i].Content = pagedata.Posts[i].Content[0:75] + "..."
			}
		}
		tmpl.ExecuteTemplate(w, "category", pagedata)
		// limit posts to 10 per page, add pagination. fetch posts by last activity (comment)?

	case "POST":

		if user.RoleID < 1 {
			//redirect to login page?
			return
		}

		log.Println("POST")
		r.ParseForm()

		userid := user.ID
		postid, _ := strconv.Atoi(r.FormValue("postid"))
		action := r.FormValue("action")
		log.Println("action:", action, "userid:", userid, "postid:", postid)

		currentreaction := database.GetReactionByPostAndUserID(postid, userid)

		//insert new record if there is no current reaction
		if currentreaction == "" && action != "" {

			statement, err := database.Db.Prepare("INSERT INTO post_reaction (type, user_id, post_id) VALUES (?,?,?);")
			if err != nil {
				log.Println(err)
			}
			defer statement.Close()

			statement.Exec(action, userid, postid)

			//delete record if current reaction and new reaction are the same
		} else if currentreaction == action {
			statement, err := database.Db.Prepare("DELETE FROM post_reaction WHERE user_id = ? AND post_id = ? ;")
			if err != nil {
				log.Println(err)
			}
			defer statement.Close()

			statement.Exec(userid, postid)

			//update post reaction if current and new reaction are different
		} else {

			statement, err := database.Db.Prepare("UPDATE post_reaction SET type = ? WHERE user_id = ? AND post_id = ? ;")
			if err != nil {
				log.Println(err)
			}
			defer statement.Close()

			statement.Exec(action, userid, postid)
		}

		http.Redirect(w, r, r.RequestURI, http.StatusSeeOther)
	}

}
