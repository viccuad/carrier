// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RemoveManagerFromSpaceCreatedCode is the HTTP code returned for type RemoveManagerFromSpaceCreated
const RemoveManagerFromSpaceCreatedCode int = 201

/*RemoveManagerFromSpaceCreated successful response

swagger:response removeManagerFromSpaceCreated
*/
type RemoveManagerFromSpaceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RemoveManagerFromSpaceResponse `json:"body,omitempty"`
}

// NewRemoveManagerFromSpaceCreated creates RemoveManagerFromSpaceCreated with default headers values
func NewRemoveManagerFromSpaceCreated() *RemoveManagerFromSpaceCreated {

	return &RemoveManagerFromSpaceCreated{}
}

// WithPayload adds the payload to the remove manager from space created response
func (o *RemoveManagerFromSpaceCreated) WithPayload(payload *models.RemoveManagerFromSpaceResponse) *RemoveManagerFromSpaceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove manager from space created response
func (o *RemoveManagerFromSpaceCreated) SetPayload(payload *models.RemoveManagerFromSpaceResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveManagerFromSpaceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}