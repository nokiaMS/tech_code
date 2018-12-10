package main

import (
	"github.com/ethereum/go-ethereum/common"
	"crypto/ecdsa"
	"os"
	"fmt"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

//erc20合约地址。
const ERC2O_CONTRACT_ADDRESS string = "0xaa55b6e94ad44def13eaeb09790216cb56db863e"
//wandappETH.sol合约地址。
const WANDAPP_ETH_CONTRACT_ADDRESS string = "0xe2a73b6a6445810fef8f2f8b49171ff40850965a"

//用户keystore文件路径。
const A_KEYSTORE_FILE string = "/home/guoxu/scenarios/singleEthereum/keystore/UTC--2018-11-23T03-39-31.125978598Z--2f8227345ab3dc0054c6ac2355d27abe584e5cc8"
const KEYSTORE_PASSWD string = "123456"

//字符串地址转换为Address结构体。
func erc20Addr( addr string ) common.Address{
	return common.HexToAddress(addr)
}

//从keystore文件中解析出私钥。
func getPrivKey(file string, passwd string) (*ecdsa.PrivateKey, error) {
	//判断文件是否存在。
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("file:", file, " ", err)
		return nil, err
	}

	//读取文件。
	keyjson, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("file:", file, " ", err)
		return nil, err
	}

	//从文件中获取私钥。
	key, err := keystore.DecryptKey(keyjson, passwd)
	if err != nil {
		fmt.Println("file:", file, " ", err)
		return nil, err
	}

	return key.PrivateKey, nil	//返回私钥。
}