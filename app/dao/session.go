package dao

import (
	"context"
	"crypto/rand"
	"math"
	"math/big"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/ProgrammingLab/toshokan/app/util"
)

type Session struct {
	SessionID uint `gorm:"primary_key"`
	CreatedAt time.Time
	ExpireAt  time.Time
	Token     string `gorm:"not null; unique_index"`
	User      User   `gorm:"foreignkey:UserID; association_foreignkey:UserID"`
	UserID    uint   `gorm:"not null"`
}

const SessionLifeTime = 48 * time.Hour

func NewSession(ctx context.Context, email, password string) (*Session, error) {
	u := &User{}
	res := db.Model(u).Where("email = ?", email).Scan(u)
	if err := res.Error; err != nil {
		if res.RecordNotFound() {
			return nil, bcrypt.ErrMismatchedHashAndPassword
		}
		return nil, err
	}

	if err := u.ComparePassword(password); err != nil {
		return nil, err
	}

	token, err := util.GenerateRandomString(32)
	if err != nil {
		return nil, err
	}

	max := big.NewInt(math.MaxInt32)
	var id uint
	tx := db.Begin()
	for {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		id = uint(n.Int64())
		if tx.Model(Session{}).Where("session_id = ?", id).Scan(&Session{}).RecordNotFound() {
			break
		}
	}

	s := &Session{
		SessionID: id,
		ExpireAt:  time.Now().Add(SessionLifeTime),
		Token:     token,
		UserID:    u.UserID,
	}

	if err := tx.Create(s).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	s.User = *u
	return s, nil
}
