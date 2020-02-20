package onode

import (
	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/common/types"
	ctypes "github.com/binance-chain/go-sdk/common/types"
)

var rpcClient = rpc.NewRPCClient("tcp://seed1.ciscox.io:80", ctypes.ProdNetwork)

func GetTokenInfo(symbol string) (*types.Token,error){
	return rpcClient.GetTokenInfo(symbol)
}
