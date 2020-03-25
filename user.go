package main

import (
	"github.com/google/uuid"
)

// User Model Type
type User struct{}

// UserDTO is used to hold the properties
// of the User model
type UserDTO struct {
	id       string
	email    string
	password string
}

var users = []ModelDTO{}

// UserModel - user model instance
var UserModel = &User{}

// UserRepo - for data store access
var UserRepo = NewModel(User{})

// Model.isDTO implementation MUST have this
func (user UserDTO) getID() string {
	return user.id
}

// Model implementation MUST have this
func (u User) getTableData() interface{} {
	return users
}

// Model implementation MUST have this
func (User) getTableName() string {
	return "users"
}

func makeDTO(t UserDTO) ModelDTO {
	return ModelDTO{t}
}

func (User) initialiseTable() {
	seed := []UserDTO{
		UserDTO{
			id:       uuid.New().String(),
			email:    "jsamchineme@example.test",
			password: "password",
		},
		UserDTO{
			id:       uuid.New().String(),
			email:    "john@example.test",
			password: "password",
		},
		UserDTO{
			id:       uuid.New().String(),
			email:    "ekene@gmail.test",
			password: "password",
		},
	}

	for i := range seed {
		user := seed[i]
		hashedPassword, _ := HashPassword(user.password)
		d := MakeDTO(UserDTO{
			id:       user.id,
			email:    user.email,
			password: hashedPassword,
		})
		UserRepo.CreateRecord(d)
		// users = append(users, d)
	}
}
