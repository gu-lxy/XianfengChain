package consensus

import (
	"XianFengChain/utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//目的：拿到区块的属性数据(属性值)
	//1、通过结构体引用，引用block结构体，然后访问

const DIFFICULTY  = 10 //难度值系数

type PoW struct {
	Block   BlockInterface
	Target *big.Int
}


func (pow PoW) FindNonce() ([32]byte,int64){
	fmt.Println("这里是共识机制中使用PoW进行的实现...")

	//1、给定一个nonce值，计算区块hash
	var nonce int64
	nonce = 0

	//无限循环
	hashBig := new(big.Int)
	for {
		hash := CalculateHash(pow.Block, nonce)
		//32 -> 256
		//2、拿到系统的目标值
		target := pow.Target
		//3、比较大小
		hashBig = hashBig.SetBytes(hash[:])
		result := hashBig.Cmp(target)
		//4、判断结果
		if result == -1 {
			return hash,nonce
		}
		nonce++//否则nonce自增
	}
}

/**
 * 根据区块已有的信息和当前nonce的赋值，计算区块的hash
 */
func CalculateHash(block BlockInterface,nonce int64) [32]byte{
	heightByte, _ := utils.Int2Byte(block.GetHeight())
	versionByte, _ := utils.Int2Byte(block.GetVersion())
	timeByte, _ := utils.Int2Byte(block.GetTimeStamp())
	nonceByte, _ := utils.Int2Byte(nonce)
	prve := block.GetPrevHash()

	blockByte := bytes.Join([][]byte{heightByte,versionByte,prve[:],timeByte,nonceByte,block.GetData()},[]byte{})
	//计算区块的哈希
	hash := sha256.Sum256(blockByte)
	return hash
}