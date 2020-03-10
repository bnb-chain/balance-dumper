package onode

import (
	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/common/types"
	ctypes "github.com/binance-chain/go-sdk/common/types"
	tmtypes "github.com/tendermint/tendermint/rpc/core/types"
)

var rpcClient = rpc.NewRPCClient("tcp://data-seed-pre-0-s3.binance.org:80", ctypes.TestNetwork)

func GetTokenInfo(symbol string) (*types.Token, error) {
	return rpcClient.GetTokenInfo(symbol)
}

func GetBlock(height *int64) (*tmtypes.ResultBlock, error) {
	return rpcClient.Block(height)
}
