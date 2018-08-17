package controllers

import (
	"github.com/go-openapi/errors"
)

var (
	errInternalServerError = errors.New(500, "internal server error")
	errInvalidToken        = errors.New(401, "invalid token")
	errBadTokenFormat      = errors.New(400, "bad token format")
)
