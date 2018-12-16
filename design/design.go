package design                                     // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("DataClient", func() {
	Title("Data client to add or delete data")
	Description("Add or delete data")
	Scheme("http")
	Host("localhost:2626")
})

/*
********************************************************
(1)  Data client
********************************************************
 */


var _ = Resource("DataClient", func() {
	BasePath("/data")

	Action("add", func() {
		Description("add data hash")
		Routing(POST("/add/:hash/:private_key"))
		Params(func() {
			Param("hash", String, "data IPFS address")
			Param("private_key", String, "ETH private key for transaction")
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("del", func() {
		Description("delete data hash")
		Routing(POST("/del/:hash/:private_key"))
		Params(func() {
			Param("hash", String, "data IPFS address")
			Param("private_key", String, "ETH private key for transaction")
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("agree", func() {
		Description("agree data request for request [ID]")
		Routing(POST("/agree/:ETH_key/:request_id"))
		Params(func() {
			Param("ETH_key", String, "ETH private key for transaction")
			Param("request_id", Integer, "request[ID]")
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("askComputing", func() {
		Description("ask for computing for [request_id] on computing resourse[hash]")
		Routing(POST("/askComputing/:hash/:ETH_key/:request_id"))
		Params(func() {
			Param("hash", String, "computing resourse hash")
			Param("ETH_key", String, "ETH private key for transaction")
			Param("request_id", Integer, "request[ID]")
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("uploadData", func() {
		Description("upload encrypted data[hash] for [request_id]")
		Routing(POST("/upload/:hash/:ETH_key/:request_id"))
		Params(func() {
			Param("hash", String, "encrypted data hash")
			Param("ETH_key", String, "ETH private key for transaction")
			Param("request_id", Integer, "request[ID]")

		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
})

var _ = Resource("swagger-ui", func() {

	Files("/swagger-ui/*filepath", "swagger-ui/")
})