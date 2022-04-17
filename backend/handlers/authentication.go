package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"git.01.kood.tech/Rostislav/real-time-forum/data"
	uuid "github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"
)

type StringInt int

type LoginResponse struct {
	Message string
	Token   uuid.UUID
	User    UserInfo
}

type UserInfo struct {
	ID       string
	Email    string
	NickName string
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
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user data.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	valid, err := user.IsValid()
	if valid {
		user.Insert()
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

// isUnique receives an email/nickname from the client and responds if it is available
func IsUnique(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user data.User

	err = json.Unmarshal(body, &user)
	if err != nil {
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists {
		resp, err := json.Marshal(false)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(resp)
		return
	}

	resp, err := json.Marshal(true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

// Login
func Login(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var nickOrEmail string = "NickName"
	var value string = user.Nickname
	// Determine whether user inserted email or nickname
	if strings.Contains(user.Nickname, "@") {
		user.Email = value
		user.Nickname = ""
		nickOrEmail = "Email"
		value = user.Email
	}
	exists, dbUser, err := data.GetUser(nickOrEmail, value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword(dbUser.Password, []byte(user.PasswordPlain))
	if err != nil {
		log.Printf("Wrong password for user with %s of %s", nickOrEmail, value)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user = dbUser

	log.Println("LOGIN: Info OK")

	token, err := user.AddSession()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := LoginResponse{
		Message: "Success",
		Token:   token,
		User: UserInfo{
			ID:       user.ID,
			Email:    user.Email,
			NickName: user.Nickname,
		},
	}

	data, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// Session checks if the client in logged in
func Session(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	token := r.Header.Get("Authorization")[7:]
	fmt.Println(token)

	session, err := data.GetSession(token)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user data.User
	user.Password = []byte("")
	_, user, err = data.GetUser("ID", session.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)

}
