/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package commands

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateCommandsInsertHandlerFunc turns a function with the right signature into a weaviate commands insert handler
type WeaviateCommandsInsertHandlerFunc func(WeaviateCommandsInsertParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateCommandsInsertHandlerFunc) Handle(params WeaviateCommandsInsertParams) middleware.Responder {
	return fn(params)
}

// WeaviateCommandsInsertHandler interface for that can handle valid weaviate commands insert params
type WeaviateCommandsInsertHandler interface {
	Handle(WeaviateCommandsInsertParams) middleware.Responder
}

// NewWeaviateCommandsInsert creates a new http.Handler for the weaviate commands insert operation
func NewWeaviateCommandsInsert(ctx *middleware.Context, handler WeaviateCommandsInsertHandler) *WeaviateCommandsInsert {
	return &WeaviateCommandsInsert{Context: ctx, Handler: handler}
}

/*WeaviateCommandsInsert swagger:route POST /commands commands weaviateCommandsInsert

Inserts and sends a new command.

*/
type WeaviateCommandsInsert struct {
	Context *middleware.Context
	Handler WeaviateCommandsInsertHandler
}

func (o *WeaviateCommandsInsert) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateCommandsInsertParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}