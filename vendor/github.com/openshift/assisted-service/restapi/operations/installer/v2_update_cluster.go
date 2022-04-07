// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2UpdateClusterHandlerFunc turns a function with the right signature into a v2 update cluster handler
type V2UpdateClusterHandlerFunc func(V2UpdateClusterParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2UpdateClusterHandlerFunc) Handle(params V2UpdateClusterParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2UpdateClusterHandler interface for that can handle valid v2 update cluster params
type V2UpdateClusterHandler interface {
	Handle(V2UpdateClusterParams, interface{}) middleware.Responder
}

// NewV2UpdateCluster creates a new http.Handler for the v2 update cluster operation
func NewV2UpdateCluster(ctx *middleware.Context, handler V2UpdateClusterHandler) *V2UpdateCluster {
	return &V2UpdateCluster{Context: ctx, Handler: handler}
}

/* V2UpdateCluster swagger:route PATCH /v2/clusters/{cluster_id} installer v2UpdateCluster

Updates an OpenShift cluster definition.

*/
type V2UpdateCluster struct {
	Context *middleware.Context
	Handler V2UpdateClusterHandler
}

func (o *V2UpdateCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV2UpdateClusterParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
