package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"git.01.kood.tech/Rostislav/real-time-forum/data"

	"golang.org/x/crypto/bcrypt"
)

type StringInt int

type LoginResponse struct {
	Token   interface{} `json:"token"`
	User    *data.User `json:"user"`
}

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

// Register gets data from the client, checks the data, if data is valid inserts a new user to the database and responds to the client
func Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user data.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	valid, err := user.IsValid()
	if valid {
		err := user.Insert()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

// isUnique receives an email/nickname from the client and responds if it is available
func IsUnique(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user data.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var field = "Email"
	var value = user.Email
	if user.Email == "" {
		field = "Nickname"
		value = user.Nickname
	}

	exists, _, err := data.GetUser(field, value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists {
		resp, err := json.Marshal(false)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(resp)
		return
	}

	resp, err := json.Marshal(true)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

// Login
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &data.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var nickOrEmail string = "Nickname"
	var value string = user.Nickname
	// Determine whether user inserted email or nickname
	if strings.Contains(user.Nickname, "@") {
		user.Email = value
		user.Nickname = ""
		nickOrEmail = "Email"
		value = user.Email
	}
	// Check if user with email/nickname exists
	exists, dbUser, err := data.GetUser(nickOrEmail, value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Check if provided password mathces the one in the db
	err = bcrypt.CompareHashAndPassword(dbUser.Password, []byte(user.PasswordPlain))
	if err != nil {
		log.Printf("Wrong password for user with %s of %s", nickOrEmail, value)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user = dbUser

	log.Println("LOGIN: Info OK")
	// Add a session to the db
	token, err := user.AddSession()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send a successful login response to the client
	response := LoginResponse{
		Token:   token,
		User:    user,
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func CurrentUser(w http.ResponseWriter, r *http.Request) {
	var user = data.CurrentUser
	user.Password = []byte("")

	data, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
