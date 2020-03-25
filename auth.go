package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

// JWT Model
type JWT struct {
	Token string `json:"token"`
}

// HashPassword - this hashes the user's raw password
func HashPassword(password string) (string, error) {
	hashingCost := 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashingCost)
	if err != nil {
		fmt.Println("Failed to encrypt", err)
	}
	return string(hashedPassword), err
}

func (user UserDTO) checkPassword(password string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))

	if result == bcrypt.ErrHashTooShort {
		fmt.Println("Error:", bcrypt.ErrHashTooShort)
		return false
	}

	if result == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("Error:", bcrypt.ErrMismatchedHashAndPassword)
		return false
	}

	return true
}

// Login - user login handler
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login called")
}

// Signup - user signup handler
func Signup(w http.ResponseWriter, r *http.Request) {
	hashedPassword, _ := HashPassword("password")
	userData := MakeDTO(UserDTO{
		id:       uuid.New().String(),
		password: hashedPassword,
		email:    "jsamchineme@gmail.com",
	})

	_, err := UserRepo.CreateRecord(userData)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := UserModel.getTableData()
	fmt.Println("Users\n", rows)
}

// TokenVerifyMiddleware handles token verification
// for JWT routes protection
func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("token middleware called")
	return nil
}
