package request

import (
	"forum-test/database"
	"html/template"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request, user database.User) {
	tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))

	var pagedata database.PageData
	pagedata.User = user

	// Get client ID-s from OS environment variables
	GCLIENT_ID = os.Getenv("GCLIENT_ID")   //Google client id
	FBCLIENT_ID = os.Getenv("FBCLIENT_ID") //Facebook app id
	GHCLIENT_ID = os.Getenv("GHCLIENT_ID") //Github client/app id

	//Data for HTML template (OAuth links)
	pagedata.Authdata.GoogleClientID = GCLIENT_ID
	pagedata.Authdata.FacebookClientID = FBCLIENT_ID
	pagedata.Authdata.GitHubClientID = GHCLIENT_ID
	pagedata.CALLBACK_URI = CALLBACK_URI

	switch r.Method {
	case "POST":

		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		user = database.GetUserByUserName(username)
		log.Println("Authenticating user: ", user.Username)

		//let users sign in by either username or e-mail

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			log.Println("Wrong password for user: ", user.Username)
			pagedata.Message.Msg1 = "Haha, wrong!"
			tmpl.ExecuteTemplate(w, "login", pagedata)
		}

		if user.Username == username || user.Email == username {
			log.Println("Success, username & password match 🔓")

		} else {
			log.Println("Access denied, no cookies for you! 😈")
			return
		}

		database.AddSession(w, r, user)
		http.Redirect(w, r, "/", http.StatusSeeOther)

	case "GET":

		tmpl.ExecuteTemplate(w, "login", pagedata)

	}

}
