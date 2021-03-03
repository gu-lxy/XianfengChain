package chain

import (
	"XianFengChain/utils"
	"bytes"
	"crypto/sha256"
	"time"
)

const VERSION = 0x00

/**
 * 区块的结构体定义
 */
type Block struct {
	Height int64
	Version int64
	PrevHash [32]byte
	//当前区块的hash
	Hash [32]byte
	//默克尔根
	TimeStamp int64
	//Difficulty int
	Nonce int64
	//区块体
	Data []byte
}

/**
 * 计算哈希值并进行赋值
 */
func (block *Block) CalculateBlockHash() {
	heightByte, _ := utils.Int2Byte(block.Height)
	versionByte, _ := utils.Int2Byte(block.Version)
	timeByte, _ := utils.Int2Byte(block.TimeStamp)
	nonceByte, _ := utils.Int2Byte(block.Nonce)

	blockByte := bytes.Join([][]byte{heightByte,versionByte,block.PrevHash[:],timeByte,nonceByte,block.Data},[]byte{})
	//委屈快的hash字段赋值
	block.Hash = sha256.Sum256(blockByte)
}

/**
* 生成创世区块的函数
 */
func CreateGenesis(data []byte) Block{
	genesis := Block{
		Height:0,
		Version: VERSION,
		PrevHash: [32]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		TimeStamp: time.Now().Unix(),
		Data: data,
	}
	//todo 计算并设置hash  寻找并设置nonce
	genesis.CalculateBlockHash()

	return genesis
}


/**
 * 生成新区块的功能函数
 */
func NewBlock(height int64,prev [32]byte,data []byte) Block{
	newBlock := Block{
		Height: height+1,
		Version: VERSION,
		PrevHash: prev,
		TimeStamp: time.Now().Unix(),
		Data: data,
	}
	//todo 设置哈希 寻找并设置nonce
	newBlock.CalculateBlockHash()
	return newBlock
}