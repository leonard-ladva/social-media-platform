package data

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

var CurrentUser *User

// lengthReq specifies the min-max lenghts of user data fields
var lengthReq = map[string][]int{
	"Nickname":  {3, 20},
	"Password":  {8, 50},
	"FirstName": {1, 50},
	"LastName":  {1, 50},
	"Gender":    {1, 50},
}

// characterReq specifies the allowed characters and format of user data fields
var characterReq = map[string]string{
	"Nickname":  "^[a-zA-Z0-9]*$", // letters, numbers
	"Password":  "^[ -~]*$",       // all ascii characres in range space to ~
	"FirstName": "^[a-zA-Z]*$",    // letters
	"LastName":  "^[a-zA-Z]*$",    // letters
	"Gender":    "^[a-zA-Z ]*$",   // letters and spaces
	"Email":     "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
}

type StringInt int

// UnmarshalJSON unmarshals user data to a struct
func (st *StringInt) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return errors.New("data: unmarshaling user data")
	}
	switch v := item.(type) {
	case int:
		*st = StringInt(v)
	case float64:
		*st = StringInt(int(v))
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("data: converting user age")
		}
		*st = StringInt(i)
	}
	return nil
}

// GetUser gets a user from the database
func GetUser(field string, value string) (bool, *User, error) {
	u := &User{}
	cmd := fmt.Sprintf(`SELECT ID, Email, Password, Nickname, FirstName, LastName, Gender, Age, Color, CreatedAt FROM user WHERE %s = ?`, field)
	row := DB.QueryRow(cmd, value)
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Nickname, &u.FirstName, &u.LastName, &u.Gender, &u.Age, &u.Color, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return false, u, nil
	} else if err != nil {
		return false, u, errors.New("data: getting user")
	}
	return true, u, nil
}

// Insert generates missing fields and insertes a new user into the database
func (user *User) Insert() error {
	stmt, err := DB.Prepare("INSERT INTO User (ID, Email, Password, Nickname, FirstName, LastName, Gender, Age, Color, CreatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return errors.New("data: inserting user")
	}
	defer stmt.Close()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordPlain), 10)
	if err != nil {
		return errors.New("data: generating user password hash")
	}
	user.Password = passwordHash

	id := uuid.NewV4()
	color := newPastelColor()
	createdAt := CurrentTime()

	stmt.Exec(id, user.Email, user.Password, user.Nickname, user.FirstName, user.LastName, user.Gender, user.Age, color, createdAt)
	return nil
}

// NewPastelColor generates a color code for a random pastel color
func newPastelColor() string {
	rand.Seed(time.Now().UnixNano())
	return "hsl(" + strconv.Itoa(rand.Intn(360-0)) + "," +
		strconv.Itoa(25+rand.Intn(70-0)) + "%," +
		strconv.Itoa(85+rand.Intn(10-0)) + "%)"
}

// IsValid check all the user fields and returns true if valid
func (u *User) IsValid() (bool, error) {
	errs := url.Values{}
	// Email
	exists, _, err := GetUser("Email", u.Email)
	if err != nil {
		return false, err
	}
	if exists {
		errs.Add("Email", "Email already in use")
	}
	if !checkCharacters("Email", u.Email) {
		errs.Add("Email", "Please enter a valid email")
	}
	// Nickname
	exists, _, err = GetUser("Nickname", u.Nickname)
	if err != nil {
		return false, err
	}
	if exists {
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
	if len(errs) != 0 {
		fmt.Println("Form Errors: ", errs)
		return false, nil
	}
	return true, nil
}

// checkLength checks a user field length based on var lenghtReq
func checkLength(field string, value string) bool {
	return (len(value) >= lengthReq[field][0] && len(value) <= lengthReq[field][1])
}

// checkCharacters checks a user field chars besed on var characterReq
func checkCharacters(field string, value string) bool {
	match := regexp.MustCompile(characterReq[field]).MatchString(value)
	return match
}

func GetAllUsers() ([]*User, error) {
	var users []*User
	query := "SELECT ID, Nickname, Color, CreatedAt FROM User ORDER BY Nickname ASC"
	rows, err := DB.Query(query)
	if err == sql.ErrNoRows {
		return users, nil
	}
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("data: getting users")
	}

	for rows.Next() {
		user := &User{}

		err := rows.Scan(&user.ID, &user.Nickname, &user.Color, &user.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("data: getting users")
		}
		users = append(users, user)
	}

	return users, nil
}
