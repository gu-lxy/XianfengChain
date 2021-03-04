package chain

//import beego "beego-develop"

/**
 * 定义区块链结构体，该结构体用于管理区块
 */
type BlockChain struct {
	Blocks []Block
}


/**
 * 创建一个创世区块，包含一个创世区块
 */
func CreateChainWithGensis(data []byte) BlockChain{
	gensis := CreateGenesis(data)
	blocks := make([]Block, 0)
	blocks = append(blocks, gensis)
	return BlockChain{blocks}
}

func (chain *BlockChain) CreateNewBlock(data []byte)  {
	blocks := chain.Blocks //获取到当前所有区块
	lastBlock := blocks[len(blocks) - 1] //最后最新的区块
	newBlock := NewBlock(lastBlock.Height, lastBlock.Hash, data)
	chain.Blocks = append(chain.Blocks,newBlock)
}