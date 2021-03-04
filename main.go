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

	fmt.Println("区块0的哈希值：",blockchain.Blocks[0])
	fmt.Println("区块1的哈希值：",blockchain.Blocks[1])
}
