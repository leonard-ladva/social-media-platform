package request

import (
	"forum-test/database"
	"html/template"
	"log"
	"net/http"
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request, user database.User) {

	tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))
	var pagedata database.PageData
	pagedata.User = user

	switch r.Method {
	case "POST":

		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		password2 := r.FormValue("password2")
		email := r.FormValue("email")

		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			//figure out what to do - internal server error? - is this correct?
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.ExecuteTemplate(w, "err500", user)
		}

		user = database.GetUserByUserName(username)
		if user.Username == username {
			log.Println("Username taken", username)
			pagedata.Message.Msg1 = "Username taken! Choose a less cool username."
			tmpl.ExecuteTemplate(w, "register", pagedata)
			return
		}

		submittedEmail := database.GetUserByUserName(email)
		if submittedEmail.Email == email {
			log.Println("E-mail taken", email)
			pagedata.Message.Msg1 = "E-mail address already registered. Did you forget?"
			tmpl.ExecuteTemplate(w, "register", pagedata)
			return
		}

		if password != password2 {
			log.Println("Passwords don't match 😔")
			pagedata.Message.Msg1 = "Provided passwords don´t match! Please be more attentive"
			tmpl.ExecuteTemplate(w, "register", pagedata)
			return
		}

		if !SecurePassword(password) {
			log.Println("Password is not secure enough 🔑")
			pagedata.Message.Msg1 = "Password is too weak. Needed:\n8-16 symbols, lowercase letter, capital letter, number"
			tmpl.ExecuteTemplate(w, "register", pagedata)
			return
		}

		if !ValidEmail(email) {
			log.Println("Email is not valid 📥")
			pagedata.Message.Msg1 = "Please check e-mail"
			tmpl.ExecuteTemplate(w, "register", pagedata)
			return
		}

		//add the user to database
		statement, err := database.Db.Prepare("INSERT INTO user (username, password, email, reg_datetime, role_id) VALUES (?, ?, ?, ?, ?);")
		if err != nil {
			log.Println("Database error, cannot insert new user.")
		}
		defer statement.Close()

		statement.Exec(username, encryptedPassword, email, time.Now(), 1)
		user = database.GetUserByUserName(username)
		database.AddSession(w, r, user)
		http.Redirect(w, r, "/", http.StatusSeeOther)

	case "GET":

		tmpl.ExecuteTemplate(w, "register", pagedata)
	}

}

func SecurePassword(pass string) bool {
	if len(pass) < 8 || len(pass) > 16 {
		return false
	}
	LowerCase := false
	UpperCase := false
	Number := false
	for i := range pass {
		if pass[i] >= 65 && pass[i] <= 90 {
			UpperCase = true
		}
		if pass[i] >= 97 && pass[i] <= 122 {
			LowerCase = true
		}
		if pass[i] >= 48 && pass[i] <= 57 {
			Number = true
		}
	}
	return UpperCase && LowerCase && Number
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
