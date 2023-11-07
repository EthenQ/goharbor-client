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

// DeleteWebhookPolicyOfProjectReader is a Reader for the DeleteWebhookPolicyOfProject structure.
type DeleteWebhookPolicyOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWebhookPolicyOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteWebhookPolicyOfProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWebhookPolicyOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWebhookPolicyOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWebhookPolicyOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWebhookPolicyOfProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWebhookPolicyOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWebhookPolicyOfProjectOK creates a DeleteWebhookPolicyOfProjectOK with default headers values
func NewDeleteWebhookPolicyOfProjectOK() *DeleteWebhookPolicyOfProjectOK {
	return &DeleteWebhookPolicyOfProjectOK{}
}

/*DeleteWebhookPolicyOfProjectOK handles this case with default header values.

Success
*/
type DeleteWebhookPolicyOfProjectOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteWebhookPolicyOfProjectOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] deleteWebhookPolicyOfProjectOK ", 200)
}

func (o *DeleteWebhookPolicyOfProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewDeleteWebhookPolicyOfProjectBadRequest creates a DeleteWebhookPolicyOfProjectBadRequest with default headers values
func NewDeleteWebhookPolicyOfProjectBadRequest() *DeleteWebhookPolicyOfProjectBadRequest {
	return &DeleteWebhookPolicyOfProjectBadRequest{}
}

/*DeleteWebhookPolicyOfProjectBadRequest handles this case with default header values.

Bad request
*/
type DeleteWebhookPolicyOfProjectBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteWebhookPolicyOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] deleteWebhookPolicyOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWebhookPolicyOfProjectBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteWebhookPolicyOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebhookPolicyOfProjectUnauthorized creates a DeleteWebhookPolicyOfProjectUnauthorized with default headers values
func NewDeleteWebhookPolicyOfProjectUnauthorized() *DeleteWebhookPolicyOfProjectUnauthorized {
	return &DeleteWebhookPolicyOfProjectUnauthorized{}
}

/*DeleteWebhookPolicyOfProjectUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteWebhookPolicyOfProjectUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteWebhookPolicyOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] deleteWebhookPolicyOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWebhookPolicyOfProjectUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteWebhookPolicyOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebhookPolicyOfProjectForbidden creates a DeleteWebhookPolicyOfProjectForbidden with default headers values
func NewDeleteWebhookPolicyOfProjectForbidden() *DeleteWebhookPolicyOfProjectForbidden {
	return &DeleteWebhookPolicyOfProjectForbidden{}
}

/*DeleteWebhookPolicyOfProjectForbidden handles this case with default header values.

Forbidden
*/
type DeleteWebhookPolicyOfProjectForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteWebhookPolicyOfProjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] deleteWebhookPolicyOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWebhookPolicyOfProjectForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteWebhookPolicyOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebhookPolicyOfProjectNotFound creates a DeleteWebhookPolicyOfProjectNotFound with default headers values
func NewDeleteWebhookPolicyOfProjectNotFound() *DeleteWebhookPolicyOfProjectNotFound {
	return &DeleteWebhookPolicyOfProjectNotFound{}
}

/*DeleteWebhookPolicyOfProjectNotFound handles this case with default header values.

Not found
*/
type DeleteWebhookPolicyOfProjectNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteWebhookPolicyOfProjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] deleteWebhookPolicyOfProjectNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWebhookPolicyOfProjectNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteWebhookPolicyOfProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebhookPolicyOfProjectInternalServerError creates a DeleteWebhookPolicyOfProjectInternalServerError with default headers values
func NewDeleteWebhookPolicyOfProjectInternalServerError() *DeleteWebhookPolicyOfProjectInternalServerError {
	return &DeleteWebhookPolicyOfProjectInternalServerError{}
}

/*DeleteWebhookPolicyOfProjectInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteWebhookPolicyOfProjectInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteWebhookPolicyOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] deleteWebhookPolicyOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWebhookPolicyOfProjectInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteWebhookPolicyOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
