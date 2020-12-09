package models

import (
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Username    string `json:"username"`
	LastSession string `json:"last_session"`
}

// CheckUserExist ...
func CheckUserExist(username string) (bool, error) {
	var user User

	err := db.Select("username").Where("username = ? ", username).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

// CreateUser ...
func CreateUser(username string) error {
	user := User{
		Username: username,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
