package auth

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	Id                uint   `json:"id" gorm:"primaryKey"`
	Username          string `json:"username" gorm:"unique;not null"`
	EncryptedPassword string `json:"-" gorm:"not null"`
}

func UpdateModels(tx *gorm.DB) error {
	log.Printf("[AUTH] Auto migrating models to database...\n")
	err := tx.AutoMigrate(&User{})
	return err
}
