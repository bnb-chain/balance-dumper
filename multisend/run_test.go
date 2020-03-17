package multisend

import (
	"fmt"
	"github.com/binance-chain/balance-dumper/common"
	"github.com/binance-chain/balance-dumper/utils"
	"github.com/binance-chain/node/app"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	ctx := app.ServerContext

	config := sdk.GetConfig()
	if app.Bech32PrefixAccAddr != "" {
		ctx.Bech32PrefixAccAddr = app.Bech32PrefixAccAddr
	}
	config.SetBech32PrefixForAccount("tbnb", ctx.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(ctx.Bech32PrefixValAddr, ctx.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(ctx.Bech32PrefixConsAddr, ctx.Bech32PrefixConsPub)
	config.Seal()

	viper.Set(client.FlagNode, "data-seed-pre-0-s1.binance.org:80")
	viper.Set(client.FlagChainID, "Binance-Chain-Nile")
	viper.Set(client.FlagTrustNode, true)
	viper.Set(client.FlagFrom, "validator1")
	viper.Set(common.FlagHome, os.ExpandEnv("$HOME/.tbairdrop"))
	viper.Set(common.FlagFile, "example.csv")
	batchSize = 3
	err := Run()
	assert.NoError(t, err)
}

func TestSplitRecords(t *testing.T) {
	records, err := utils.ReadCSV("example.csv")
	assert.NoError(t, err)
	groupRs := splitRecords(records, 3)
	for i := 0; i < len(groupRs); i++ {
		fmt.Println(groupRs[i])
	}

}
