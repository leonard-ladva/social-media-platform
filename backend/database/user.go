package database

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"git.01.kood.tech/Rostislav/real-time-forum/errors"

	uuid "github.com/satori/go.uuid"
)

// func GetUserByUserName(username string) User {
// 	var user User
// 	row := DB.QueryRow("SELECT u.user_id, u.username, u.password, u.email")
// }

func IfUserExist(field string, value string) bool {
	var user User
	cmd := fmt.Sprintf(`SELECT Nickname, Email FROM user WHERE %s = ?`, field)
	// row := DB.QueryRow("SELECT Nickname, Email FROM user WHERE Nickname = ?", value)
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
