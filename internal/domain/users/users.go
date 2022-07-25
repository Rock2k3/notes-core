package users

import "github.com/google/uuid"

type MyUser struct {
	UserUUID uuid.UUID
	Name     string
}

type MyUsers interface {
	GetUserByUUID(uuid.UUID) (*MyUser, error)
	GetUserByName(string) (*MyUser, error)
	AddUser(string) (*MyUser, error)
}
