package request

import (
	"forum-test/database"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func Config(w http.ResponseWriter, r *http.Request, user database.User) {
	var pagedata database.PageData
	pagedata.User = user
	tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))
	pagedata.Categories, _ = database.GetCategories()

	switch r.Method {

	case "GET":
		tmpl.ExecuteTemplate(w, "config", pagedata)

	case "POST":

		r.ParseForm()
		buttonpressed := r.FormValue("submit")
		username := r.FormValue("username")
		role, _ := strconv.Atoi(r.FormValue("role"))
		category_title := r.FormValue("category-title")
		category_desc := r.FormValue("category-description")
		category_link := r.FormValue("category-link")
		category_id, _ := strconv.Atoi(r.FormValue("category-id"))

		if buttonpressed == "delete" {
			database.Db.Exec("DELETE FROM user")
			database.Db.Exec("DELETE FROM role")
			database.Db.Exec("DELETE FROM session")
			database.Db.Exec("DELETE FROM post")
			database.Db.Exec("DELETE FROM comment")
			database.Db.Exec("DELETE FROM category")
			database.Db.Exec("DELETE FROM post_reaction")
			database.Db.Exec("DELETE FROM comment_reaction")
			database.Db.Exec("DELETE FROM post_category")
			database.Db.Exec("DELETE FROM notification")
			database.Db.Exec("DELETE FROM sqlite_sequence")

			log.Println("Deleting all database records")
			pagedata.Message.Msg2 = "Deleted all database records"
		}

		//Adding dummy data to database
		if buttonpressed == "create" {
			err := GenerateDummyData()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.ExecuteTemplate(w, "err500", user)
			}
			pagedata.Message.Msg2 = "Data was successfully added to database. Created users: admin, moderator, user1...user3 (password for every user is the same as login)"
			log.Println("Dummy database records created")
		}

		//Changing user role
		if buttonpressed == "change-role" {
			if username == "" {
				pagedata.Message.Msg1 = "Please enter username!"
				tmpl.ExecuteTemplate(w, "config", pagedata)
				return
			}

			if pagedata.User.Username == username {
				pagedata.Message.Msg1 = "You can not change your own role!"
				tmpl.ExecuteTemplate(w, "config", pagedata)
				return
			}

			usr := database.GetUserByUserName(username)
			if usr.ID == 0 {
				pagedata.Message.Msg1 = "User " + username + " not found!"
				tmpl.ExecuteTemplate(w, "config", pagedata)
				return
			}

			statement, err := database.Db.Prepare("UPDATE user SET role_id = ? WHERE username = ?;")
			if err != nil {
				log.Println("Database error, cannot change role.")
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
			defer statement.Close()
			statement.Exec(role, username)
			notification := "Your role was changed by " + pagedata.User.Username
			link := ""
			_ = database.AddNotificationForUser(notification, link, 10, usr.ID)
			pagedata.Message.Msg2 = "Role for user " + username + " was successfully changed!"
		}

		//Adding new category
		if buttonpressed == "add-category" {
			if category_title == "" || category_desc == "" || category_link == "" {
				pagedata.Message.Msg1 = "Please fill in all fields to add new category"
				tmpl.ExecuteTemplate(w, "config", pagedata)
				return
			}
			_, err := url.ParseRequestURI(category_link)
			if err != nil {
				pagedata.Message.Msg1 = "Please provide valid link"
				tmpl.ExecuteTemplate(w, "config", pagedata)
				return
			}
			statement, err := database.Db.Prepare("INSERT INTO category (title, description, img_link) VALUES (?, ?, ?);")
			if err != nil {
				log.Println("Database error, cannot insert new category")
			}
			defer statement.Close()
			statement.Exec(category_title, category_desc, category_link)
			pagedata.Message.Msg2 = "Category " + category_title + " was successfully created!"
		}

		//Deleting category
		if buttonpressed == "delete-category" {
			statement, _ := database.Db.Prepare("DELETE FROM post_category WHERE category_id = ? ;")
			defer statement.Close()
			statement.Exec(category_id)

			statement, _ = database.Db.Prepare("DELETE FROM category WHERE category_id = ?;")
			defer statement.Close()
			statement.Exec(category_id)

			pagedata.Message.Msg2 = "Category was successfully deleted!"
		}

		tmpl.ExecuteTemplate(w, "config", pagedata)
	}

}
