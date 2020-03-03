package main

import (
	"fmt"
	"github.com/binance-chain/balance-dumper/common"
	"github.com/binance-chain/balance-dumper/multisend"
	"github.com/binance-chain/node/app"
	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	DefaultCLIHome = os.ExpandEnv("$HOME/.tbairdrop")

	airdropCmd = &cobra.Command{
		Use:   "tbairdrop",
		Short: "Testnet BlockChain Airdrop",
	}

	addressPrefix = "tbnb"
)

func main() {

	// disable sorting
	cobra.EnableCommandSorting = false

	ctx := app.ServerContext

	config := sdk.GetConfig()
	if app.Bech32PrefixAccAddr != "" {
		ctx.Bech32PrefixAccAddr = app.Bech32PrefixAccAddr
	}
	config.SetBech32PrefixForAccount(addressPrefix, ctx.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(ctx.Bech32PrefixValAddr, ctx.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(ctx.Bech32PrefixConsAddr, ctx.Bech32PrefixConsPub)
	config.Seal()

	airdropCmd.AddCommand(keys.Commands())
	airdropCmd.AddCommand(multisend.Command())

	airdropCmd.PersistentFlags().StringP(common.FlagHome, "", DefaultCLIHome, "directory for config and data")

	airdropCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		return nil
	}

	executor := common.NewExecutor(airdropCmd)
	err := executor.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
