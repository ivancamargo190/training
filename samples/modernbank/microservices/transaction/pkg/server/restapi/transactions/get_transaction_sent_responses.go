// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	model "github.com/tetrateio/training/samples/modernbank/microservices/transaction/pkg/model"
)

// GetTransactionSentOKCode is the HTTP code returned for type GetTransactionSentOK
const GetTransactionSentOKCode int = 200

/*GetTransactionSentOK Success!

swagger:response getTransactionSentOK
*/
type GetTransactionSentOK struct {

	/*
	  In: Body
	*/
	Payload *model.Transaction `json:"body,omitempty"`
}

// NewGetTransactionSentOK creates GetTransactionSentOK with default headers values
func NewGetTransactionSentOK() *GetTransactionSentOK {

	return &GetTransactionSentOK{}
}

// WithPayload adds the payload to the get transaction sent o k response
func (o *GetTransactionSentOK) WithPayload(payload *model.Transaction) *GetTransactionSentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transaction sent o k response
func (o *GetTransactionSentOK) SetPayload(payload *model.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionSentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTransactionSentNotFoundCode is the HTTP code returned for type GetTransactionSentNotFound
const GetTransactionSentNotFoundCode int = 404

/*GetTransactionSentNotFound Transaction not found

swagger:response getTransactionSentNotFound
*/
type GetTransactionSentNotFound struct {
}

// NewGetTransactionSentNotFound creates GetTransactionSentNotFound with default headers values
func NewGetTransactionSentNotFound() *GetTransactionSentNotFound {

	return &GetTransactionSentNotFound{}
}

// WriteResponse to the client
func (o *GetTransactionSentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetTransactionSentInternalServerErrorCode is the HTTP code returned for type GetTransactionSentInternalServerError
const GetTransactionSentInternalServerErrorCode int = 500

/*GetTransactionSentInternalServerError Internal server error

swagger:response getTransactionSentInternalServerError
*/
type GetTransactionSentInternalServerError struct {
}

// NewGetTransactionSentInternalServerError creates GetTransactionSentInternalServerError with default headers values
func NewGetTransactionSentInternalServerError() *GetTransactionSentInternalServerError {

	return &GetTransactionSentInternalServerError{}
}

// WriteResponse to the client
func (o *GetTransactionSentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
