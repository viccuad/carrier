// Code generated by go-swagger; DO NOT EDIT.

package feature_flags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetServiceInstanceCreationFeatureFlagHandlerFunc turns a function with the right signature into a get service instance creation feature flag handler
type GetServiceInstanceCreationFeatureFlagHandlerFunc func(GetServiceInstanceCreationFeatureFlagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetServiceInstanceCreationFeatureFlagHandlerFunc) Handle(params GetServiceInstanceCreationFeatureFlagParams) middleware.Responder {
	return fn(params)
}

// GetServiceInstanceCreationFeatureFlagHandler interface for that can handle valid get service instance creation feature flag params
type GetServiceInstanceCreationFeatureFlagHandler interface {
	Handle(GetServiceInstanceCreationFeatureFlagParams) middleware.Responder
}

// NewGetServiceInstanceCreationFeatureFlag creates a new http.Handler for the get service instance creation feature flag operation
func NewGetServiceInstanceCreationFeatureFlag(ctx *middleware.Context, handler GetServiceInstanceCreationFeatureFlagHandler) *GetServiceInstanceCreationFeatureFlag {
	return &GetServiceInstanceCreationFeatureFlag{Context: ctx, Handler: handler}
}

/*GetServiceInstanceCreationFeatureFlag swagger:route GET /config/feature_flags/service_instance_creation featureFlags getServiceInstanceCreationFeatureFlag

Get the Service Instance Creation feature flag

curl --insecure -i %s/v2/config/feature_flags/service_instance_creation -X GET -H 'Authorization: %s'

*/
type GetServiceInstanceCreationFeatureFlag struct {
	Context *middleware.Context
	Handler GetServiceInstanceCreationFeatureFlagHandler
}

func (o *GetServiceInstanceCreationFeatureFlag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetServiceInstanceCreationFeatureFlagParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}