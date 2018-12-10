package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"goSmartSC/wandappETH"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum"
	"fmt"
	"os"
	"os/signal"
)

func getTransactionOpts(priv *ecdsa.PrivateKey, client *ethclient.Client) (*bind.TransactOpts,error) {
	signer := bind.NewKeyedTransactor(priv)	//创建一个交易签名对象。

	//获得账户地址。
	pubKey := priv.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)  //类型断言。
	if !ok {
		log.Fatal("Error casting pub key to ECDSA.")
	}
	addr := crypto.PubkeyToAddress(*pubKeyECDSA).Hex()

	//获得账户对应的nonce值。
	nonce,err := client.PendingNonceAt(context.Background(), common.HexToAddress(addr))
	if err != nil {
		log.Fatal(err)
	}
	signer.Nonce = big.NewInt(int64(nonce))

	signer.Value = big.NewInt(0)
	signer.GasLimit = uint64(1000000)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signer.GasPrice = gasPrice

	return signer, nil
}

func showLogInfo(log types.Log)  {
	addr := log.Address.Hex()
	blockNumber := log.BlockNumber
	blockHash := log.BlockHash.Hex()
	txHash := log.TxHash.Hex()
	txIndex := log.TxIndex

	fmt.Println("====Log Info====")
	fmt.Println("====    Contract address:", addr)
	fmt.Println("====    Block number: ", blockNumber)
	fmt.Println("====    Block hash: ", blockHash)
	fmt.Println("====    transaction hash: ", txHash)
	fmt.Println("====    tx index: ", txIndex)
	fmt.Println("====    data: ", log.Data)
}

//时间处理函数。
func dealWithEvents(client *ethclient.Client, chl chan(types.Log))  {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
						common.HexToAddress(WANDAPP_ETH_CONTRACT_ADDRESS),
						common.HexToAddress(ERC2O_CONTRACT_ADDRESS)},
	}
	sub,err := client.SubscribeFilterLogs(context.Background(), query, chl)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-chl:
			showLogInfo(vLog)
		}
	}
}

func main() {
	/*
	//获得系统环境变量。
	gopath := os.Getenv("GOPATH")
	fmt.Println("GOPATH is: " + gopath)
	*/

	//连接ethereum节点。
	client,err := ethclient.Dial("ws://127.0.0.1:50001")	//不能用http协议，因为http协议是短连接，无法从node返回消息。
	if err != nil {
		log.Fatal(err)
	}

	//创建erc20合约实例。
	erc20Inst,err := wandappETH.NewDemoERC20(erc20Addr(ERC2O_CONTRACT_ADDRESS), client)
	if err != nil {
		log.Fatal(err)
	}

	//启动事件监听器。
	logsChan := make(chan types.Log)	//构造日志类型的通道。
	go dealWithEvents(client, logsChan)	//事件监听例程。

	//获得账户私钥。
	privKey,err := getPrivKey(A_KEYSTORE_FILE, KEYSTORE_PASSWD)
	if err != nil {
		log.Fatal(err)
	}

	signer,err := getTransactionOpts(privKey,client)
	if err != nil {
		log.Fatal(err)
	}

	//用户授权wandappETH.sol合约转账额度。
	erc20Inst.Approve(signer, common.HexToAddress(WANDAPP_ETH_CONTRACT_ADDRESS),big.NewInt(10000));

	/*
	_, err := KeystoreLoad("UTC--2018-07-30T01-39-03.462860004Z--b4173fddaa6e5a3b496460ba440cff0f984b98b8", "123456")
	if err != nil {
		fmt.Println("Priv error:", err)
		return
	}
	*/

	//捕获信号，程序退出。
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("Got singal: ",s)
}
