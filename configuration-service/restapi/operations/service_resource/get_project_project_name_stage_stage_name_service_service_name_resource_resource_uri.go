// Code generated by go-swagger; DO NOT EDIT.

package service_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURIHandlerFunc turns a function with the right signature into a get project project name stage stage name service service name resource resource URI handler
type GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURIHandlerFunc func(GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURIParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURIHandlerFunc) Handle(params GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURIParams) middleware.Responder {
	return fn(params)
}

// GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURIHandler interface for that can handle valid get project project name stage stage name service service name resource resource URI params
type GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURIHandler interface {
	Handle(GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURIParams) middleware.Responder
}

// NewGetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURI creates a new http.Handler for the get project project name stage stage name service service name resource resource URI operation
func NewGetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURI(ctx *middleware.Context, handler GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURIHandler) *GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURI {
	return &GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURI{Context: ctx, Handler: handler}
}

/*GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURI swagger:route GET /project/{projectName}/stage/{stageName}/service/{serviceName}/resource/{resourceURI} Service Resource getProjectProjectNameStageStageNameServiceServiceNameResourceResourceUri

Get the specified resource

*/
type GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURI struct {
	Context *middleware.Context
	Handler GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURIHandler
}

func (o *GetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURI) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetProjectProjectNameStageStageNameServiceServiceNameResourceResourceURIParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
