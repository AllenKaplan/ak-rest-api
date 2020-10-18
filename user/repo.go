package user

import (
	"math/rand"
)

var users map[int]*User

func get() ([]*User, error) {
	//s.db.get
	var userList []*User
	for _, u := range users {
		userList = append(userList, u)
	}
	return userList, nil
}

func create(user *User) (*User, error) {
	user.UserID = rand.Intn(100000) + 1000000

	if users == nil {
		users = make(map[int]*User)
	}

	users[user.UserID] = user
	return user, nil
}
