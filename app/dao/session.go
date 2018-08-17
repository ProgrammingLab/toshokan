package dao

import (
	"context"
	"crypto/rand"
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/ProgrammingLab/toshokan/app/util"
)

type Session struct {
	SessionID  uint `gorm:"primary_key"`
	CreatedAt  time.Time
	ExpireAt   time.Time
	PrivateKey string `gorm:"not null; unique_index"`
	User       User   `gorm:"foreignkey:UserID; association_foreignkey:UserID"`
	UserID     uint   `gorm:"not null"`
}

const (
	SessionLifeTime = 48 * time.Hour
)

var (
	ErrBadSessionTokenFormat = errors.New("bad session token format")
	ErrInvalidSessionToken   = errors.New("invalid session token")
)

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

	key, err := util.GenerateRandomString(32)
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
		SessionID:  id,
		ExpireAt:   time.Now().Add(SessionLifeTime),
		PrivateKey: key,
		UserID:     u.UserID,
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

func GetSession(ctx context.Context, token string) (*Session, error) {
	i := strings.Index(token, "_")
	if i == -1 || len(token) < i+2 {
		return nil, ErrBadSessionTokenFormat
	}

	rawID := token[:i]
	key := token[i+1:]
	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return nil, ErrBadSessionTokenFormat
	}

	s := &Session{}
	res := db.Model(Session{}).Where("session_id = ?", id).Scan(s)
	if err := res.Error; err != nil {
		if res.RecordNotFound() {
			return nil, ErrInvalidSessionToken
		}
		return nil, err
	}

	if s.PrivateKey != key || time.Now().After(s.ExpireAt) {
		return nil, ErrInvalidSessionToken
	}

	s.PrivateKey = ""

	return s, nil
}

func (s *Session) Delete() error {
	return db.Where("session_id = ?", s.SessionID).Delete(Session{}).Error
}
