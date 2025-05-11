package auth

import (
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/siddhantmadhur/ocelot-media-server/internal/storage"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username          string `json:"username" gorm:"unique;not null"`
	DisplayName       string `json:"display_name" gorm:"not null"`
	EncryptedPassword string `json:"encrypted_password" gorm:"not null"`
	Permission        int    `json:"permission" gorm:"not null"`
}

func UpdateModels(tx *gorm.DB) error {
	log.Printf("[AUTH] Auto migrating models to database...\n")

	userMigrateErr := tx.AutoMigrate(&User{})
	sessionMigrateErr := tx.AutoMigrate(&User{})

	if userMigrateErr != nil || sessionMigrateErr != nil {
		log.Printf("[ERROR] Error in migration")
		return errors.New("Error in migration")
	}

	return nil
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

// Return a string of the displayname for every user
func GetAllUsers() ([]string, error) {
	var usernames []string

	tx, err := storage.GetConnection()
	defer storage.CloseConnection(tx)

	if err != nil {
		return nil, err
	}

	res := tx.Raw("SELECT display_name FROM users").Scan(&usernames)
	if res.Error != nil {
		return nil, res.Error
	}

	return usernames, nil

}

func LoginUser(username string, password string) (*jwt.Token, error) {
	var u User
	tx, err := storage.GetConnection()
	defer storage.CloseConnection(tx)

	if err != nil {
		return nil, err
	}

	res := tx.Raw("SELECT * FROM users WHERE username = ? LIMIT 1", username).Scan(&u)
	if res.Error != nil {
		return nil, res.Error
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	if err != nil {
		return nil, errors.New("Password does not match!")
	}

	return nil, nil
}
