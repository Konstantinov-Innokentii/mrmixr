// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListChecksHandlerFunc turns a function with the right signature into a list checks handler
type ListChecksHandlerFunc func(ListChecksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListChecksHandlerFunc) Handle(params ListChecksParams) middleware.Responder {
	return fn(params)
}

// ListChecksHandler interface for that can handle valid list checks params
type ListChecksHandler interface {
	Handle(ListChecksParams) middleware.Responder
}

// NewListChecks creates a new http.Handler for the list checks operation
func NewListChecks(ctx *middleware.Context, handler ListChecksHandler) *ListChecks {
	return &ListChecks{Context: ctx, Handler: handler}
}

/* ListChecks swagger:route GET /checks listChecks

ListChecks list checks API

*/
type ListChecks struct {
	Context *middleware.Context
	Handler ListChecksHandler
}

func (o *ListChecks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListChecksParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}