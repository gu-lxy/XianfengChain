package main

import (
	"XianFengChain/chain"
	"fmt"
	"github.com/bolt-master/bolt-master"
)

const BLOCKS = "xianfengchain.db"

func main(){
	//打开数据库文件
	db, err := bolt.Open(BLOCKS, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() //xxx.db.lock

	blockChain := chain.CerateChain(db)
	//创世区块
	err = blockChain.CreateGensis([]byte("hello world"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//新增一个区块
	err = blockChain.CreateNewBlock([]byte("hello"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//测试
	lastBlock := blockChain.GetLastBlock()
	fmt.Println("最新区块是:", lastBlock)

	blocks, err := blockChain.GetAllBlocks()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for index, block := range blocks {
		fmt.Printf("第%d个区块：", index)
		fmt.Println(block)
	}

	//db, err := bolt.Open()
	//db.View(func(tx *bolt.Tx) error {
	//		return err
	//	})
	//db.Update(func(tx *bolt.Tx) error {
	//	return err
	//})
}