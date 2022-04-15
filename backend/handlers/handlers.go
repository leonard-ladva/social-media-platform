package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
	"git.01.kood.tech/Rostislav/real-time-forum/errors"

	"golang.org/x/crypto/bcrypt"
)

type StringInt int

func (st *StringInt) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = StringInt(v)
	case float64:
		*st = StringInt(int(v))
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*st = StringInt(i)
	}
	return nil
}
func Register(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var user data.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		errors.ErrorLog.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errs := user.IsValid()
	fmt.Println(errs)
	if len(errs) == 0 {
		fmt.Println("Success your accounts info is superb")
		// Successful register
		// business logic
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordPlain), 10)
		if err != nil {
			errors.ServerError(w, err)
		}
		user.Password = passwordHash
		user.Insert()
		w.WriteHeader(http.StatusOK)
		return
	}
	formErrs, err := json.Marshal(errs)
	if err != nil {
		errors.ErrorLog.Println("encoding formErrs to JSON")
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write(formErrs)
	// display errors on register form
	fmt.Println("Errors in form!!!!")
	return
}

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	enableCors(&w)
// 	if r.Method == "OPTIONS" {
// 		return
// 	}

// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var user data.User

// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		log.Println(err)
// 		errors.ServerError(w, err)
// 		return
// 	}

// 	if data.IfUserExist("Email", user.Email) || data.IfUserExist("Nickname", user.Email) {
// 		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
// 		if err != nil {
// 			log.Println("Wrong password for user: ", user.Username)
// 			pagedata.Message.Msg1 = "Haha, wrong!"
// 			tmpl.ExecuteTemplate(w, "login", pagedata)
// 		}

// 		if user.Username == username || user.Email == username {
// 			log.Println("Success, username & password match 🔓")

// 		} else {
// 			log.Println("Access denied, no cookies for you! 😈")
// 			return
// 		}

// 		data.AddSession(w, r, user)
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 	}
// }
func submitPost(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var post data.Post

	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Println(err)
		errors.ServerError(w, err)
		return
	}

	post.Insert()
}

func IsUnique(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var user data.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		errors.ServerError(w, err)
		return
	}

	var field = "Email" 
	var value = user.Email
	if user.Email == "" {field = "Nickname"; value = user.Nickname} 
	
	if data.IfUserExist(field, value) {
		resp, err := json.Marshal(false)
		if err != nil {
			errors.ServerError(w, err)
		}
		w.Write(resp)
		return
	}
	resp, err := json.Marshal(true)
	if err != nil {
		errors.ServerError(w, err)
	}
	w.Write(resp)
}
