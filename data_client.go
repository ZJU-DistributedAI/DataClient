package main

import (
	"data-client/app"
	"github.com/goadesign/goa"
	"fmt"
)

// DataClientController implements the DataClient resource.
type DataClientController struct {
	*goa.Controller
}

// NewDataClientController creates a DataClient controller.
func NewDataClientController(service *goa.Service) *DataClientController {
	return &DataClientController{Controller: service.NewController("DataClientController")}
}

var transaction_api = ""

// Add runs the add action.
func (c *DataClientController) Add(ctx *app.AddDataClientContext) error {
	// DataClientController_Add: start_implement

	// Put your logic here
	if len(ctx.Hash) != 46{
		fmt.Println("ssss")
	}


	return ctx.OK([]byte("OK"))
	// DataClientController_Add: end_implement
}

// Del runs the del action.
func (c *DataClientController) Del(ctx *app.DelDataClientContext) error {
	// DataClientController_Del: start_implement

	// Put your logic here

	return nil
	// DataClientController_Del: end_implement
}
