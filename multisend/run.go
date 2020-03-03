package multisend

import (
	"fmt"
	"github.com/binance-chain/balance-dumper/common"
	"github.com/binance-chain/balance-dumper/utils"
	"github.com/binance-chain/node/app"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	txbuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/crypto/tmhash"
	cmn "github.com/tendermint/tendermint/libs/common"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"time"
)

type Transfer struct {
	To     string `json:"to"`
	Amount string `json:"amount"`
}

type Transfers []Transfer

var batchSize = 300

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run airdrop",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run()
		},
	}

	cmd.Flags().StringP(common.FlagFile, "f", "", "File of Transfer details")
	cmd.Flags().Bool(client.FlagTrustNode, true, "Trust connected full node (don't verify proofs for responses)")
	cmd.Flags().String(client.FlagFrom, "", "Name or address of private key with which to sign")
	cmd.Flags().String(client.FlagMemo, "", "Memo to send along with transaction")
	cmd.Flags().Int64(client.FlagSource, 0, "Source of tx")
	cmd.Flags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	cmd.Flags().String(client.FlagNode, "tcp://localhost:26657", "<host>:<port> to tendermint rpc interface for this chain")
	if err := cmd.MarkFlagRequired(common.FlagFile); err != nil {
		panic(err)
	}

	return cmd
}

func Run() error {

	cdc := app.Codec

	txBuilder := txbuilder.NewTxBuilderFromCLI().WithCodec(cdc)
	ctx := context.NewCLIContext().
		WithCodec(cdc).
		WithAccountDecoder(authcmd.GetAccountDecoder(cdc))

	if err := ctx.EnsureAccountExists(); err != nil {
		return err
	}

	from, err := ctx.GetFromAddress()
	if err != nil {
		return err
	}

	records, err := utils.ReadCSV(viper.GetString(common.FlagFile))
	if err != nil {
		return err
	}

	fmt.Printf("==>Start to run with file: %s\n", viper.GetString(common.FlagFile))
	batchRecords := splitRecords(records, batchSize)
	passphrase := ""
	for count, br := range batchRecords {

		txs := make(Transfers, len(br))
		for i, row := range br {
			transfer := Transfer{row[0], row[1] + ":" + row[2]}
			txs[i] = transfer
		}

		fmt.Printf("==>Start batch %d(from %s to %s)\n", count+1, br[0][0], br[len(br)-1][0])
		msg, err := buildMsg(from, txs)
		if err != nil {
			fmt.Printf("==>Fail to build msg at batch %d\n", count+1)
			return err
		}

		txBytes, err := sign(txBuilder, ctx, []sdk.Msg{msg}, &passphrase)
		if err != nil {
			fmt.Printf("==>Fail to sign at batch %d\n", count+1)
			return err
		}
		txHash := cmn.HexBytes(tmhash.Sum(txBytes)).String()
		fmt.Printf("==>Transaction hash: %s, sending...\n", txHash)

		if res, err := broadcast(ctx, txBytes); err != nil {
			fmt.Printf("==>Sending fail at batch %d, please recheck if the tx(hash: %s) is on chain to prevent repeated transfer\n", count+1, txHash)
			return err
		} else {
			fmt.Printf("==>Sending completed, committed at block %d (tx hash: %s)\n", res.Height, res.Hash.String())
		}

		if count != len(batchRecords)-1 {
			time.Sleep(2 * time.Second)
			fmt.Println()
		}

	}
	return nil
}

func buildMsg(from sdk.AccAddress, txs Transfers) (sdk.Msg, error) {
	if len(txs) == 0 {
		return nil, errors.New("tx is empty")
	}

	toAddrs := make([]sdk.AccAddress, 0, len(txs))
	toCoins := make([]sdk.Coins, 0, len(txs))

	for _, tx := range txs {
		to, err := sdk.AccAddressFromBech32(tx.To)
		if err != nil {
			return nil, err
		}
		toAddrs = append(toAddrs, to)

		toCoin, err := sdk.ParseCoins(tx.Amount)
		if err != nil {
			return nil, err
		}
		toCoins = append(toCoins, toCoin)
	}

	fromCoins := sdk.Coins{}
	for _, toCoin := range toCoins {
		fromCoins = fromCoins.Plus(toCoin)
	}

	if !fromCoins.IsPositive() {
		return nil, errors.Errorf("The number of coins you want to send(%s) should be positive!", fromCoins.String())
	}

	msg := BuildMultiSendMsg(from, fromCoins, toAddrs, toCoins)

	return msg, nil
}

func sign(txBldr txbuilder.TxBuilder, cliCtx context.CLIContext, msgs []sdk.Msg, passphrase *string) ([]byte, error) {
	txBldr, err := prepareTxBuilder(txBldr, cliCtx)
	if err != nil {
		return nil, err
	}
	name, err := cliCtx.GetFromName()
	if err != nil {
		return nil, err
	}

	if passphrase == nil {
		panic("bug found, passphrase can not refer to nil point")
	}
	if len(*passphrase) == 0 {
		*passphrase, err = keys.GetPassphrase(name)
		if err != nil {
			return nil, err
		}

	}

	// build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(name, *passphrase, msgs)

	return txBytes, err
}

func broadcast(cliCtx context.CLIContext, txBytes []byte) (*ctypes.ResultBroadcastTxCommit, error) {
	return cliCtx.BroadcastTxAndAwaitCommit(txBytes)
}

func prepareTxBuilder(txBldr txbuilder.TxBuilder, cliCtx context.CLIContext) (txbuilder.TxBuilder, error) {
	from, err := cliCtx.GetFromAddress()
	if err != nil {
		return txBldr, err
	}

	account, err := cliCtx.GetAccount(from)
	if err != nil {
		return txBldr, err
	}
	txBldr = txBldr.WithAccountNumber(account.GetAccountNumber())
	txBldr = txBldr.WithSequence(account.GetSequence())

	return txBldr, nil
}

func BuildMultiSendMsg(from sdk.AccAddress, fromCoins sdk.Coins, toAddrs []sdk.AccAddress, toCoins []sdk.Coins) sdk.Msg {
	input := bank.NewInput(from, fromCoins)

	output := make([]bank.Output, 0, len(toAddrs))
	for idx, toAddr := range toAddrs {
		output = append(output, bank.NewOutput(toAddr, toCoins[idx]))
	}
	msg := bank.NewMsgSend([]bank.Input{input}, output)
	return msg
}

func splitRecords(records [][]string, batchSize int) [][][]string {
	if records == nil || len(records) == 0 {
		return nil
	}

	var length int
	rs := len(records) / batchSize
	length = rs
	if rmd := len(records) % batchSize; rmd > 0 {
		length = rs + 1
	}

	groupRs := make([][][]string, length)

	for i := 0; i < length; i++ {

		begin := i * batchSize
		if i == length-1 {
			groupRs[i] = records[begin:]
		} else {
			end := (i + 1) * batchSize
			groupRs[i] = records[begin:end]
		}

	}

	return groupRs
}
