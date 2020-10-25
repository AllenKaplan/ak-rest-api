package user

import "math/rand"

func (s *UserService) Create(user *User) (*User, error) {
	//user given random id lol
	user.UserID = rand.Intn(100000) + 1000000

	resp, _ := s.db.create(user)
	return resp, nil
}

func (s *UserService) Get() ([]*User, error) {
	resp, _ := s.db.get()
	return resp, nil
}
