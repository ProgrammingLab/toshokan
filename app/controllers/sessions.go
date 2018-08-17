package controllers

import (
	"context"
	"log"
	"strconv"

	"github.com/ProgrammingLab/toshokan/app/conv"
	"github.com/ProgrammingLab/toshokan/app/dao"
	"github.com/ProgrammingLab/toshokan/restapi/operations/sessions"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(token string) (interface{}, error) {
	if 128 < len(token) {
		return nil, errBadTokenFormat
	}
	s, err := dao.GetSession(context.Background(), token)
	switch err {
	case dao.ErrBadSessionTokenFormat:
		return nil, errBadTokenFormat
	case dao.ErrInvalidSessionToken:
		return nil, errInvalidToken
	case nil:
		return s, nil
	default:
		return nil, errInternalServerError
	}
}

func Login(params sessions.LoginParams) middleware.Responder {
	req := params.EmailAndPassword
	s, err := dao.NewSession(params.HTTPRequest.Context(), *req.Email, *req.Password)

	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return sessions.NewLoginUnauthorized()
		}

		log.Print(err)
		return InteranlServerError{}
	}

	b := &sessions.LoginCreatedBody{
		Token: strconv.FormatInt(int64(s.SessionID), 10) + "_" + s.PrivateKey,
		User:  conv.UserDAOToUserModel(&s.User, true),
	}
	return sessions.NewLoginCreated().WithPayload(b)
}

func Logout(params sessions.LogoutParams, session interface{}) middleware.Responder {
	s := session.(*dao.Session)
	if err := s.Delete(); err != nil {
		log.Print(err)
		return InteranlServerError{}
	}
	return sessions.NewLogoutNoContent()
}
