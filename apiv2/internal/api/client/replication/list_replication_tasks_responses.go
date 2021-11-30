// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// ListReplicationTasksReader is a Reader for the ListReplicationTasks structure.
type ListReplicationTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReplicationTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListReplicationTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListReplicationTasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListReplicationTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListReplicationTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListReplicationTasksOK creates a ListReplicationTasksOK with default headers values
func NewListReplicationTasksOK() *ListReplicationTasksOK {
	return &ListReplicationTasksOK{}
}

/*ListReplicationTasksOK handles this case with default header values.

Success
*/
type ListReplicationTasksOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of the resources
	 */
	XTotalCount int64

	Payload []*model.ReplicationTask
}

func (o *ListReplicationTasksOK) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks][%d] listReplicationTasksOK  %+v", 200, o.Payload)
}

func (o *ListReplicationTasksOK) GetPayload() []*model.ReplicationTask {
	return o.Payload
}

func (o *ListReplicationTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReplicationTasksUnauthorized creates a ListReplicationTasksUnauthorized with default headers values
func NewListReplicationTasksUnauthorized() *ListReplicationTasksUnauthorized {
	return &ListReplicationTasksUnauthorized{}
}

/*ListReplicationTasksUnauthorized handles this case with default header values.

Unauthorized
*/
type ListReplicationTasksUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListReplicationTasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks][%d] listReplicationTasksUnauthorized  %+v", 401, o.Payload)
}

func (o *ListReplicationTasksUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListReplicationTasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReplicationTasksForbidden creates a ListReplicationTasksForbidden with default headers values
func NewListReplicationTasksForbidden() *ListReplicationTasksForbidden {
	return &ListReplicationTasksForbidden{}
}

/*ListReplicationTasksForbidden handles this case with default header values.

Forbidden
*/
type ListReplicationTasksForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListReplicationTasksForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks][%d] listReplicationTasksForbidden  %+v", 403, o.Payload)
}

func (o *ListReplicationTasksForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListReplicationTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReplicationTasksInternalServerError creates a ListReplicationTasksInternalServerError with default headers values
func NewListReplicationTasksInternalServerError() *ListReplicationTasksInternalServerError {
	return &ListReplicationTasksInternalServerError{}
}

/*ListReplicationTasksInternalServerError handles this case with default header values.

Internal server error
*/
type ListReplicationTasksInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListReplicationTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks][%d] listReplicationTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *ListReplicationTasksInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListReplicationTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}