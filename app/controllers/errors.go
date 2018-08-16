package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

type InteranlServerError struct{}

func (e InteranlServerError) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	w.WriteHeader(http.StatusInternalServerError)
}
