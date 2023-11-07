// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ethenq/goharbor-client/v5/apiv2/model"
)

// ListExecutionsOfWebhookPolicyReader is a Reader for the ListExecutionsOfWebhookPolicy structure.
type ListExecutionsOfWebhookPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListExecutionsOfWebhookPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListExecutionsOfWebhookPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListExecutionsOfWebhookPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListExecutionsOfWebhookPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListExecutionsOfWebhookPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListExecutionsOfWebhookPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListExecutionsOfWebhookPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListExecutionsOfWebhookPolicyOK creates a ListExecutionsOfWebhookPolicyOK with default headers values
func NewListExecutionsOfWebhookPolicyOK() *ListExecutionsOfWebhookPolicyOK {
	return &ListExecutionsOfWebhookPolicyOK{}
}

/*ListExecutionsOfWebhookPolicyOK handles this case with default header values.

List webhook executions success
*/
type ListExecutionsOfWebhookPolicyOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of executions
	 */
	XTotalCount int64

	Payload []*model.Execution
}

func (o *ListExecutionsOfWebhookPolicyOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions][%d] listExecutionsOfWebhookPolicyOK  %+v", 200, o.Payload)
}

func (o *ListExecutionsOfWebhookPolicyOK) GetPayload() []*model.Execution {
	return o.Payload
}

func (o *ListExecutionsOfWebhookPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListExecutionsOfWebhookPolicyBadRequest creates a ListExecutionsOfWebhookPolicyBadRequest with default headers values
func NewListExecutionsOfWebhookPolicyBadRequest() *ListExecutionsOfWebhookPolicyBadRequest {
	return &ListExecutionsOfWebhookPolicyBadRequest{}
}

/*ListExecutionsOfWebhookPolicyBadRequest handles this case with default header values.

Bad request
*/
type ListExecutionsOfWebhookPolicyBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListExecutionsOfWebhookPolicyBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions][%d] listExecutionsOfWebhookPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *ListExecutionsOfWebhookPolicyBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListExecutionsOfWebhookPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExecutionsOfWebhookPolicyUnauthorized creates a ListExecutionsOfWebhookPolicyUnauthorized with default headers values
func NewListExecutionsOfWebhookPolicyUnauthorized() *ListExecutionsOfWebhookPolicyUnauthorized {
	return &ListExecutionsOfWebhookPolicyUnauthorized{}
}

/*ListExecutionsOfWebhookPolicyUnauthorized handles this case with default header values.

Unauthorized
*/
type ListExecutionsOfWebhookPolicyUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListExecutionsOfWebhookPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions][%d] listExecutionsOfWebhookPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *ListExecutionsOfWebhookPolicyUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListExecutionsOfWebhookPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExecutionsOfWebhookPolicyForbidden creates a ListExecutionsOfWebhookPolicyForbidden with default headers values
func NewListExecutionsOfWebhookPolicyForbidden() *ListExecutionsOfWebhookPolicyForbidden {
	return &ListExecutionsOfWebhookPolicyForbidden{}
}

/*ListExecutionsOfWebhookPolicyForbidden handles this case with default header values.

Forbidden
*/
type ListExecutionsOfWebhookPolicyForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListExecutionsOfWebhookPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions][%d] listExecutionsOfWebhookPolicyForbidden  %+v", 403, o.Payload)
}

func (o *ListExecutionsOfWebhookPolicyForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListExecutionsOfWebhookPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExecutionsOfWebhookPolicyNotFound creates a ListExecutionsOfWebhookPolicyNotFound with default headers values
func NewListExecutionsOfWebhookPolicyNotFound() *ListExecutionsOfWebhookPolicyNotFound {
	return &ListExecutionsOfWebhookPolicyNotFound{}
}

/*ListExecutionsOfWebhookPolicyNotFound handles this case with default header values.

Not found
*/
type ListExecutionsOfWebhookPolicyNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListExecutionsOfWebhookPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions][%d] listExecutionsOfWebhookPolicyNotFound  %+v", 404, o.Payload)
}

func (o *ListExecutionsOfWebhookPolicyNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListExecutionsOfWebhookPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExecutionsOfWebhookPolicyInternalServerError creates a ListExecutionsOfWebhookPolicyInternalServerError with default headers values
func NewListExecutionsOfWebhookPolicyInternalServerError() *ListExecutionsOfWebhookPolicyInternalServerError {
	return &ListExecutionsOfWebhookPolicyInternalServerError{}
}

/*ListExecutionsOfWebhookPolicyInternalServerError handles this case with default header values.

Internal server error
*/
type ListExecutionsOfWebhookPolicyInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListExecutionsOfWebhookPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions][%d] listExecutionsOfWebhookPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *ListExecutionsOfWebhookPolicyInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListExecutionsOfWebhookPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
