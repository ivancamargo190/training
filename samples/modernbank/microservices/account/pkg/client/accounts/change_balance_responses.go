// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeBalanceReader is a Reader for the ChangeBalance structure.
type ChangeBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeBalanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewChangeBalanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewChangeBalanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeBalanceOK creates a ChangeBalanceOK with default headers values
func NewChangeBalanceOK() *ChangeBalanceOK {
	return &ChangeBalanceOK{}
}

/*ChangeBalanceOK handles this case with default header values.

OK
*/
type ChangeBalanceOK struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *ChangeBalanceOK) Error() string {
	return fmt.Sprintf("[PUT /accounts/{number}/balance/{delta}][%d] changeBalanceOK ", 200)
}

func (o *ChangeBalanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}

// NewChangeBalanceNotFound creates a ChangeBalanceNotFound with default headers values
func NewChangeBalanceNotFound() *ChangeBalanceNotFound {
	return &ChangeBalanceNotFound{}
}

/*ChangeBalanceNotFound handles this case with default header values.

Account not found
*/
type ChangeBalanceNotFound struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *ChangeBalanceNotFound) Error() string {
	return fmt.Sprintf("[PUT /accounts/{number}/balance/{delta}][%d] changeBalanceNotFound ", 404)
}

func (o *ChangeBalanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}

// NewChangeBalanceInternalServerError creates a ChangeBalanceInternalServerError with default headers values
func NewChangeBalanceInternalServerError() *ChangeBalanceInternalServerError {
	return &ChangeBalanceInternalServerError{}
}

/*ChangeBalanceInternalServerError handles this case with default header values.

Internal server error
*/
type ChangeBalanceInternalServerError struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *ChangeBalanceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /accounts/{number}/balance/{delta}][%d] changeBalanceInternalServerError ", 500)
}

func (o *ChangeBalanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}
