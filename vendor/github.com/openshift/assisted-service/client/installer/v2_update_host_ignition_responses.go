// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2UpdateHostIgnitionReader is a Reader for the V2UpdateHostIgnition structure.
type V2UpdateHostIgnitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2UpdateHostIgnitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV2UpdateHostIgnitionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV2UpdateHostIgnitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV2UpdateHostIgnitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2UpdateHostIgnitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2UpdateHostIgnitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2UpdateHostIgnitionMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2UpdateHostIgnitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewV2UpdateHostIgnitionNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2UpdateHostIgnitionCreated creates a V2UpdateHostIgnitionCreated with default headers values
func NewV2UpdateHostIgnitionCreated() *V2UpdateHostIgnitionCreated {
	return &V2UpdateHostIgnitionCreated{}
}

/* V2UpdateHostIgnitionCreated describes a response with status code 201, with default header values.

Success.
*/
type V2UpdateHostIgnitionCreated struct {
}

func (o *V2UpdateHostIgnitionCreated) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/ignition][%d] v2UpdateHostIgnitionCreated ", 201)
}

func (o *V2UpdateHostIgnitionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewV2UpdateHostIgnitionBadRequest creates a V2UpdateHostIgnitionBadRequest with default headers values
func NewV2UpdateHostIgnitionBadRequest() *V2UpdateHostIgnitionBadRequest {
	return &V2UpdateHostIgnitionBadRequest{}
}

/* V2UpdateHostIgnitionBadRequest describes a response with status code 400, with default header values.

Error.
*/
type V2UpdateHostIgnitionBadRequest struct {
	Payload *models.Error
}

func (o *V2UpdateHostIgnitionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/ignition][%d] v2UpdateHostIgnitionBadRequest  %+v", 400, o.Payload)
}
func (o *V2UpdateHostIgnitionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostIgnitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostIgnitionUnauthorized creates a V2UpdateHostIgnitionUnauthorized with default headers values
func NewV2UpdateHostIgnitionUnauthorized() *V2UpdateHostIgnitionUnauthorized {
	return &V2UpdateHostIgnitionUnauthorized{}
}

/* V2UpdateHostIgnitionUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2UpdateHostIgnitionUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2UpdateHostIgnitionUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/ignition][%d] v2UpdateHostIgnitionUnauthorized  %+v", 401, o.Payload)
}
func (o *V2UpdateHostIgnitionUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UpdateHostIgnitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostIgnitionForbidden creates a V2UpdateHostIgnitionForbidden with default headers values
func NewV2UpdateHostIgnitionForbidden() *V2UpdateHostIgnitionForbidden {
	return &V2UpdateHostIgnitionForbidden{}
}

/* V2UpdateHostIgnitionForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2UpdateHostIgnitionForbidden struct {
	Payload *models.InfraError
}

func (o *V2UpdateHostIgnitionForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/ignition][%d] v2UpdateHostIgnitionForbidden  %+v", 403, o.Payload)
}
func (o *V2UpdateHostIgnitionForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UpdateHostIgnitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostIgnitionNotFound creates a V2UpdateHostIgnitionNotFound with default headers values
func NewV2UpdateHostIgnitionNotFound() *V2UpdateHostIgnitionNotFound {
	return &V2UpdateHostIgnitionNotFound{}
}

/* V2UpdateHostIgnitionNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2UpdateHostIgnitionNotFound struct {
	Payload *models.Error
}

func (o *V2UpdateHostIgnitionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/ignition][%d] v2UpdateHostIgnitionNotFound  %+v", 404, o.Payload)
}
func (o *V2UpdateHostIgnitionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostIgnitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostIgnitionMethodNotAllowed creates a V2UpdateHostIgnitionMethodNotAllowed with default headers values
func NewV2UpdateHostIgnitionMethodNotAllowed() *V2UpdateHostIgnitionMethodNotAllowed {
	return &V2UpdateHostIgnitionMethodNotAllowed{}
}

/* V2UpdateHostIgnitionMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2UpdateHostIgnitionMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2UpdateHostIgnitionMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/ignition][%d] v2UpdateHostIgnitionMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *V2UpdateHostIgnitionMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostIgnitionMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostIgnitionInternalServerError creates a V2UpdateHostIgnitionInternalServerError with default headers values
func NewV2UpdateHostIgnitionInternalServerError() *V2UpdateHostIgnitionInternalServerError {
	return &V2UpdateHostIgnitionInternalServerError{}
}

/* V2UpdateHostIgnitionInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2UpdateHostIgnitionInternalServerError struct {
	Payload *models.Error
}

func (o *V2UpdateHostIgnitionInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/ignition][%d] v2UpdateHostIgnitionInternalServerError  %+v", 500, o.Payload)
}
func (o *V2UpdateHostIgnitionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostIgnitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostIgnitionNotImplemented creates a V2UpdateHostIgnitionNotImplemented with default headers values
func NewV2UpdateHostIgnitionNotImplemented() *V2UpdateHostIgnitionNotImplemented {
	return &V2UpdateHostIgnitionNotImplemented{}
}

/* V2UpdateHostIgnitionNotImplemented describes a response with status code 501, with default header values.

Not implemented.
*/
type V2UpdateHostIgnitionNotImplemented struct {
	Payload *models.Error
}

func (o *V2UpdateHostIgnitionNotImplemented) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/ignition][%d] v2UpdateHostIgnitionNotImplemented  %+v", 501, o.Payload)
}
func (o *V2UpdateHostIgnitionNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostIgnitionNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
