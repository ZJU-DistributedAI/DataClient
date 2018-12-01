package design                                     // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Data Client", func() {
	Title("Data client to add or delete data")
	Description("Add or delete data")
	Scheme("http")
	Host("localhost:2626")
})

var _ = Resource("DataClient", func() {
	BasePath("/data")

	Action("add", func() {
		Description("add data hash")
		Routing(POST("/:hash"))
		Params(func() {
			Param("hash", String, "data IPFS address")
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("del", func() {
		Description("delete data hash")
		Routing(DELETE("/:hash"))
		Params(func() {
			Param("hash", String, "data IPFS address")
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