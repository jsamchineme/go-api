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

var users = []UserDTO{}

type fieldList struct {
	id       string
	email    string
	password string
}

var tableFields = fieldList{
	id:       "id",
	email:    "email",
	password: "password",
}

// UserRepo - for data store access
var UserRepo = &User{}

func (u User) updateDTO(record UserDTO, data UserDTO) UserDTO {
	newData := UserDTO{}
	fields := []string{"email", "password"}

	for _, field := range fields {
		switch field {
		case "email":
			if data.email != "" {
				newData.email = data.email
			} else {
				newData.email = record.email
			}
		case "password":
			if data.password != "" {
				newData.password = data.password
			} else {
				newData.password = record.password
			}
		}
	}

	return newData
}

func (u User) getTableData() []UserDTO {
	return users
}

func (User) getTableName() string {
	return "users"
}

func (User) setTableData(data []UserDTO) {
	users = data
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

		d := UserDTO{
			id:       user.id,
			email:    user.email,
			password: hashedPassword,
		}

		UserRepo.createRecord(d)
		UserRepo.updateRecord(d)
	}
}

// CreateRecord is used to insert a new row into a table
func (u User) createRecord(d UserDTO) (UserDTO, error) {
	rows := u.getTableData()
	rows = append(rows, d)
	u.setTableData(rows)

	return d, nil
}

// UpdateRecord update a record
func (u User) updateRecord(d UserDTO) (UserDTO, error) {
	rows := []UserDTO{}
	for _, record := range u.getTableData() {
		if d.id == record.id {
			updatedDTO := u.updateDTO(record, d)
			rows = append(rows, updatedDTO)
			continue
		}
		rows = append(rows, d)
	}
	u.setTableData(rows)
	return d, nil
}
