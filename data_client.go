package main

import (
	"DataClient/app"
	"github.com/goadesign/goa"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
)

type DataClientController struct {
	*goa.Controller
}

func NewDataClientController(service *goa.Service) *DataClientController {
	return &DataClientController{Controller: service.NewController("DataClientController")}
}


var keyStoreDir 	= "UTC--2018-12-01T11-00-43.186381425Z--0d3bf5d5e8101cf0f49f2c0013f470d39527d01b"
var ETH_HOST 		= "http://47.52.163.119:8545"
var password  		= "   "

func (c *DataClientController) Add(ctx *app.AddDataClientContext) error {

	if len(ctx.Hash) != 46 {
		return ctx.BadRequest(
			goa.ErrBadRequest("Hash invalid!"))
	}

	var to = "3dfd86891ea4a634e4a6e1c8e75d1a92a0928346"
	value := big.NewInt(0)
	gasPrice := big.NewInt(2000000000)
	gasLimit := uint64(2711301)
	data := "add " + ctx.Hash

	transactionHash, err := sendTransaction(to, gasLimit, gasPrice, value, data)
	if err != nil{
		return ctx.InternalServerError(
			goa.ErrInternal("Fail to send transation"))
	}

	return ctx.OK([]byte(transactionHash))
}


func (c *DataClientController) Del(ctx *app.DelDataClientContext) error {


	return nil
}

func sendTransaction(to string, gasLimite uint64, gasPirce *big.Int, value *big.Int, data string) (string ,error){
	fromKeystore,err := ioutil.ReadFile(keyStoreDir)
	if err != nil{
		fmt.Println("Read key fail")
		return "",err
	}
	fromKey,err := keystore.DecryptKey(fromKeystore, password)
	fromPrivkey := fromKey.PrivateKey

	// a new Transaction
	tx := types.NewTransaction(
		0x0,
		common.HexToAddress(to),
		value,
		gasLimite,
		gasPirce,
		[]byte(data))

	signature, signatureErr := crypto.Sign(tx.Hash().Bytes(), fromPrivkey)
	if signatureErr != nil {
		fmt.Println("signature create error")
		return "",signatureErr
	}

	signedTx, signErr := tx.WithSignature(types.HomesteadSigner{}, signature)
	if signErr != nil{
		fmt.Println("signature create error")
		return "",signErr
	}

	client, err := ethclient.Dial(ETH_HOST)
	if err != nil {
		fmt.Println("client connection error")
		return "",err
	}

	txErr := client.SendTransaction(context.Background(), signedTx)
	if txErr != nil {
		fmt.Println("send tx error")
		panic(txErr)
		return "",txErr
	}
	return signedTx.Hash().String(), nil
}

