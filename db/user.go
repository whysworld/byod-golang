package db

import "log"

func ValidUser(username, password string) bool {
	passwordFromDB := "123456"
	usernameFromDB := "test@gmail.com"
	log.Print("validating user ", username)
	if password == passwordFromDB && username == usernameFromDB{
		return true
	}
	return false
}
