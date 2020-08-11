package onode

import (
	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/common/types"
	ctypes "github.com/binance-chain/go-sdk/common/types"
	tmtypes "github.com/tendermint/tendermint/rpc/core/types"
)

var rpcClient = rpc.NewRPCClient("tcp://dataseed1.binance.org:80", ctypes.ProdNetwork)

func GetTokenInfo(symbol string, dataseed string) (*types.Token, error) {
	if len(dataseed) > 0 {
		rpcClient = rpc.NewRPCClient("tcp://"+dataseed, ctypes.ProdNetwork)
	}
	return rpcClient.GetTokenInfo(symbol)
}

func GetBlock(height *int64, dataseed string) (*tmtypes.ResultBlock, error) {
	if len(dataseed) > 0 {
		rpcClient = rpc.NewRPCClient("tcp://"+dataseed, ctypes.ProdNetwork)
	}
	return rpcClient.Block(height)
}
