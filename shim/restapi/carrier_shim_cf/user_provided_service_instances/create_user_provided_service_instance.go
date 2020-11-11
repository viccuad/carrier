// Code generated by go-swagger; DO NOT EDIT.

package user_provided_service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateUserProvidedServiceInstanceHandlerFunc turns a function with the right signature into a create user provided service instance handler
type CreateUserProvidedServiceInstanceHandlerFunc func(CreateUserProvidedServiceInstanceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateUserProvidedServiceInstanceHandlerFunc) Handle(params CreateUserProvidedServiceInstanceParams) middleware.Responder {
	return fn(params)
}

// CreateUserProvidedServiceInstanceHandler interface for that can handle valid create user provided service instance params
type CreateUserProvidedServiceInstanceHandler interface {
	Handle(CreateUserProvidedServiceInstanceParams) middleware.Responder
}

// NewCreateUserProvidedServiceInstance creates a new http.Handler for the create user provided service instance operation
func NewCreateUserProvidedServiceInstance(ctx *middleware.Context, handler CreateUserProvidedServiceInstanceHandler) *CreateUserProvidedServiceInstance {
	return &CreateUserProvidedServiceInstance{Context: ctx, Handler: handler}
}

/*CreateUserProvidedServiceInstance swagger:route POST /user_provided_service_instances userProvidedServiceInstances createUserProvidedServiceInstance

Creating a User Provided Service Instance

curl --insecure -i %s/v2/user_provided_service_instances -X POST -H 'Authorization: %s' -d '%s'

*/
type CreateUserProvidedServiceInstance struct {
	Context *middleware.Context
	Handler CreateUserProvidedServiceInstanceHandler
}

func (o *CreateUserProvidedServiceInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateUserProvidedServiceInstanceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}