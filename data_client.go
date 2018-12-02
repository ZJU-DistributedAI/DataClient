package main

import (
	"DataClient/app"
	"github.com/goadesign/goa"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"crypto/ecdsa"
)

type DataClientController struct {
	*goa.Controller
}

func NewDataClientController(service *goa.Service) *DataClientController {
	return &DataClientController{Controller: service.NewController("DataClientController")}
}

var used_key_file = true
var keyStoreDir 	= "UTC--2018-12-02T05-46-18.847380197Z--8bcf2f7c9445a5768d7ec98658fa317eb8be1cc8"
var privateKeyStr 	= "9eefbcbd4e6061d49722fe04f9f14127aa2bed5160202c009549c73c4099445a"
var ETH_HOST 		= "http://47.52.163.119:8545"
var password  		= "   "
var ChainID 		= 999

func (c *DataClientController) Add(ctx *app.AddDataClientContext) error {

	if len(ctx.Hash) != 46 {
		return ctx.BadRequest(
			goa.ErrBadRequest("Hash invalid!"))
	}

	var to = "4a09e270bf5bae6ccda090cea401fae587b87ba6"
	value := big.NewInt(0x0)
	gasPrice := big.NewInt(200000000)
	gasLimit := uint64(0xfffff)
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
	var fromPrivkey *ecdsa.PrivateKey
	if used_key_file {
		fromKeystore,err := ioutil.ReadFile(keyStoreDir)
		if err != nil{
			fmt.Println("Read key fail")
			return "",err
		}
		fromKey,err := keystore.DecryptKey(fromKeystore, password)
		fromPrivkey = fromKey.PrivateKey
	} else{
		privkey,err := crypto.HexToECDSA(privateKeyStr)
		if err != nil{
			fmt.Println("Get key fail")
			return "",err
		}
		fromPrivkey = privkey
	}

	client, err := ethclient.Dial(ETH_HOST)
	if err != nil {
		fmt.Println("client connection error")
		return "",err
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(fromPrivkey.PublicKey))

	auth := bind.NewKeyedTransactor(fromPrivkey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = gasLimite
	auth.GasPrice = gasPirce
	auth.From = crypto.PubkeyToAddress(fromPrivkey.PublicKey)

	// a new Transaction
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(to),
		value,
		gasLimite,
		gasPirce,
		[]byte(data))

	chainID := big.NewInt(int64(ChainID))
	signer := types.NewEIP155Signer(chainID)
	//signer := types.HomesteadSigner{}

	signedTx ,err:= auth.Signer(signer, auth.From, tx)

	txErr := client.SendTransaction(context.Background(), signedTx)
	fmt.Println(client.BalanceAt(context.Background(), crypto.PubkeyToAddress(fromPrivkey.PublicKey), nil))
	if txErr != nil {
		fmt.Println("send tx error")
		panic(txErr)
		return "",txErr
	}
	fmt.Println(signedTx.Hash().String())
	bind.WaitMined(context.Background(), client, signedTx)
	return signedTx.Hash().String(), nil
}

