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

func (user UserDTO) getID() string {
	return user.id
}

// TODO: make the "users" variable private
// declare the value to holder table collections
var users []UserDTO = []UserDTO{}

// UserModel instance
var UserModel User = User{}

// This method MUST be implemented by an object
// to be qualified as a [Model]
func (user User) getTableData() []UserDTO {
	return users
}

// This method MUST be implemented by an object
// to be qualified as a [Model]
func (user User) getTableName() string {
	return "users"
}

func (User) initialiseTable() {
	users = []UserDTO{
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

	// to update a property of "user" by refering to it like so - "for i, user := range users" would not work
	// because range creates a copy of the user item from the splice,
	// hence the syntax below - user := &users[i], this user value is now a pointer that can be updated
	for i := range users {
		// There's the need to wrap this block within the type assertion for User
		// because each user is seen as an Entity type
		// Wrapping the code in the assertion block makes it so that the "u" variable refers to a User type
		// which will then make field password, email accessible
		user := &users[i]

		hashedPassword, _ := HashPassword(user.password)

		user.password = string(hashedPassword)
	}
}
