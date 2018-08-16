package dao

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID         uint `gorm:"primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Name           string `gorm:"not null; unique_index"`
	DisplayName    string `gorm:"not null"`
	Email          string `gorm:"not null; unique_index"`
	PasswordDigest string `gorm:"not null"`
	IsAdmin        bool   `gorm:"not null"`
}

func NewUser(ctx context.Context, user User) error {
	tx := db.Begin()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func FindUserByName(ctx context.Context, name string) (User, error) {
	u := User{}
	if err := db.Model(u).Where("name = ?", name).Scan(&u).Error; err != nil {
		return User{}, err
	}
	return u, nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
}
