// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2CompleteInstallationAcceptedCode is the HTTP code returned for type V2CompleteInstallationAccepted
const V2CompleteInstallationAcceptedCode int = 202

/*V2CompleteInstallationAccepted Success.

swagger:response v2CompleteInstallationAccepted
*/
type V2CompleteInstallationAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewV2CompleteInstallationAccepted creates V2CompleteInstallationAccepted with default headers values
func NewV2CompleteInstallationAccepted() *V2CompleteInstallationAccepted {

	return &V2CompleteInstallationAccepted{}
}

// WithPayload adds the payload to the v2 complete installation accepted response
func (o *V2CompleteInstallationAccepted) WithPayload(payload *models.Cluster) *V2CompleteInstallationAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 complete installation accepted response
func (o *V2CompleteInstallationAccepted) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2CompleteInstallationAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2CompleteInstallationUnauthorizedCode is the HTTP code returned for type V2CompleteInstallationUnauthorized
const V2CompleteInstallationUnauthorizedCode int = 401

/*V2CompleteInstallationUnauthorized Unauthorized.

swagger:response v2CompleteInstallationUnauthorized
*/
type V2CompleteInstallationUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2CompleteInstallationUnauthorized creates V2CompleteInstallationUnauthorized with default headers values
func NewV2CompleteInstallationUnauthorized() *V2CompleteInstallationUnauthorized {

	return &V2CompleteInstallationUnauthorized{}
}

// WithPayload adds the payload to the v2 complete installation unauthorized response
func (o *V2CompleteInstallationUnauthorized) WithPayload(payload *models.InfraError) *V2CompleteInstallationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 complete installation unauthorized response
func (o *V2CompleteInstallationUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2CompleteInstallationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2CompleteInstallationForbiddenCode is the HTTP code returned for type V2CompleteInstallationForbidden
const V2CompleteInstallationForbiddenCode int = 403

/*V2CompleteInstallationForbidden Forbidden.

swagger:response v2CompleteInstallationForbidden
*/
type V2CompleteInstallationForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2CompleteInstallationForbidden creates V2CompleteInstallationForbidden with default headers values
func NewV2CompleteInstallationForbidden() *V2CompleteInstallationForbidden {

	return &V2CompleteInstallationForbidden{}
}

// WithPayload adds the payload to the v2 complete installation forbidden response
func (o *V2CompleteInstallationForbidden) WithPayload(payload *models.InfraError) *V2CompleteInstallationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 complete installation forbidden response
func (o *V2CompleteInstallationForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2CompleteInstallationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2CompleteInstallationNotFoundCode is the HTTP code returned for type V2CompleteInstallationNotFound
const V2CompleteInstallationNotFoundCode int = 404

/*V2CompleteInstallationNotFound Error.

swagger:response v2CompleteInstallationNotFound
*/
type V2CompleteInstallationNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2CompleteInstallationNotFound creates V2CompleteInstallationNotFound with default headers values
func NewV2CompleteInstallationNotFound() *V2CompleteInstallationNotFound {

	return &V2CompleteInstallationNotFound{}
}

// WithPayload adds the payload to the v2 complete installation not found response
func (o *V2CompleteInstallationNotFound) WithPayload(payload *models.Error) *V2CompleteInstallationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 complete installation not found response
func (o *V2CompleteInstallationNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2CompleteInstallationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2CompleteInstallationMethodNotAllowedCode is the HTTP code returned for type V2CompleteInstallationMethodNotAllowed
const V2CompleteInstallationMethodNotAllowedCode int = 405

/*V2CompleteInstallationMethodNotAllowed Method Not Allowed.

swagger:response v2CompleteInstallationMethodNotAllowed
*/
type V2CompleteInstallationMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2CompleteInstallationMethodNotAllowed creates V2CompleteInstallationMethodNotAllowed with default headers values
func NewV2CompleteInstallationMethodNotAllowed() *V2CompleteInstallationMethodNotAllowed {

	return &V2CompleteInstallationMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 complete installation method not allowed response
func (o *V2CompleteInstallationMethodNotAllowed) WithPayload(payload *models.Error) *V2CompleteInstallationMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 complete installation method not allowed response
func (o *V2CompleteInstallationMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2CompleteInstallationMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2CompleteInstallationConflictCode is the HTTP code returned for type V2CompleteInstallationConflict
const V2CompleteInstallationConflictCode int = 409

/*V2CompleteInstallationConflict Error.

swagger:response v2CompleteInstallationConflict
*/
type V2CompleteInstallationConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2CompleteInstallationConflict creates V2CompleteInstallationConflict with default headers values
func NewV2CompleteInstallationConflict() *V2CompleteInstallationConflict {

	return &V2CompleteInstallationConflict{}
}

// WithPayload adds the payload to the v2 complete installation conflict response
func (o *V2CompleteInstallationConflict) WithPayload(payload *models.Error) *V2CompleteInstallationConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 complete installation conflict response
func (o *V2CompleteInstallationConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2CompleteInstallationConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2CompleteInstallationInternalServerErrorCode is the HTTP code returned for type V2CompleteInstallationInternalServerError
const V2CompleteInstallationInternalServerErrorCode int = 500

/*V2CompleteInstallationInternalServerError Error.

swagger:response v2CompleteInstallationInternalServerError
*/
type V2CompleteInstallationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2CompleteInstallationInternalServerError creates V2CompleteInstallationInternalServerError with default headers values
func NewV2CompleteInstallationInternalServerError() *V2CompleteInstallationInternalServerError {

	return &V2CompleteInstallationInternalServerError{}
}

// WithPayload adds the payload to the v2 complete installation internal server error response
func (o *V2CompleteInstallationInternalServerError) WithPayload(payload *models.Error) *V2CompleteInstallationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 complete installation internal server error response
func (o *V2CompleteInstallationInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2CompleteInstallationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2CompleteInstallationServiceUnavailableCode is the HTTP code returned for type V2CompleteInstallationServiceUnavailable
const V2CompleteInstallationServiceUnavailableCode int = 503

/*V2CompleteInstallationServiceUnavailable Unavailable.

swagger:response v2CompleteInstallationServiceUnavailable
*/
type V2CompleteInstallationServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2CompleteInstallationServiceUnavailable creates V2CompleteInstallationServiceUnavailable with default headers values
func NewV2CompleteInstallationServiceUnavailable() *V2CompleteInstallationServiceUnavailable {

	return &V2CompleteInstallationServiceUnavailable{}
}

// WithPayload adds the payload to the v2 complete installation service unavailable response
func (o *V2CompleteInstallationServiceUnavailable) WithPayload(payload *models.Error) *V2CompleteInstallationServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 complete installation service unavailable response
func (o *V2CompleteInstallationServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2CompleteInstallationServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
