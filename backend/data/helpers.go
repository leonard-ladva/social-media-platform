package data

import (
	"errors"
	"regexp"
	"time"
)

func CurrentTime() int64 {
	return time.Now().UnixMilli()
}

// lengthReq specifies the min-max lenghts of data fields
var lengthReq = map[string][]int{
	"Nickname":  {3, 20},
	"Password":  {8, 50},
	"FirstName": {1, 50},
	"LastName":  {1, 50},
	"Gender":    {1, 50},
	"Content":   {1, 500},
	"Tag":       {1, 50},
}

// characterReq specifies the allowed characters and format of data fields
var characterReq = map[string]string{
	"Nickname":  "^[a-zA-Z0-9]*$", // letters, numbers
	"Password":  "^[ -~]*$",       // all ascii characters in range space to ~
	"FirstName": "^[a-zA-Z]*$",    // letters
	"LastName":  "^[a-zA-Z]*$",    // letters
	"Gender":    "^[a-zA-Z ]*$",   // letters and spaces
	"Email":     "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
	"Content":   "(?m)^[ -~]*$", // all ascii characters in range space to ~
	"Tag":       "^[ -~]*$", // all ascii characters in range space to ~
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

func ChatID(user1ID string, user2ID string) (string, error) {
	if len(user1ID) != 36 || len(user2ID) != 36 {
		return "", errors.New("datahelpers: one or both users IDs not valid")
	}
	abc := user1ID < user2ID
	if abc {
		return (user1ID + user2ID), nil
	}
	return (user2ID + user1ID), nil
}
