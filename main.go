//go:generate goagen bootstrap -d data-client/design

package main

import (
	"data-client/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("Data Client")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "DataClient" controller
	c := NewDataClientController(service)
	app.MountDataClientController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)
	// Mount "swagger-ui" controller
	c3 := NewSwaggerUIController(service)
	app.MountSwaggerUIController(service, c3)

	// Start service
	if err := service.ListenAndServe(":2626"); err != nil {
		service.LogError("startup", "err", err)
	}

}
