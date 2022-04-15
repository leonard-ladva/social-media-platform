package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"git.01.kood.tech/Rostislav/real-time-forum/errors"

	uuid "github.com/satori/go.uuid"
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
func GetUserByUserName(username string) User {
	var user User
	cmd := "SELECT ID, Nickname, Password, Email, FirstName, LastName, Gender, Age, Color, CreatedAt FROM user WHERE Nickname = ? OR Email = ?"
	row := DB.QueryRow(cmd, username)

	row.Scan(&user.ID, &user.Nickname, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Gender, &user.Age, &user.Color, &user.CreatedAt)
	log.Println("Getting user from database | User: ", username)
	return user
}

func IfUserExist(field string, value string) bool {
	var user User
	cmd := fmt.Sprintf(`SELECT Nickname, Email FROM user WHERE %s = ?`, field)
	row := DB.QueryRow(cmd, value)
	err := row.Scan(&user.Nickname, &user.Email)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		errors.ErrorLog.Print(err)
	}
	return true
}

func (user *User) Insert() {
	stmt, err := DB.Prepare("INSERT INTO User (ID, Email, Password, Nickname, FirstName, LastName, Gender, Age, Color, CreatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		errors.ErrorLog.Print("Database: failed when inserting user to database.")
	}
	defer stmt.Close()

	id := uuid.NewV4()
	color := newPastelColor()
	createdAt := time.Now().UnixNano()

	stmt.Exec(id, user.Email, user.Password, user.Nickname, user.FirstName, user.LastName, user.Gender, user.Age, color, createdAt)
}

func newPastelColor() string {
	rand.Seed(time.Now().UnixNano())

	return "hsl(" + strconv.Itoa(rand.Intn(360-0)) + "," +
		strconv.Itoa(25+rand.Intn(70-0)) + "%," +
		strconv.Itoa(85+rand.Intn(10-0)) + "%)"
}

func (u *User) IsValid() (url.Values) {
	errs := url.Values{} 
	// Email
	if IfUserExist("Email", u.Email) {
		errs.Add("Email", "Email already in use")
	}
	if !checkCharacters("Email", u.Email) {
		errs.Add("Email", "Please enter a valid email")
	}
	// Nickname
	if IfUserExist("Nickname", u.Nickname) {
		errs.Add("Nickname", "Sorry! Nickname already taken")
	}
	if !checkCharacters("Nickname", u.Nickname) || !checkLength("Nickname", u.Nickname) {
		errs.Add("Nickname", fmt.Sprintf("Nickname has to be between %d and %d characters and contain only letters and numbers",
			lengthReq["Nickname"][0], lengthReq["Nickname"][1]))
	}
	// Password
	if !checkCharacters("Password", u.PasswordPlain) || !checkLength("Password", u.PasswordPlain) {
		errs.Add("Password", fmt.Sprintf("Password has to be between %d and %d characters",
			lengthReq["Password"][0], lengthReq["Password"][1]))
	}
	// Password Confirm
	if u.PasswordPlain != u.PasswordConfirm {
		errs.Add("PasswordConfirm", "Passwords don't match")
	}
	// FirstName
	if !checkCharacters("FirstName", u.FirstName) || !checkLength("FirstName", u.FirstName) {
		errs.Add("FirstName", fmt.Sprintf("FirstName has to be between %d and %d characters and contain only letters",
			lengthReq["FirstName"][0], lengthReq["FirstName"][1]))
	}
	// LastName
	if !checkCharacters("LastName", u.LastName) || !checkLength("LastName", u.LastName) {
		errs.Add("LastName", fmt.Sprintf("LastName has to be between %d and %d characters and contain only letters",
			lengthReq["LastName"][0], lengthReq["LastName"][1]))
	}
	// Gender
	if !checkCharacters("Gender", u.Gender) || !checkLength("Gender", u.Gender) {
		errs.Add("Gender", fmt.Sprintf("Gender has to be between %d and %d characters and contain only letters",
			lengthReq["Gender"][0], lengthReq["Gender"][1]))
	}
	
	return errs
}

func checkLength(field string, value string) bool {
	return (len(value) >= lengthReq[field][0] && len(value) <= lengthReq[field][1])
}

func checkCharacters(field string, value string) bool {
	match := regexp.MustCompile(characterReq[field]).MatchString(value)
	return match
}
