// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListServiceBrokerDeleteEventsExperimentalOKCode is the HTTP code returned for type ListServiceBrokerDeleteEventsExperimentalOK
const ListServiceBrokerDeleteEventsExperimentalOKCode int = 200

/*ListServiceBrokerDeleteEventsExperimentalOK successful response

swagger:response listServiceBrokerDeleteEventsExperimentalOK
*/
type ListServiceBrokerDeleteEventsExperimentalOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListServiceBrokerDeleteEventsExperimentalResponsePaged `json:"body,omitempty"`
}

// NewListServiceBrokerDeleteEventsExperimentalOK creates ListServiceBrokerDeleteEventsExperimentalOK with default headers values
func NewListServiceBrokerDeleteEventsExperimentalOK() *ListServiceBrokerDeleteEventsExperimentalOK {

	return &ListServiceBrokerDeleteEventsExperimentalOK{}
}

// WithPayload adds the payload to the list service broker delete events experimental o k response
func (o *ListServiceBrokerDeleteEventsExperimentalOK) WithPayload(payload *models.ListServiceBrokerDeleteEventsExperimentalResponsePaged) *ListServiceBrokerDeleteEventsExperimentalOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service broker delete events experimental o k response
func (o *ListServiceBrokerDeleteEventsExperimentalOK) SetPayload(payload *models.ListServiceBrokerDeleteEventsExperimentalResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceBrokerDeleteEventsExperimentalOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}