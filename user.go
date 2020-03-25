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

var users = &[]ModelDTO{}

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
	return *users
}

// Model implementation MUST have this
func (User) getTableName() string {
	return "users"
}

func makeDTO(t UserDTO) ModelDTO {
	return ModelDTO{t}
}

func (User) initialiseTable() {

	dummy := []UserDTO{
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

	for _, user := range dummy {
		// There's the need to wrap this block within the type assertion for User
		// because each user is seen as an Entity type
		// Wrapping the code in the assertion block makes it so that the "u" variable refers to a User type
		// which will then make field password, email accessible
		hashedPassword, _ := HashPassword(user.password)
		userData := MakeDTO(UserDTO{
			id:       uuid.New().String(),
			password: hashedPassword,
			email:    "jsamchineme@gmail.com",
		})
		UserRepo.CreateRecord(userData)
	}
}
