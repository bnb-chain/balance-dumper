package explorer

import (
	"encoding/json"
	"github.com/binance-chain/balance-dumper/remote"
	"strconv"
)

const (
	uri  = "https://explorer.binance.org/api/v1/"
	blocksEndpoint  = "blocks"
	blockEndpoint = "block"
)

type Blocks struct {
	Total  int64   `json:"total"`
	BlockArray []Block `json:"blockArray"`
}

type Block struct {
	BlockHeight int64 `json:"blockHeight"`
	BlockHash string `json:"blockHash"`
	TxNum int `json:"txNum"`
	ProposalAddr string `json:"proposalAddr"`
	ParentHash string `json:"parentHash"`
	TimeStamp int64 `json:"timeStamp"`
}

func QueryBlocks(startTime int64,endTime int64,page int) (*Blocks,error){

	url := uri + blocksEndpoint
	params := make(map[string]string,3)
	params["page"] = strconv.Itoa(page)
	params["rows"] = "50"
	params["startTime"] = strconv.FormatInt(startTime,10)
	params["endTime"] = strconv.FormatInt(endTime,10)
	var blocks = Blocks{}
	respB,err := remote.Get(url,params)
	if err != nil {
		return nil,err
	}
	if respB != nil {
		err = json.Unmarshal(respB,&blocks)
		if err != nil {
			return nil,err
		}
	}

	return &blocks,err
}

func QueryBlock(blockHeight int64) (*Block,error) {

	url := uri + blockEndpoint + "/" + strconv.FormatInt(blockHeight,10)
	respB,err := remote.Get(url,nil)
	if err != nil {
		return nil,err
	}

	var block = Block{}
	if respB != nil {
		err = json.Unmarshal(respB,&block)
		if err != nil {
			return nil,err
		}
	}

	return &block,err

}
