package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        string "json:`gorm:'primaryKey'`"
	FirstName string "json:'first_name'"
	LastName  string "json:'last_name'"
	Email     string "json:email"
	Phone     string "json:phone"
	Password  []byte "json:'-'"
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.Id = uuid.NewString()
	return
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
