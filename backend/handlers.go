package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"git.01.kood.tech/Rostislav/real-time-forum/database"
	"git.01.kood.tech/Rostislav/real-time-forum/errors"

	"golang.org/x/crypto/bcrypt"
)

var lengthReq = map[string][]int{
	"Nickname":  {3, 20},
	"Password":  {8, 50},
	"FirstName": {1, 50},
	"LastName":  {1, 50},
	"Gender":    {1, 50},
}

var characterReq = map[string]string{
	"Nickname":  "^[a-zA-Z0-9]*$", // letters, numbers
	"Password":  "^[ -~]*$",       // all ascii characres in range space to ~
	"FirstName": "^[a-zA-Z]*$",    // letters
	"LastName":  "^[a-zA-Z]*$",    // letters
	"Gender":    "^[a-zA-Z ]*$",   // letters and spaces
	"Email":     "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
}

type formError struct {
	field   string
	message string
}

var fields = []string{"Email", "Nickname", "FirstName", "LastName", "Gender"}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var user database.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		errors.ServerError(w, err)
		return
	}

	var formErrors []formError
	fmt.Println(formErrors)

	// v := reflect.ValueOf(user)
	//  Check characters
	// for _, field := range fields {
	// 	fmt.Println(v.FieldByName(field))
	// 	checkCharacters(field, string(v.FieldByName(field)))
	// }

	// Email
	fmt.Println("Email okay?",
		!database.IfUserExist("Email", user.Email),
		checkCharacters("Email", user.Email))

	// Nickname
	fmt.Println("Nickn okay?",
		!database.IfUserExist("Nickname", user.Nickname),
		checkCharacters("Nickname", user.Nickname),
		checkLength("Nickname", user.Nickname))

	// Password
	fmt.Println("Psswd okay?",
		checkCharacters("Password", user.PasswordPlain),
		checkLength("Password", user.PasswordPlain))

	// Password Confirm
	fmt.Println("Cnfrm okay?", user.PasswordConfirm == user.PasswordPlain)

	// FirstName
	fmt.Println("FName okay?",
		checkCharacters("FirstName", user.FirstName),
		checkLength("FirstName", user.FirstName))

	// LastName
	fmt.Println("LName okay?",
		checkCharacters("LastName", user.LastName),
		checkLength("LastName", user.LastName))

	// Gender
	fmt.Println("Gendr okay?",
		checkCharacters("Gender", user.Gender),
		checkLength("Gender", user.Gender))

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordPlain), 10)
	if err != nil {
		errors.ServerError(w, err)
	}
	user.Password = passwordHash

	user.Insert()

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var user database.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		errors.ServerError(w, err)
		return
	}

	var formErrors []formError
	fmt.Println(formErrors)

	if database.IfUserExist("Email", user.Email) || database.IfUserExist("Nickname", user.Email) {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
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
	}
}
func checkLength(field string, value string) bool {
	return (len(value) >= lengthReq[field][0] && len(value) <= lengthReq[field][1])
}

func checkCharacters(field string, value string) bool {
	match := regexp.MustCompile(characterReq[field]).MatchString(value)
	return match
}
