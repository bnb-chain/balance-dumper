package onode

import (
	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/common/types"
	ctypes "github.com/binance-chain/go-sdk/common/types"
	tmtypes "github.com/tendermint/tendermint/rpc/core/types"
)

var rpcClient = rpc.NewRPCClient("tcp://seed1.ciscox.io:80", ctypes.ProdNetwork)

func GetTokenInfo(symbol string) (*types.Token,error){
	return rpcClient.GetTokenInfo(symbol)
}

func GetBlock(height *int64) (*tmtypes.ResultBlock, error){
	return rpcClient.Block(height)
}
