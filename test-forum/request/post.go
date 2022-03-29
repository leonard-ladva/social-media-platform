package request

import (
	"forum-test/database"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Post(w http.ResponseWriter, r *http.Request, user database.User) {

	tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))

	var pagedata database.PageData

	pagedata.User = user

	keys, ok := r.URL.Query()["id"]
	var err error
	if !ok || len(keys[0]) < 1 {
		//write bad request header and execute template
		w.WriteHeader(http.StatusBadRequest)
		tmpl.ExecuteTemplate(w, "err400", pagedata)
		return
	}

	// Query()["key"] will return an array of items, we only want the single item.
	postid, err := strconv.Atoi(keys[0])
	if err != nil {
		//bad request
		w.WriteHeader(http.StatusBadRequest)
		//execute bad request template
		tmpl.ExecuteTemplate(w, "err400", pagedata)
		return
	}

	pagedata.Post, err = database.GetPostByPostAndUsedID(postid, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "err500", pagedata)
		return
	}

	if postid != pagedata.Post.ID {
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "err404", pagedata)
		return
	}

	pagedata.Comments, err = database.GetCommentsByPostAndUserId(postid, user.ID)
	if err != nil {
		log.Println(err)
		//write internal server error header and template
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "err500", pagedata)
		return
	}

	switch r.Method {
	case "GET":
		tmpl.ExecuteTemplate(w, "view-post", pagedata)
		return

	case "POST":
		if user.RoleID < 1 {
			w.WriteHeader(http.StatusUnauthorized)
			tmpl.ExecuteTemplate(w, "err401", pagedata)
			return
		}

		r.ParseForm()
		postid := pagedata.Post.ID
		userid := pagedata.User.ID

		content := r.FormValue("comment")
		buttonpressed := r.FormValue("submit")
		postreaction := r.FormValue("postreaction")
		commentreaction := r.FormValue("commentreaction")
		currentpostreaction := database.GetReactionByPostAndUserID(postid, userid)

		//if user has submitted a comment reaction, do comment reaction stuff here
		if commentreaction != "" {
			commentid, err := strconv.Atoi(r.FormValue("commentid"))
			if err != nil {
				log.Println(err)
				//write internal server error header and template
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
			currentcommentreaction := database.GetReactionByCommentAndUserID(commentid, userid)
			//do comment reaction stuff here

			//insert new record if there is no current reaction
			if currentcommentreaction == "" && commentreaction != "" {
				statement, err := database.Db.Prepare("INSERT INTO comment_reaction (type, user_id, comment_id) VALUES (?,?,?);")
				if err != nil {
					log.Println(err)
				}
				defer statement.Close()
				statement.Exec(commentreaction, userid, commentid)

				//delete record if current reaction and new reaction are the same
			} else if currentcommentreaction == commentreaction {
				statement, err := database.Db.Prepare("DELETE FROM comment_reaction WHERE user_id = ? AND comment_id = ? ;")
				if err != nil {
					log.Println(err)
				}
				defer statement.Close()
				statement.Exec(userid, commentid)

				//update record if current and new reaction are different
			} else {
				statement, err := database.Db.Prepare("UPDATE comment_reaction SET type = ? WHERE user_id = ? AND comment_id = ? ;")
				if err != nil {
					log.Println(err)
				}
				defer statement.Close()
				statement.Exec(commentreaction, userid, commentid)
			}
		}

		//if user has submitted a post reaction, do post reaction stuff here
		if postreaction != "" {
			//insert new record if there is no current reaction
			if currentpostreaction == "" && postreaction != "" {

				statement, err := database.Db.Prepare("INSERT INTO post_reaction (type, user_id, post_id) VALUES (?,?,?);")
				if err != nil {
					log.Println(err)
				}
				defer statement.Close()
				statement.Exec(postreaction, userid, postid)

				//notify post author about new reaction
				if pagedata.User.Username != pagedata.Post.Author {
					notification := `Your post was ` + postreaction + `d by ` + pagedata.User.Username + ". Check it out: "
					link := "/post?id=" + strconv.Itoa(postid)
					usr := database.GetUserByUserName(pagedata.Post.Author)
					_ = database.AddNotificationForUser(notification, link, 1, usr.ID)
				}

				//delete record if current reaction and new reaction are the same
			} else if currentpostreaction == postreaction {
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
				statement.Exec(postreaction, userid, postid)
			}

		}

		//if user submitted the comment
		if buttonpressed == "submit-comment" {
			commentid, err := strconv.Atoi(r.FormValue("commentid"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				log.Println("Cannot get categories from DB - category.go")
				return
			}

			commentauthor := database.GetCommentAuthorByID(commentid)

			if commentid > 0 && pagedata.User.Username != commentauthor && pagedata.User.RoleID != 3 {
				w.WriteHeader(http.StatusUnauthorized)
				tmpl.ExecuteTemplate(w, "err401", pagedata)
				return
			}

			//check that comment is not empty string or whitespace
			if strings.TrimSpace(content) == "" {
				pagedata.Message.Msg1 = "Cannot submit empty comment"
				tmpl.ExecuteTemplate(w, "view-post", pagedata)
				return
			}

			if commentid == 0 { //add new comment to database
				statement, err := database.Db.Prepare("INSERT INTO comment ( post_id, user_id, content, datetime) VALUES ( ?, ?, ?, ?);")
				if err != nil {
					log.Println("Database error, cannot insert new comment.")
				}
				defer statement.Close()
				statement.Exec(postid, userid, content, time.Now())

				//notify post author about new comment
				if pagedata.User.Username != pagedata.Post.Author {
					notification := `Your post was commented by ` + pagedata.User.Username + ". Check it out: "
					link := "/post?id=" + strconv.Itoa(postid)
					usr := database.GetUserByUserName(pagedata.Post.Author)
					_ = database.AddNotificationForUser(notification, link, 3, usr.ID)
				}
			} else { //edit existing comment in database
				t := time.Now()
				content = content + "\n\nEdited by " + user.Username + " on " + t.Format(time.Stamp)
				statement, err := database.Db.Prepare("UPDATE comment SET content = ? WHERE comment_id = ?;")
				if err != nil {
					log.Println("Database error, cannot edit new comment.")
				}
				defer statement.Close()
				statement.Exec(content, commentid)

				//If comment was edited by admin, notify comment author about it
				if pagedata.User.Username != commentauthor {
					notification := `Your comment was edited by ` + pagedata.User.Username + ". Check it out: "
					link := "/post?id=" + strconv.Itoa(pagedata.Post.ID)
					usr := database.GetUserByUserName(commentauthor)
					_ = database.AddNotificationForUser(notification, link, 3, usr.ID)
				}
			}
		}

		//Editing post
		if buttonpressed == "edit-post" {
			pagedata.Categories, err = database.GetCategoriesForPostEdit(postid)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				log.Println("Cannot get categories from DB - category.go")
				return
			}
			link := "/new-post?id=" + strconv.Itoa(postid)
			http.Redirect(w, r, link, http.StatusMovedPermanently)

		}
		//Deleting post
		if buttonpressed == "delete-post" {
			log.Println("Deleting post with post_id=", postid)
			statement, _ := database.Db.Prepare("DELETE FROM comment_reaction WHERE comment_id IN (SELECT comment_id FROM comment WHERE post_id = ?)")
			defer statement.Close()
			statement.Exec(postid)

			statement, _ = database.Db.Prepare("DELETE FROM post_reaction WHERE post_id = ? ;")
			defer statement.Close()
			statement.Exec(postid)

			statement, _ = database.Db.Prepare("DELETE FROM comment WHERE post_id = ? ;")
			defer statement.Close()
			statement.Exec(postid)

			statement, _ = database.Db.Prepare("DELETE FROM post_category WHERE post_id = ? ;")
			defer statement.Close()
			statement.Exec(postid)

			statement, _ = database.Db.Prepare("DELETE FROM post WHERE post_id = ?")
			defer statement.Close()
			statement.Exec(postid)

			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

		//Reporting post - creates notification for admins
		if buttonpressed == "report-post" {
			notification := `Please review the following post: `
			link := "/post?id=" + strconv.Itoa(postid)
			_ = database.AddNotificationForRole(notification, link, 7, "administrator")
			pagedata.Message.Msg2 = "Post was successfully reported, admins will review it!"
		}

		//Approving post - creates notification for admins
		if buttonpressed == "approve-post" {
			if pagedata.User.RoleID < 2 {
				w.WriteHeader(http.StatusUnauthorized)
				tmpl.ExecuteTemplate(w, "err401", pagedata)
				return
			}
			statement, err := database.Db.Prepare("UPDATE post SET status = 'approved' WHERE post_id = ?;")
			if err != nil {
				log.Println("Database error, cannot approve post.")
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
			defer statement.Close()
			statement.Exec(pagedata.Post.ID)

			//Notification for post author that post was approved
			notification := `Your post was approved by ` + pagedata.User.Username + ". Check it out: "
			link := "/post?id=" + strconv.Itoa(pagedata.Post.ID)
			usr := database.GetUserByUserName(pagedata.Post.Author)
			_ = database.AddNotificationForUser(notification, link, 7, usr.ID)
		}

		//Declining post - creates notification for admins
		if buttonpressed == "decline-post" {
			if pagedata.User.RoleID < 2 {
				w.WriteHeader(http.StatusUnauthorized)
				tmpl.ExecuteTemplate(w, "err401", pagedata)
				return
			}
			statement, err := database.Db.Prepare("UPDATE post SET status = 'declined' WHERE post_id = ?;")
			if err != nil {
				log.Println("Database error, cannot decline post.")
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
			defer statement.Close()
			statement.Exec(pagedata.Post.ID)

			//Notification for post author that post was declined
			notification := `Your post was declined by ` + pagedata.User.Username + ". Check it out: "
			link := "/post?id=" + strconv.Itoa(pagedata.Post.ID)
			usr := database.GetUserByUserName(pagedata.Post.Author)
			_ = database.AddNotificationForUser(notification, link, 7, usr.ID)
		}

		//Deleting comment
		if buttonpressed == "delete-comment" {
			commentid, err := strconv.Atoi(r.FormValue("commentid"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				log.Println("Cannot get categories from DB - category.go")
				return
			}
			commentauthor := database.GetCommentAuthorByID(commentid)
			if pagedata.User.Username != commentauthor && pagedata.User.RoleID != 3 {
				w.WriteHeader(http.StatusUnauthorized)
				tmpl.ExecuteTemplate(w, "err401", pagedata)
				return
			}

			statement, err := database.Db.Prepare("DELETE FROM comment_reaction WHERE comment_id = ?;")
			if err != nil {
				log.Println("Database error, cannot delete comment.")
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
			defer statement.Close()
			statement.Exec(commentid)

			statement, err = database.Db.Prepare("DELETE FROM comment WHERE comment_id = ?;")
			if err != nil {
				log.Println("Database error, cannot delete comment.")
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
			defer statement.Close()
			statement.Exec(commentid)
		}

		//Reporting comment - creates notification for admins
		if buttonpressed == "report-comment" {
			commentid, err := strconv.Atoi(r.FormValue("commentid"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				log.Println("Cannot get comment from DB")
				return
			}

			commentauthor := database.GetCommentAuthorByID(commentid)
			notification := `Please review comment written by ` + commentauthor + ". Link to post: "
			link := "/post?id=" + strconv.Itoa(postid)
			_ = database.AddNotificationForRole(notification, link, 5, "administrator")
			pagedata.Message.Msg2 = "Comment was successfully reported, admins will review it!"
		}

		//Editing comment
		if buttonpressed == "edit-comment" {
			commentid, err := strconv.Atoi(r.FormValue("commentid"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				log.Println("Cannot get comment from DB")
				return
			}
			commentauthor := database.GetCommentAuthorByID(commentid)
			if pagedata.User.Username != commentauthor && pagedata.User.RoleID != 3 {
				w.WriteHeader(http.StatusUnauthorized)
				tmpl.ExecuteTemplate(w, "err401", pagedata)
				return
			}

			pagedata.Comment, err = database.GetCommentById(commentid)
			if err != nil {
				log.Println(err)
				//write internal server error header and template
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
		}

		pagedata.Comments, err = database.GetCommentsByPostAndUserId(postid, user.ID)
		if err != nil {
			log.Println(err)
			//write internal server error header and template
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.ExecuteTemplate(w, "err500", pagedata)
			return
		}

		pagedata.Post, err = database.GetPostByPostAndUsedID(postid, user.ID)
		if err != nil {
			log.Println(err)
			//write internal server error header and template
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.ExecuteTemplate(w, "err500", pagedata)
			return
		}

	}

	tmpl.ExecuteTemplate(w, "view-post", pagedata)

}
