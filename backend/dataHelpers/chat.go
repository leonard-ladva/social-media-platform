package dataHelpers

import (
	"errors"
)

func ChatID(user1ID string, user2ID string) (string, error) {
	if len(user1ID) != 36 || len(user2ID) != 36 {
		return "", errors.New("datahelpers: one or both users IDs not acceptable")
	}
	return (user1ID + user2ID), nil
}

// func ValidMessage
