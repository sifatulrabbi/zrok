// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/zrok/rest_model_zrok"
)

// GetAccountDetailReader is a Reader for the GetAccountDetail structure.
type GetAccountDetailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountDetailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountDetailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAccountDetailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountDetailOK creates a GetAccountDetailOK with default headers values
func NewGetAccountDetailOK() *GetAccountDetailOK {
	return &GetAccountDetailOK{}
}

/*
GetAccountDetailOK describes a response with status code 200, with default header values.

ok
*/
type GetAccountDetailOK struct {
	Payload rest_model_zrok.Environments
}

// IsSuccess returns true when this get account detail o k response has a 2xx status code
func (o *GetAccountDetailOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get account detail o k response has a 3xx status code
func (o *GetAccountDetailOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account detail o k response has a 4xx status code
func (o *GetAccountDetailOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get account detail o k response has a 5xx status code
func (o *GetAccountDetailOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get account detail o k response a status code equal to that given
func (o *GetAccountDetailOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAccountDetailOK) Error() string {
	return fmt.Sprintf("[GET /detail/account][%d] getAccountDetailOK  %+v", 200, o.Payload)
}

func (o *GetAccountDetailOK) String() string {
	return fmt.Sprintf("[GET /detail/account][%d] getAccountDetailOK  %+v", 200, o.Payload)
}

func (o *GetAccountDetailOK) GetPayload() rest_model_zrok.Environments {
	return o.Payload
}

func (o *GetAccountDetailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountDetailInternalServerError creates a GetAccountDetailInternalServerError with default headers values
func NewGetAccountDetailInternalServerError() *GetAccountDetailInternalServerError {
	return &GetAccountDetailInternalServerError{}
}

/*
GetAccountDetailInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type GetAccountDetailInternalServerError struct {
}

// IsSuccess returns true when this get account detail internal server error response has a 2xx status code
func (o *GetAccountDetailInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get account detail internal server error response has a 3xx status code
func (o *GetAccountDetailInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account detail internal server error response has a 4xx status code
func (o *GetAccountDetailInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get account detail internal server error response has a 5xx status code
func (o *GetAccountDetailInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get account detail internal server error response a status code equal to that given
func (o *GetAccountDetailInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAccountDetailInternalServerError) Error() string {
	return fmt.Sprintf("[GET /detail/account][%d] getAccountDetailInternalServerError ", 500)
}

func (o *GetAccountDetailInternalServerError) String() string {
	return fmt.Sprintf("[GET /detail/account][%d] getAccountDetailInternalServerError ", 500)
}

func (o *GetAccountDetailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
