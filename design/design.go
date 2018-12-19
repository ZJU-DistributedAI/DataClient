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
		Routing(POST("/add/:hash/:ETH_key"))
		Params(func() {
			Param("hash", String, "data IPFS address")						// 数据的ipfs地址
			Param("ETH_key", String, "ETH private key for transaction")		// 以太坊交易秘钥，以后会隐藏
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("del", func() {
		Description("delete data hash")
		Routing(POST("/del/:hash/:ETH_key"))
		Params(func() {
			Param("hash", String, "data IPFS address")						// 数据的ipfs地址
			Param("ETH_key", String, "ETH private key for transaction")		// 以太坊交易秘钥，以后会隐藏
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("agree", func() {
		Description("agree data request")
		Routing(POST("/agree/:ETH_key/:data_hash/:contract_hash"))
		Params(func() {
			Param("ETH_key", String, "ETH private key for transaction")		// 以太坊交易秘钥，以后会隐藏
			Param("data_hash", String, "data hash")							// 被请求的数据的ipfs地址
			Param("contract_hash", String, "smart contract hash")			// 智能合约的地址，智能合约地址和被请求的数据的地址可以成为数据请求的唯一标识
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotImplemented, ErrorMedia)
	})

	Action("askComputing", func() {
		Description("ask for computing for data request")
		Routing(POST("/askComputing/:ETH_key/:computing_hash/:contract_hash/:public_key"))
		Params(func() {
			// 智能合约地址，被请求的运算资源地址，请求运算资源的客户端钱包地址可以成为运算资源请求的唯一标识
			Param("ETH_key", String, "ETH private key for transaction")		// 以太坊交易秘钥，以后会隐藏
			Param("computing_hash", String, "computing resourse hash")		// 被请求的数据的运算资源地址
			Param("contract_hash", String, "smart contract hash")			// 智能合约的地址
			Param("public_key", String, "ETH public key(Wallet address)")	// 数据方客户端的公钥，即钱包地址
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotImplemented, ErrorMedia)
	})

	Action("uploadData", func() {
		Description("upload encrypted data[hash] for data request")
		Routing(POST("/upload/:encrypt_data_hash/:ETH_key/:data_hash/:contract_hash"))
		Params(func() {
			Param("encrypt_data_hash", String, "encrypted data hash")			// 加密数据的最终上传
			Param("ETH_key", String, "ETH private key for transaction")		// 以太坊交易秘钥，以后会隐藏
			Param("data_hash", String, "data hash")							// 被请求的数据的ipfs地址
			Param("contract_hash", String, "smart contract hash")			// 智能合约的地址，智能合约地址和被请求的数据的地址可以成为数据请求的唯一标识

		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotImplemented, ErrorMedia)
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