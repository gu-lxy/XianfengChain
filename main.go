package main

import (
	"XianFengChain/chain"
	"fmt"
)

func main(){
	fmt.Println("hello World")

	blockchain := chain.CreateChainWithGensis([]byte("hello world"))
	blockchain.CreateNewBlock([]byte("hello"))

	fmt.Println("区块链中的区块的个数：",len(blockchain.Blocks))

	fmt.Println("区块0：",blockchain.Blocks[0])
	//fmt.Println("区块1：",blockchain.Blocks[1])

	firstBlock := blockchain.Blocks[0]
	firstBytes, err := firstBlock.Sweialize()
	if err != nil {
		panic(err.Error())
	}
	//反序列化，验证逆过程
	deFirstBlock, err := chain.Deserialize(firstBytes)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(deFirstBlock)
}
