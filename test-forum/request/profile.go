package request

import (
	"forum-test/database"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func Profile(w http.ResponseWriter, r *http.Request, user database.User) {

	tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))

	var pagedata database.PageData

	pagedata.User = user
	pagedata.Notifications, _ = database.GetNotificationsByUserId(user.ID)

	switch r.Method {
	case "GET":
		//Marking all unread notifications as read
		for i := range pagedata.Notifications {
			statement, err := database.Db.Prepare("UPDATE notification SET read = 1 WHERE notification_id = ?;")
			if err != nil {
				log.Println("Database error, cannot update notification.")
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
			defer statement.Close()
			statement.Exec(pagedata.Notifications[i].ID)
		}
		tmpl.ExecuteTemplate(w, "profile", pagedata)
		return

	case "POST":
		if user.RoleID < 1 {
			w.WriteHeader(http.StatusUnauthorized)
			tmpl.ExecuteTemplate(w, "err401", pagedata)
			return
		}

		r.ParseForm()
		userid := pagedata.User.ID
		oldPass := r.FormValue("old-pass")
		newPass1 := r.FormValue("new-pass1")
		newPass2 := r.FormValue("new-pass2")

		application := r.FormValue("application")
		buttonpressed := r.FormValue("submit")

		//If user decided to change password
		if buttonpressed == "change-password" {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPass))
			if err != nil {
				log.Println("Wrong old password provided 🔑")
				pagedata.Message.Msg1 = "Wrong old password! Please retry"
				tmpl.ExecuteTemplate(w, "profile", pagedata)
				return
			}

			if newPass1 != newPass2 {
				log.Println("Passwords don´t match 😔")
				pagedata.Message.Msg1 = "Provided new passwords don´t match! Please be more attentive"
				tmpl.ExecuteTemplate(w, "profile", pagedata)
				return
			}

			if !SecurePassword(newPass1) {
				log.Println("Password is not secure enough 🔑")
				pagedata.Message.Msg1 = "New password is too weak. Needed:\n8-16 symbols, lowercase letter, capital letter, number"
				tmpl.ExecuteTemplate(w, "profile", pagedata)
				return
			}

			encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newPass1), 10)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}

			//change password
			statement, err := database.Db.Prepare("UPDATE user SET password = ? WHERE user_id = ?;")
			if err != nil {
				log.Println("Database error, cannot change password.")
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
			defer statement.Close()
			statement.Exec(encryptedPassword, userid)
			pagedata.Message.Msg2 = "Password was successfully changed!"
		}

		//If user submitted application to become moderator
		if buttonpressed == "become-moderator" {
			if application == "" {
				log.Println("Empty application submitted")
				pagedata.Message.Msg1 = "Please add a good reasoning why you want to become moderator!"
				tmpl.ExecuteTemplate(w, "profile", pagedata)
				return
			}
			notification := "User " + pagedata.User.Username + " wants to become moderator. Application: " + application
			link := "/config"
			err := database.AddNotificationForRole(notification, link, 10, "administrator")
			if err != nil {
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
			pagedata.Message.Msg2 = "Application succesfully submitted. Admins will review it!"
		}

		//If user wants to hide notification
		if buttonpressed == "hide-notification" {
			notification, err := strconv.Atoi(r.FormValue("notificationid"))
			statement, err := database.Db.Prepare("UPDATE notification SET priority = -1 WHERE notification_id = ?;")
			if err != nil {
				log.Println("Database error, cannot hide notification.")
				tmpl.ExecuteTemplate(w, "err500", pagedata)
				return
			}
			defer statement.Close()
			statement.Exec(notification)
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
		}

		tmpl.ExecuteTemplate(w, "profile", pagedata)
	}
}
