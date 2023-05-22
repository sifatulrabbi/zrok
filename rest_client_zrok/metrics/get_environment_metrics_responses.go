// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/zrok/rest_model_zrok"
)

// GetEnvironmentMetricsReader is a Reader for the GetEnvironmentMetrics structure.
type GetEnvironmentMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvironmentMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnvironmentMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEnvironmentMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEnvironmentMetricsOK creates a GetEnvironmentMetricsOK with default headers values
func NewGetEnvironmentMetricsOK() *GetEnvironmentMetricsOK {
	return &GetEnvironmentMetricsOK{}
}

/*
GetEnvironmentMetricsOK describes a response with status code 200, with default header values.

environment metrics
*/
type GetEnvironmentMetricsOK struct {
	Payload *rest_model_zrok.Metrics
}

// IsSuccess returns true when this get environment metrics o k response has a 2xx status code
func (o *GetEnvironmentMetricsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get environment metrics o k response has a 3xx status code
func (o *GetEnvironmentMetricsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment metrics o k response has a 4xx status code
func (o *GetEnvironmentMetricsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get environment metrics o k response has a 5xx status code
func (o *GetEnvironmentMetricsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment metrics o k response a status code equal to that given
func (o *GetEnvironmentMetricsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEnvironmentMetricsOK) Error() string {
	return fmt.Sprintf("[GET /metrics/environment/{envId}][%d] getEnvironmentMetricsOK  %+v", 200, o.Payload)
}

func (o *GetEnvironmentMetricsOK) String() string {
	return fmt.Sprintf("[GET /metrics/environment/{envId}][%d] getEnvironmentMetricsOK  %+v", 200, o.Payload)
}

func (o *GetEnvironmentMetricsOK) GetPayload() *rest_model_zrok.Metrics {
	return o.Payload
}

func (o *GetEnvironmentMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model_zrok.Metrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnvironmentMetricsUnauthorized creates a GetEnvironmentMetricsUnauthorized with default headers values
func NewGetEnvironmentMetricsUnauthorized() *GetEnvironmentMetricsUnauthorized {
	return &GetEnvironmentMetricsUnauthorized{}
}

/*
GetEnvironmentMetricsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type GetEnvironmentMetricsUnauthorized struct {
}

// IsSuccess returns true when this get environment metrics unauthorized response has a 2xx status code
func (o *GetEnvironmentMetricsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get environment metrics unauthorized response has a 3xx status code
func (o *GetEnvironmentMetricsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment metrics unauthorized response has a 4xx status code
func (o *GetEnvironmentMetricsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get environment metrics unauthorized response has a 5xx status code
func (o *GetEnvironmentMetricsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment metrics unauthorized response a status code equal to that given
func (o *GetEnvironmentMetricsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetEnvironmentMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /metrics/environment/{envId}][%d] getEnvironmentMetricsUnauthorized ", 401)
}

func (o *GetEnvironmentMetricsUnauthorized) String() string {
	return fmt.Sprintf("[GET /metrics/environment/{envId}][%d] getEnvironmentMetricsUnauthorized ", 401)
}

func (o *GetEnvironmentMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
