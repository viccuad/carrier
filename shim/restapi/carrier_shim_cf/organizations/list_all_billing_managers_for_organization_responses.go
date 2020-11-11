// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllBillingManagersForOrganizationOKCode is the HTTP code returned for type ListAllBillingManagersForOrganizationOK
const ListAllBillingManagersForOrganizationOKCode int = 200

/*ListAllBillingManagersForOrganizationOK successful response

swagger:response listAllBillingManagersForOrganizationOK
*/
type ListAllBillingManagersForOrganizationOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllBillingManagersForOrganizationResponsePaged `json:"body,omitempty"`
}

// NewListAllBillingManagersForOrganizationOK creates ListAllBillingManagersForOrganizationOK with default headers values
func NewListAllBillingManagersForOrganizationOK() *ListAllBillingManagersForOrganizationOK {

	return &ListAllBillingManagersForOrganizationOK{}
}

// WithPayload adds the payload to the list all billing managers for organization o k response
func (o *ListAllBillingManagersForOrganizationOK) WithPayload(payload *models.ListAllBillingManagersForOrganizationResponsePaged) *ListAllBillingManagersForOrganizationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all billing managers for organization o k response
func (o *ListAllBillingManagersForOrganizationOK) SetPayload(payload *models.ListAllBillingManagersForOrganizationResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllBillingManagersForOrganizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}