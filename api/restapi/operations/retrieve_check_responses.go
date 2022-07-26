// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Konstantinov-Innokentii/mrmixr/api/models"
)

// RetrieveCheckOKCode is the HTTP code returned for type RetrieveCheckOK
const RetrieveCheckOKCode int = 200

/*RetrieveCheckOK OK

swagger:response retrieveCheckOK
*/
type RetrieveCheckOK struct {

	/*
	  In: Body
	*/
	Payload *models.Check `json:"body,omitempty"`
}

// NewRetrieveCheckOK creates RetrieveCheckOK with default headers values
func NewRetrieveCheckOK() *RetrieveCheckOK {

	return &RetrieveCheckOK{}
}

// WithPayload adds the payload to the retrieve check o k response
func (o *RetrieveCheckOK) WithPayload(payload *models.Check) *RetrieveCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve check o k response
func (o *RetrieveCheckOK) SetPayload(payload *models.Check) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*RetrieveCheckDefault General errors

swagger:response retrieveCheckDefault
*/
type RetrieveCheckDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveCheckDefault creates RetrieveCheckDefault with default headers values
func NewRetrieveCheckDefault(code int) *RetrieveCheckDefault {
	if code <= 0 {
		code = 500
	}

	return &RetrieveCheckDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the retrieve check default response
func (o *RetrieveCheckDefault) WithStatusCode(code int) *RetrieveCheckDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the retrieve check default response
func (o *RetrieveCheckDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the retrieve check default response
func (o *RetrieveCheckDefault) WithPayload(payload *models.Error) *RetrieveCheckDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve check default response
func (o *RetrieveCheckDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveCheckDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}