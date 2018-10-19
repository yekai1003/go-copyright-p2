package eths

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

//创建一个账户
func NewAcc(pass, connstr string) (string, error) {
	//连接到geth
	client, err := rpc.Dial(connstr)
	if err != nil {
		fmt.Println("failed to connect to geth", err)
		return "", err
	}

	//创建账户
	var account string
	err = client.Call(&account, "personal_newAccount", pass)
	if err != nil {
		fmt.Println("failed to create acc", err)
		return "", err
	}
	fmt.Println("acc==", account)
	return account, nil
}
