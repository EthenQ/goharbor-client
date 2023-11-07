// Code generated by go-swagger; DO NOT EDIT.

package purge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ethenq/goharbor-client/v5/apiv2/model"
)

// GetPurgeJobReader is a Reader for the GetPurgeJob structure.
type GetPurgeJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPurgeJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPurgeJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPurgeJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPurgeJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPurgeJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPurgeJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPurgeJobOK creates a GetPurgeJobOK with default headers values
func NewGetPurgeJobOK() *GetPurgeJobOK {
	return &GetPurgeJobOK{}
}

/*GetPurgeJobOK handles this case with default header values.

Get purge job results successfully.
*/
type GetPurgeJobOK struct {
	Payload *model.ExecHistory
}

func (o *GetPurgeJobOK) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/{purge_id}][%d] getPurgeJobOK  %+v", 200, o.Payload)
}

func (o *GetPurgeJobOK) GetPayload() *model.ExecHistory {
	return o.Payload
}

func (o *GetPurgeJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ExecHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeJobUnauthorized creates a GetPurgeJobUnauthorized with default headers values
func NewGetPurgeJobUnauthorized() *GetPurgeJobUnauthorized {
	return &GetPurgeJobUnauthorized{}
}

/*GetPurgeJobUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPurgeJobUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/{purge_id}][%d] getPurgeJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPurgeJobUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeJobForbidden creates a GetPurgeJobForbidden with default headers values
func NewGetPurgeJobForbidden() *GetPurgeJobForbidden {
	return &GetPurgeJobForbidden{}
}

/*GetPurgeJobForbidden handles this case with default header values.

Forbidden
*/
type GetPurgeJobForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeJobForbidden) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/{purge_id}][%d] getPurgeJobForbidden  %+v", 403, o.Payload)
}

func (o *GetPurgeJobForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeJobNotFound creates a GetPurgeJobNotFound with default headers values
func NewGetPurgeJobNotFound() *GetPurgeJobNotFound {
	return &GetPurgeJobNotFound{}
}

/*GetPurgeJobNotFound handles this case with default header values.

Not found
*/
type GetPurgeJobNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeJobNotFound) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/{purge_id}][%d] getPurgeJobNotFound  %+v", 404, o.Payload)
}

func (o *GetPurgeJobNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeJobInternalServerError creates a GetPurgeJobInternalServerError with default headers values
func NewGetPurgeJobInternalServerError() *GetPurgeJobInternalServerError {
	return &GetPurgeJobInternalServerError{}
}

/*GetPurgeJobInternalServerError handles this case with default header values.

Internal server error
*/
type GetPurgeJobInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/{purge_id}][%d] getPurgeJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPurgeJobInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
