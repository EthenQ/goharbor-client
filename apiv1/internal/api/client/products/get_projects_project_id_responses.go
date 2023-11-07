// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ethenq/goharbor-client/v5/apiv1/model"
)

// GetProjectsProjectIDReader is a Reader for the GetProjectsProjectID structure.
type GetProjectsProjectIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetProjectsProjectIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectsProjectIDOK creates a GetProjectsProjectIDOK with default headers values
func NewGetProjectsProjectIDOK() *GetProjectsProjectIDOK {
	return &GetProjectsProjectIDOK{}
}

/*GetProjectsProjectIDOK handles this case with default header values.

Return matched project information.
*/
type GetProjectsProjectIDOK struct {
	Payload *model.Project
}

func (o *GetProjectsProjectIDOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}][%d] getProjectsProjectIdOK  %+v", 200, o.Payload)
}

func (o *GetProjectsProjectIDOK) GetPayload() *model.Project {
	return o.Payload
}

func (o *GetProjectsProjectIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDUnauthorized creates a GetProjectsProjectIDUnauthorized with default headers values
func NewGetProjectsProjectIDUnauthorized() *GetProjectsProjectIDUnauthorized {
	return &GetProjectsProjectIDUnauthorized{}
}

/*GetProjectsProjectIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetProjectsProjectIDUnauthorized struct {
}

func (o *GetProjectsProjectIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}][%d] getProjectsProjectIdUnauthorized ", 401)
}

func (o *GetProjectsProjectIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDInternalServerError creates a GetProjectsProjectIDInternalServerError with default headers values
func NewGetProjectsProjectIDInternalServerError() *GetProjectsProjectIDInternalServerError {
	return &GetProjectsProjectIDInternalServerError{}
}

/*GetProjectsProjectIDInternalServerError handles this case with default header values.

Internal errors.
*/
type GetProjectsProjectIDInternalServerError struct {
}

func (o *GetProjectsProjectIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}][%d] getProjectsProjectIdInternalServerError ", 500)
}

func (o *GetProjectsProjectIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
