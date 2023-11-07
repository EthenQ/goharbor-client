// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ethenq/goharbor-client/v5/apiv2/model"
)

// SearchLdapGroupReader is a Reader for the SearchLdapGroup structure.
type SearchLdapGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchLdapGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchLdapGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchLdapGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchLdapGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchLdapGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchLdapGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchLdapGroupOK creates a SearchLdapGroupOK with default headers values
func NewSearchLdapGroupOK() *SearchLdapGroupOK {
	return &SearchLdapGroupOK{}
}

/*SearchLdapGroupOK handles this case with default header values.

Search ldap group successfully.
*/
type SearchLdapGroupOK struct {
	Payload []*model.UserGroup
}

func (o *SearchLdapGroupOK) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupOK  %+v", 200, o.Payload)
}

func (o *SearchLdapGroupOK) GetPayload() []*model.UserGroup {
	return o.Payload
}

func (o *SearchLdapGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLdapGroupBadRequest creates a SearchLdapGroupBadRequest with default headers values
func NewSearchLdapGroupBadRequest() *SearchLdapGroupBadRequest {
	return &SearchLdapGroupBadRequest{}
}

/*SearchLdapGroupBadRequest handles this case with default header values.

Bad request
*/
type SearchLdapGroupBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SearchLdapGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupBadRequest  %+v", 400, o.Payload)
}

func (o *SearchLdapGroupBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SearchLdapGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLdapGroupUnauthorized creates a SearchLdapGroupUnauthorized with default headers values
func NewSearchLdapGroupUnauthorized() *SearchLdapGroupUnauthorized {
	return &SearchLdapGroupUnauthorized{}
}

/*SearchLdapGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type SearchLdapGroupUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SearchLdapGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchLdapGroupUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SearchLdapGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLdapGroupForbidden creates a SearchLdapGroupForbidden with default headers values
func NewSearchLdapGroupForbidden() *SearchLdapGroupForbidden {
	return &SearchLdapGroupForbidden{}
}

/*SearchLdapGroupForbidden handles this case with default header values.

Forbidden
*/
type SearchLdapGroupForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SearchLdapGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupForbidden  %+v", 403, o.Payload)
}

func (o *SearchLdapGroupForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SearchLdapGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLdapGroupInternalServerError creates a SearchLdapGroupInternalServerError with default headers values
func NewSearchLdapGroupInternalServerError() *SearchLdapGroupInternalServerError {
	return &SearchLdapGroupInternalServerError{}
}

/*SearchLdapGroupInternalServerError handles this case with default header values.

Internal server error
*/
type SearchLdapGroupInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SearchLdapGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchLdapGroupInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SearchLdapGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
