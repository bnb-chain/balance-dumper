package main

import (
	"fmt"
	"github.com/binance-chain/balance-dumper/common"
	"github.com/binance-chain/balance-dumper/multisend"
	"github.com/binance-chain/node/app"
	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	DefaultCLIHome = os.ExpandEnv("$HOME/.bairdrop")

	airdropCmd = &cobra.Command{
		Use:   "bairdrop",
		Short: "BlockChain Airdrop",
	}

	addressPrefix = "bnb"
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
	airdropCmd.PersistentFlags().StringP(common.FlagOutput, "o", "text", "Output format (text|json)")

	airdropCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {

		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		// validate output format
		output := viper.GetString(common.FlagOutput)
		switch output {
		case "text", "json":
		default:
			return errors.Errorf("Unsupported output format: %s", output)
		}

		return nil
	}

	executor := common.NewExecutor(airdropCmd)
	err := executor.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
