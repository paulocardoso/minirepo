package main

type User struct {
	ID string `json:"id,omitempty"`
}

func (u User) Error() string {
	panic("implement me")
}

type Repository interface {
	GetUser(user User) error
}

func (u User) GetID() error {
	return User{
		ID: "asd",
	}
}