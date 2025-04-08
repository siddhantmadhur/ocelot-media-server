package auth

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id                uint   `json:"id" gorm:"primaryKey"`
	Username          string `json:"username" gorm:"unique;not null"`
	DisplayName       string `json:"display_name" gorm:"not null"`
	EncryptedPassword string `json:"-" gorm:"not null"`
	Permission        int    `json:"permission" gorm:"not null"`
}

func UpdateModels(tx *gorm.DB) error {
	log.Printf("[AUTH] Auto migrating models to database...\n")
	err := tx.AutoMigrate(&User{})
	return err
}

func CreateUser(username string, displayName string, password string) (User, error) {
	var u User

	if len(username) <= 3 || len(username) >= 14 {
		return u, errors.New("Username is not within 3 and 14 characters")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return u, err
	}

	u.EncryptedPassword = string(encryptedPassword)
	u.DisplayName = displayName
	u.Username = username

	return u, err
}
