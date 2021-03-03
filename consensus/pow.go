package consensus

import (
	"XianfengChain/chain"
	"fmt"
)

type PoW struct {
	Block chain.Block
}


func (pow PoW) FindNonce()int64{
	fmt.Println("这里是共识机制中使用PoW进行的实现...")
	return 0
}