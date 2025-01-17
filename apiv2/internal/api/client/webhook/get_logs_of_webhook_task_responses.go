// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ethenq/goharbor-client/v5/apiv2/model"
)

// GetLogsOfWebhookTaskReader is a Reader for the GetLogsOfWebhookTask structure.
type GetLogsOfWebhookTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsOfWebhookTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsOfWebhookTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsOfWebhookTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLogsOfWebhookTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLogsOfWebhookTaskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLogsOfWebhookTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLogsOfWebhookTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLogsOfWebhookTaskOK creates a GetLogsOfWebhookTaskOK with default headers values
func NewGetLogsOfWebhookTaskOK() *GetLogsOfWebhookTaskOK {
	return &GetLogsOfWebhookTaskOK{}
}

/*GetLogsOfWebhookTaskOK handles this case with default header values.

Get log success
*/
type GetLogsOfWebhookTaskOK struct {
	/*Content type of response
	 */
	ContentType string

	Payload string
}

func (o *GetLogsOfWebhookTaskOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskOK  %+v", 200, o.Payload)
}

func (o *GetLogsOfWebhookTaskOK) GetPayload() string {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Type
	o.ContentType = response.GetHeader("Content-Type")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsOfWebhookTaskBadRequest creates a GetLogsOfWebhookTaskBadRequest with default headers values
func NewGetLogsOfWebhookTaskBadRequest() *GetLogsOfWebhookTaskBadRequest {
	return &GetLogsOfWebhookTaskBadRequest{}
}

/*GetLogsOfWebhookTaskBadRequest handles this case with default header values.

Bad request
*/
type GetLogsOfWebhookTaskBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLogsOfWebhookTaskBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskBadRequest  %+v", 400, o.Payload)
}

func (o *GetLogsOfWebhookTaskBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsOfWebhookTaskUnauthorized creates a GetLogsOfWebhookTaskUnauthorized with default headers values
func NewGetLogsOfWebhookTaskUnauthorized() *GetLogsOfWebhookTaskUnauthorized {
	return &GetLogsOfWebhookTaskUnauthorized{}
}

/*GetLogsOfWebhookTaskUnauthorized handles this case with default header values.

Unauthorized
*/
type GetLogsOfWebhookTaskUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLogsOfWebhookTaskUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLogsOfWebhookTaskUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsOfWebhookTaskForbidden creates a GetLogsOfWebhookTaskForbidden with default headers values
func NewGetLogsOfWebhookTaskForbidden() *GetLogsOfWebhookTaskForbidden {
	return &GetLogsOfWebhookTaskForbidden{}
}

/*GetLogsOfWebhookTaskForbidden handles this case with default header values.

Forbidden
*/
type GetLogsOfWebhookTaskForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLogsOfWebhookTaskForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskForbidden  %+v", 403, o.Payload)
}

func (o *GetLogsOfWebhookTaskForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsOfWebhookTaskNotFound creates a GetLogsOfWebhookTaskNotFound with default headers values
func NewGetLogsOfWebhookTaskNotFound() *GetLogsOfWebhookTaskNotFound {
	return &GetLogsOfWebhookTaskNotFound{}
}

/*GetLogsOfWebhookTaskNotFound handles this case with default header values.

Not found
*/
type GetLogsOfWebhookTaskNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLogsOfWebhookTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetLogsOfWebhookTaskNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsOfWebhookTaskInternalServerError creates a GetLogsOfWebhookTaskInternalServerError with default headers values
func NewGetLogsOfWebhookTaskInternalServerError() *GetLogsOfWebhookTaskInternalServerError {
	return &GetLogsOfWebhookTaskInternalServerError{}
}

/*GetLogsOfWebhookTaskInternalServerError handles this case with default header values.

Internal server error
*/
type GetLogsOfWebhookTaskInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLogsOfWebhookTaskInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLogsOfWebhookTaskInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
