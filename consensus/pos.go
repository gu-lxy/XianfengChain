package consensus

import (
	"XianfengChain/chain"
	"fmt"
)

type PoS struct {
	Block chain.Block
}

func (pos PoS) FindNonce() int64{
	fmt.Println("这是共识机制使用PoS算法的实现")
	return  0
}
