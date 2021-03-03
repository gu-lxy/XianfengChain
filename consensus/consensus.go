package consensus

import "XianfengChain/chain"

type Consensus interface {
	FindNonce() int64
}

func NewPoW(block chain.Block) Consensus{
	return PoW{block}
}

func NewPoS(block chain.Block) Consensus{
	return PoS{Block:block}
}

