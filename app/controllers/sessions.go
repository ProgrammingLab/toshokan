package controllers

import (
	"log"
	"strconv"

	"github.com/ProgrammingLab/toshokan/app/conv"
	"github.com/ProgrammingLab/toshokan/app/dao"
	"github.com/ProgrammingLab/toshokan/restapi/operations/sessions"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

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
		Token: strconv.FormatInt(int64(s.SessionID), 10) + "_" + s.Token,
		User:  conv.UserDAOToUserModel(&s.User, true),
	}
	return sessions.NewLoginCreated().WithPayload(b)
}

func Logout(params sessions.LogoutParams) middleware.Responder {
	return middleware.NotImplemented("operation sessions.Logout has not yet been implemented")
}
