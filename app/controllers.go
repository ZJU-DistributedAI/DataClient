// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "DataClient": Application Controllers
//
// Command:
// $ goagen
// --design=DataClient/design
// --out=$(GOPATH)\src\DataClient
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// DataClientController is the controller interface for the DataClient actions.
type DataClientController interface {
	goa.Muxer
	Add(*AddDataClientContext) error
	Agree(*AgreeDataClientContext) error
	AskComputing(*AskComputingDataClientContext) error
	Del(*DelDataClientContext) error
	UploadData(*UploadDataDataClientContext) error
}

// MountDataClientController "mounts" a DataClient resource controller on the given service.
func MountDataClientController(service *goa.Service, ctrl DataClientController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddDataClientContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Add(rctx)
	}
	service.Mux.Handle("POST", "/data/add/:hash/:ETH_key", ctrl.MuxHandler("add", h, nil))
	service.LogInfo("mount", "ctrl", "DataClient", "action", "Add", "route", "POST /data/add/:hash/:ETH_key")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAgreeDataClientContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Agree(rctx)
	}
	service.Mux.Handle("POST", "/data/agree/:ETH_key/:data_hash/:contract_hash", ctrl.MuxHandler("agree", h, nil))
	service.LogInfo("mount", "ctrl", "DataClient", "action", "Agree", "route", "POST /data/agree/:ETH_key/:data_hash/:contract_hash")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAskComputingDataClientContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.AskComputing(rctx)
	}
	service.Mux.Handle("POST", "/data/askComputing/:ETH_key/:computing_hash/:contract_hash/:public_key", ctrl.MuxHandler("askComputing", h, nil))
	service.LogInfo("mount", "ctrl", "DataClient", "action", "AskComputing", "route", "POST /data/askComputing/:ETH_key/:computing_hash/:contract_hash/:public_key")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDelDataClientContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Del(rctx)
	}
	service.Mux.Handle("POST", "/data/del/:hash/:ETH_key", ctrl.MuxHandler("del", h, nil))
	service.LogInfo("mount", "ctrl", "DataClient", "action", "Del", "route", "POST /data/del/:hash/:ETH_key")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUploadDataDataClientContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.UploadData(rctx)
	}
	service.Mux.Handle("POST", "/data/upload/:encrypt_data_hash/:ETH_key/:data_hash/:contract_hash", ctrl.MuxHandler("uploadData", h, nil))
	service.LogInfo("mount", "ctrl", "DataClient", "action", "UploadData", "route", "POST /data/upload/:encrypt_data_hash/:ETH_key/:data_hash/:contract_hash")
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerUIController is the controller interface for the SwaggerUI actions.
type SwaggerUIController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerUIController "mounts" a SwaggerUI resource controller on the given service.
func MountSwaggerUIController(service *goa.Service, ctrl SwaggerUIController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/swagger-ui/*filepath", "swagger-ui/")
	service.Mux.Handle("GET", "/swagger-ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "SwaggerUI", "files", "swagger-ui/", "route", "GET /swagger-ui/*filepath")

	h = ctrl.FileHandler("/swagger-ui/", "swagger-ui\\index.html")
	service.Mux.Handle("GET", "/swagger-ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "SwaggerUI", "files", "swagger-ui\\index.html", "route", "GET /swagger-ui/")
}
