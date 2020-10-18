package user

var users []*User

func get() ([]*User, error) {
	//s.db.get
	return users, nil
}

func create(user *User) (*User, error) {
	users = append(users, user)
	return user, nil
}
