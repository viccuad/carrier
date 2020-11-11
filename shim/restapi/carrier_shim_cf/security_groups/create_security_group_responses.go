// Code generated by go-swagger; DO NOT EDIT.

package security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// CreateSecurityGroupCreatedCode is the HTTP code returned for type CreateSecurityGroupCreated
const CreateSecurityGroupCreatedCode int = 201

/*CreateSecurityGroupCreated successful response

swagger:response createSecurityGroupCreated
*/
type CreateSecurityGroupCreated struct {

	/*
	  In: Body
	*/
	Payload *models.CreateSecurityGroupResponse `json:"body,omitempty"`
}

// NewCreateSecurityGroupCreated creates CreateSecurityGroupCreated with default headers values
func NewCreateSecurityGroupCreated() *CreateSecurityGroupCreated {

	return &CreateSecurityGroupCreated{}
}

// WithPayload adds the payload to the create security group created response
func (o *CreateSecurityGroupCreated) WithPayload(payload *models.CreateSecurityGroupResponse) *CreateSecurityGroupCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create security group created response
func (o *CreateSecurityGroupCreated) SetPayload(payload *models.CreateSecurityGroupResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSecurityGroupCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}