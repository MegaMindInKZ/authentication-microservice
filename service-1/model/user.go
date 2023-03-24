package models

import (
	"github.com/MegaMindInKZ/authentication-microservice.git/service-1/internalErrors"
	"github.com/MegaMindInKZ/authentication-microservice.git/service-1/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func LoginCheck(username string, password string) (string, error) {
	var err error
	
	u := User{}

	if err := DB.Model(User{}).Where("username = ?", username).Take(&u).Error; err != nil {
		return "", internalErrors.IncorrectPasswordOrUsernameError
	}

	if !VerifyPassword(u.Password, password) {
		return "", internalErrors.IncorrectPasswordOrUsernameError

	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", internalErrors.CannotCreateTokenError
	}

	return token, nil
}

func VerifyPassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}

func (u *User) AfterCreate(tx *gorm.DB) error {
	if u.ID == 1 {

	}
	return nil
}
