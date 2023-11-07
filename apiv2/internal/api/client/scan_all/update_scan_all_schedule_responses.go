// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ethenq/goharbor-client/v5/apiv2/model"
)

// UpdateScanAllScheduleReader is a Reader for the UpdateScanAllSchedule structure.
type UpdateScanAllScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateScanAllScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateScanAllScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateScanAllScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateScanAllScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateScanAllScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewUpdateScanAllSchedulePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateScanAllScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateScanAllScheduleOK creates a UpdateScanAllScheduleOK with default headers values
func NewUpdateScanAllScheduleOK() *UpdateScanAllScheduleOK {
	return &UpdateScanAllScheduleOK{}
}

/*UpdateScanAllScheduleOK handles this case with default header values.

Success
*/
type UpdateScanAllScheduleOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateScanAllScheduleOK) Error() string {
	return fmt.Sprintf("[PUT /system/scanAll/schedule][%d] updateScanAllScheduleOK ", 200)
}

func (o *UpdateScanAllScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewUpdateScanAllScheduleBadRequest creates a UpdateScanAllScheduleBadRequest with default headers values
func NewUpdateScanAllScheduleBadRequest() *UpdateScanAllScheduleBadRequest {
	return &UpdateScanAllScheduleBadRequest{}
}

/*UpdateScanAllScheduleBadRequest handles this case with default header values.

Bad request
*/
type UpdateScanAllScheduleBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateScanAllScheduleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /system/scanAll/schedule][%d] updateScanAllScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateScanAllScheduleBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateScanAllScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScanAllScheduleUnauthorized creates a UpdateScanAllScheduleUnauthorized with default headers values
func NewUpdateScanAllScheduleUnauthorized() *UpdateScanAllScheduleUnauthorized {
	return &UpdateScanAllScheduleUnauthorized{}
}

/*UpdateScanAllScheduleUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateScanAllScheduleUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateScanAllScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /system/scanAll/schedule][%d] updateScanAllScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateScanAllScheduleUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateScanAllScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScanAllScheduleForbidden creates a UpdateScanAllScheduleForbidden with default headers values
func NewUpdateScanAllScheduleForbidden() *UpdateScanAllScheduleForbidden {
	return &UpdateScanAllScheduleForbidden{}
}

/*UpdateScanAllScheduleForbidden handles this case with default header values.

Forbidden
*/
type UpdateScanAllScheduleForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateScanAllScheduleForbidden) Error() string {
	return fmt.Sprintf("[PUT /system/scanAll/schedule][%d] updateScanAllScheduleForbidden  %+v", 403, o.Payload)
}

func (o *UpdateScanAllScheduleForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateScanAllScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScanAllSchedulePreconditionFailed creates a UpdateScanAllSchedulePreconditionFailed with default headers values
func NewUpdateScanAllSchedulePreconditionFailed() *UpdateScanAllSchedulePreconditionFailed {
	return &UpdateScanAllSchedulePreconditionFailed{}
}

/*UpdateScanAllSchedulePreconditionFailed handles this case with default header values.

Precondition failed
*/
type UpdateScanAllSchedulePreconditionFailed struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateScanAllSchedulePreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /system/scanAll/schedule][%d] updateScanAllSchedulePreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateScanAllSchedulePreconditionFailed) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateScanAllSchedulePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScanAllScheduleInternalServerError creates a UpdateScanAllScheduleInternalServerError with default headers values
func NewUpdateScanAllScheduleInternalServerError() *UpdateScanAllScheduleInternalServerError {
	return &UpdateScanAllScheduleInternalServerError{}
}

/*UpdateScanAllScheduleInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateScanAllScheduleInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateScanAllScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /system/scanAll/schedule][%d] updateScanAllScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateScanAllScheduleInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateScanAllScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
