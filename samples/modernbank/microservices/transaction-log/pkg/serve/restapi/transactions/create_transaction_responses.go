// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	model "github.com/tetrateio/training/samples/modernbank/microservices/transaction-log/pkg/model"
)

// CreateTransactionCreatedCode is the HTTP code returned for type CreateTransactionCreated
const CreateTransactionCreatedCode int = 201

/*CreateTransactionCreated Created

swagger:response createTransactionCreated
*/
type CreateTransactionCreated struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`

	/*
	  In: Body
	*/
	Payload *model.Transaction `json:"body,omitempty"`
}

// NewCreateTransactionCreated creates CreateTransactionCreated with default headers values
func NewCreateTransactionCreated() *CreateTransactionCreated {

	return &CreateTransactionCreated{}
}

// WithVersion adds the version to the create transaction created response
func (o *CreateTransactionCreated) WithVersion(version string) *CreateTransactionCreated {
	o.Version = version
	return o
}

// SetVersion sets the version to the create transaction created response
func (o *CreateTransactionCreated) SetVersion(version string) {
	o.Version = version
}

// WithPayload adds the payload to the create transaction created response
func (o *CreateTransactionCreated) WithPayload(payload *model.Transaction) *CreateTransactionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create transaction created response
func (o *CreateTransactionCreated) SetPayload(payload *model.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTransactionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTransactionBadRequestCode is the HTTP code returned for type CreateTransactionBadRequest
const CreateTransactionBadRequestCode int = 400

/*CreateTransactionBadRequest Nice try! You can't send negative amounts...

swagger:response createTransactionBadRequest
*/
type CreateTransactionBadRequest struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`
}

// NewCreateTransactionBadRequest creates CreateTransactionBadRequest with default headers values
func NewCreateTransactionBadRequest() *CreateTransactionBadRequest {

	return &CreateTransactionBadRequest{}
}

// WithVersion adds the version to the create transaction bad request response
func (o *CreateTransactionBadRequest) WithVersion(version string) *CreateTransactionBadRequest {
	o.Version = version
	return o
}

// SetVersion sets the version to the create transaction bad request response
func (o *CreateTransactionBadRequest) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *CreateTransactionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateTransactionInternalServerErrorCode is the HTTP code returned for type CreateTransactionInternalServerError
const CreateTransactionInternalServerErrorCode int = 500

/*CreateTransactionInternalServerError Internal server error

swagger:response createTransactionInternalServerError
*/
type CreateTransactionInternalServerError struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`
}

// NewCreateTransactionInternalServerError creates CreateTransactionInternalServerError with default headers values
func NewCreateTransactionInternalServerError() *CreateTransactionInternalServerError {

	return &CreateTransactionInternalServerError{}
}

// WithVersion adds the version to the create transaction internal server error response
func (o *CreateTransactionInternalServerError) WithVersion(version string) *CreateTransactionInternalServerError {
	o.Version = version
	return o
}

// SetVersion sets the version to the create transaction internal server error response
func (o *CreateTransactionInternalServerError) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *CreateTransactionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}