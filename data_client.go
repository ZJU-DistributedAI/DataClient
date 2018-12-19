package main

import (
	"github.com/goadesign/goa"
	"math/big"
	"io/ioutil"
	"encoding/json"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"DataClient/app"
	"context"
	"strconv"
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

type DataClientConfig struct {
	// add info
	Add_to_address		string
	Add_data_prefix		string

	// del info
	Del_to_address		string
	Del_data_prefix		string

	// public info
	ETH_HOST 			string
	Value 				string
	Gas_price			string
	Gas_limit			string
}


// Add runs the add action.
func (c *DataClientController) Add(ctx *app.AddDataClientContext) error {
	// check
	if check_valid_arguments(ctx.Hash, ctx.ETHKey) == false {
		return ctx.BadRequest(
			goa.ErrBadRequest("Invalid arguments!"))
	}
	// read config
	config := read_config()
	if config == nil{
		goa.LogInfo(context.Background(), "Config of data client error")
		return ctx.InternalServerError(
			goa.ErrInternal("Config of data client error"))
	}

	// generate transaction
	tx, err := generate_transaction("add", ctx.Hash, ctx.ETHKey, config)
	if err != nil{
		return ctx.InternalServerError(
			goa.ErrInternal("Generate transaction failed!"))
	}

	// sign transaction
	signedTx, err := signTransaction(tx, ctx.ETHKey)
	if err != nil{
		return ctx.InternalServerError(
			goa.ErrInternal("Fail to sign transaction"))
	}
	// send transaction
	transactionHash, err := sendTransaction(signedTx, config.ETH_HOST)
	if err != nil{
		return ctx.InternalServerError(
			goa.ErrInternal("Fail to send transaction"))
	}
	return ctx.OK([]byte(transactionHash))
}

// Agree runs the agree action.
func (c *DataClientController) Agree(ctx *app.AgreeDataClientContext) error {
	return ctx.NotImplemented(goa.ErrInternal("Not implemented"))
}

// AskComputing runs the askComputing action.
func (c *DataClientController) AskComputing(ctx *app.AskComputingDataClientContext) error {
	return ctx.NotImplemented(goa.ErrInternal("Not implemented"))
}

// Del runs the del action.
func (c *DataClientController) Del(ctx *app.DelDataClientContext) error {

	// check
	if check_valid_arguments(ctx.Hash, ctx.ETHKey) == false {
		return ctx.BadRequest(
			goa.ErrBadRequest("Invalid arguments!"))
	}

	// read config
	config := read_config()
	if config == nil{
		goa.LogInfo(context.Background(), "Config of data client error")
		return ctx.InternalServerError(
			goa.ErrInternal("Config of data client error"))
	}

	// generate transaction
	tx, err := generate_transaction("del", ctx.Hash, ctx.ETHKey, config)
	if err != nil{
		return ctx.BadRequest(
			goa.ErrBadRequest("Generate transaction failed!"))
	}

	// sign transaction
	signedTx, err := signTransaction(tx, ctx.ETHKey)
	if err != nil{
		return ctx.InternalServerError(
			goa.ErrInternal("Fail to sign transaction"))
	}

	// send transaction
	transactionHash, err := sendTransaction(signedTx, config.ETH_HOST)
	if err != nil{
		return ctx.InternalServerError(
			goa.ErrInternal("Fail to send transaction"))
	}
	return ctx.OK([]byte(transactionHash))
}

// UploadData runs the uploadData action.
func (c *DataClientController) UploadData(ctx *app.UploadDataDataClientContext) error {

	return ctx.NotImplemented(goa.ErrInternal("Not implemented"))
}


func generate_transaction(op string, hash string, private_key_str string, config * DataClientConfig) (*types.Transaction, error){

	// get paraments of  transaction
	value, gasLimite, gasPrice, err := trans_type(config)
	if err != nil{
		return new(types.Transaction), err
	}

	// data
	to := config.Add_to_address
	data := config.Add_data_prefix + hash
	if op != "add"{
		to 		= config.Del_to_address
		data	= config.Del_data_prefix + hash
	}
	fmt.Println(data)

	// get valid nonce
	privity_key,err := crypto.HexToECDSA(private_key_str)
	if err != nil{
		return new(types.Transaction), err
	}
	client, err := ethclient.Dial(config.ETH_HOST)
	if err != nil {
		return new(types.Transaction),err
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privity_key.PublicKey))
	if err != nil {
		return new(types.Transaction),err
	}

	// a new Transaction
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(to),
		value,
		gasLimite,
		gasPrice,
		[]byte(data))

	return tx, nil
}

func trans_type(config *DataClientConfig)( *big.Int, uint64, *big.Int, error){

	// trans value
	value, err := new(big.Int).SetString(config.Value, 10)
	if err == false{
		goa.LogInfo(context.Background(), "Trans value failed")
		return new(big.Int), uint64(0), new(big.Int), fmt.Errorf("Trans value failed")
	}

	// trans gasLimit
	gas_limit, err_gas :=  strconv.ParseInt(config.Gas_limit, 16, 64)
	if err_gas != nil{
		goa.LogInfo(context.Background(), "Trans value failed")
		return new(big.Int), uint64(0), new(big.Int), fmt.Errorf("Trans value failed")
	}
	gasLimit := uint64(gas_limit)

	// trans gasPrice
	gasPrice, err_price := new(big.Int).SetString(config.Gas_price, 10)
	if err_price == false{
		goa.LogInfo(context.Background(), "Trans gasPrice failed")
		return new(big.Int), uint64(0), new(big.Int), fmt.Errorf("Trans gasPrice failed")
	}

	return value, gasLimit, gasPrice, nil
}

func check_valid_arguments(hash string, private_key string) bool{
	// easy check
	if len(hash) != 46 || len(private_key) != 64{
		return false
	}
	return true
}

func read_config() *DataClientConfig{

	// read file
	config_json,err := ioutil.ReadFile("config.json")
	if err != nil{
		return nil
	}

	// parse json string
	config := &DataClientConfig{}
	err = json.Unmarshal([]byte(config_json), &config)
	if err != nil{
		return nil
	}

	return config
}

func signTransaction(transaction * types.Transaction, private_key_str string) (*types.Transaction, error){

	// get private key
	privity_key,err := crypto.HexToECDSA(private_key_str)
	if err != nil{
		return new(types.Transaction), err
	}

	// get auth for sign
	auth := bind.NewKeyedTransactor(privity_key)
	auth.Nonce = big.NewInt(int64(transaction.Nonce()))
	auth.Value = transaction.Value()
	auth.GasLimit = transaction.Gas()
	auth.GasPrice = transaction.GasPrice()
	auth.From = crypto.PubkeyToAddress(privity_key.PublicKey)

	//chainID := big.NewInt(int64(ChainID))
	//signer := types.NewEIP155Signer(chainID)

	// sign
	signer := types.HomesteadSigner{}
	signedTx ,err:= auth.Signer(signer, auth.From, transaction)
	return signedTx, err
}

func sendTransaction(signedTx * types.Transaction, ETH_HOST string) (string, error) {
	// get client
	client, err := ethclient.Dial(ETH_HOST)
	if err != nil {
		return "",err
	}

	// send
	txErr := client.SendTransaction(context.Background(), signedTx)
	if txErr != nil {
		return "",txErr
	}

	_, bind_err := bind.WaitMined(context.Background(), client, signedTx)
	if bind_err != nil{
		return "",bind_err
	}

	return signedTx.Hash().String(), nil
}