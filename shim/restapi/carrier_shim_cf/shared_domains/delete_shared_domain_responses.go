// Code generated by go-swagger; DO NOT EDIT.

package shared_domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteSharedDomainNoContentCode is the HTTP code returned for type DeleteSharedDomainNoContent
const DeleteSharedDomainNoContentCode int = 204

/*DeleteSharedDomainNoContent successful response

swagger:response deleteSharedDomainNoContent
*/
type DeleteSharedDomainNoContent struct {
}

// NewDeleteSharedDomainNoContent creates DeleteSharedDomainNoContent with default headers values
func NewDeleteSharedDomainNoContent() *DeleteSharedDomainNoContent {

	return &DeleteSharedDomainNoContent{}
}

// WriteResponse to the client
func (o *DeleteSharedDomainNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}