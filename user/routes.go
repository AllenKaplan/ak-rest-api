package user

func (s UserService) Create(user *User) (*User, error) {
	//user still has no id
	resp, _ := create(user)
	return resp, nil
}

func (s UserService) Get() ([]*User, error) {
	resp, _ := get()
	return resp, nil
}
