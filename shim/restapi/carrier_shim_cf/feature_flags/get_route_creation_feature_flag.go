// Code generated by go-swagger; DO NOT EDIT.

package feature_flags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRouteCreationFeatureFlagHandlerFunc turns a function with the right signature into a get route creation feature flag handler
type GetRouteCreationFeatureFlagHandlerFunc func(GetRouteCreationFeatureFlagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRouteCreationFeatureFlagHandlerFunc) Handle(params GetRouteCreationFeatureFlagParams) middleware.Responder {
	return fn(params)
}

// GetRouteCreationFeatureFlagHandler interface for that can handle valid get route creation feature flag params
type GetRouteCreationFeatureFlagHandler interface {
	Handle(GetRouteCreationFeatureFlagParams) middleware.Responder
}

// NewGetRouteCreationFeatureFlag creates a new http.Handler for the get route creation feature flag operation
func NewGetRouteCreationFeatureFlag(ctx *middleware.Context, handler GetRouteCreationFeatureFlagHandler) *GetRouteCreationFeatureFlag {
	return &GetRouteCreationFeatureFlag{Context: ctx, Handler: handler}
}

/*GetRouteCreationFeatureFlag swagger:route GET /config/feature_flags/route_creation featureFlags getRouteCreationFeatureFlag

Get the Route Creation feature flag

curl --insecure -i %s/v2/config/feature_flags/route_creation -X GET -H 'Authorization: %s'

*/
type GetRouteCreationFeatureFlag struct {
	Context *middleware.Context
	Handler GetRouteCreationFeatureFlagHandler
}

func (o *GetRouteCreationFeatureFlag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRouteCreationFeatureFlagParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}