// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostSessionsCreatedCode is the HTTP code returned for type PostSessionsCreated
const PostSessionsCreatedCode int = 201

/*PostSessionsCreated Created

swagger:response postSessionsCreated
*/
type PostSessionsCreated struct {

	/*
	  In: Body
	*/
	Payload *PostSessionsCreatedBody `json:"body,omitempty"`
}

// NewPostSessionsCreated creates PostSessionsCreated with default headers values
func NewPostSessionsCreated() *PostSessionsCreated {

	return &PostSessionsCreated{}
}

// WithPayload adds the payload to the post sessions created response
func (o *PostSessionsCreated) WithPayload(payload *PostSessionsCreatedBody) *PostSessionsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post sessions created response
func (o *PostSessionsCreated) SetPayload(payload *PostSessionsCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSessionsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostSessionsUnauthorizedCode is the HTTP code returned for type PostSessionsUnauthorized
const PostSessionsUnauthorizedCode int = 401

/*PostSessionsUnauthorized Unauthorized

swagger:response postSessionsUnauthorized
*/
type PostSessionsUnauthorized struct {
}

// NewPostSessionsUnauthorized creates PostSessionsUnauthorized with default headers values
func NewPostSessionsUnauthorized() *PostSessionsUnauthorized {

	return &PostSessionsUnauthorized{}
}

// WriteResponse to the client
func (o *PostSessionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
