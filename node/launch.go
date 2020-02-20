package node

import (
	"github.com/binance-chain/node/app"
	"github.com/binance-chain/node/common/tx"
	"github.com/binance-chain/node/common/types"
	"github.com/binance-chain/node/common/upgrade"
	"github.com/binance-chain/node/plugins/account"
	"github.com/binance-chain/node/plugins/dex"
	"github.com/binance-chain/node/plugins/param"
	"github.com/binance-chain/node/plugins/tokens"
	"github.com/binance-chain/node/wire"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/concurrent"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/stake"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	tlog "github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	pvm "github.com/tendermint/tendermint/privval"
	tmTypes "github.com/tendermint/tendermint/types"
	"io"
	"log"
	"time"
)

var Codec = MakeCodec()


// MakeCodec creates a custom tx codec.
func MakeCodec() *wire.Codec {
	var cdc = wire.NewCodec()

	wire.RegisterCrypto(cdc) // Register crypto.
	bank.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc) // Register Msgs
	dex.RegisterWire(cdc)
	tokens.RegisterWire(cdc)
	account.RegisterWire(cdc)
	types.RegisterWire(cdc)
	tx.RegisterWire(cdc)
	stake.RegisterCodec(cdc)
	gov.RegisterCodec(cdc)
	param.RegisterWire(cdc)
	return cdc
}



type XBinanceChain struct {
	*app.BinanceChain
	stopAt int64
	stopSignal chan bool
}

func (app *XBinanceChain) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) (res abci.ResponseBeginBlock) {
	upgrade.Mgr.BeginBlocker(ctx)
	if ctx.BlockHeight() > app.stopAt {
		app.stopSignal <- true
	}
	return
}

func newBinanceChain(logger tlog.Logger, db dbm.DB, storeTracer io.Writer,stopAt int64) *XBinanceChain {

	bc := app.NewBinanceChain(logger, db, storeTracer)
	xc := &XBinanceChain{bc,stopAt,make(chan bool)}
	bc.SetBeginBlocker(xc.BeginBlocker)

	return xc
}


func Start(stopAt int64) (err error) {

	ctx := app.ServerContext

	err = prepare(ctx)
	if err != nil {
		return err
	}

	log.Printf("===>start node,home = %s, stopAt = %d, StateSyncHeight = %d\n",ctx.Config.RootDir,stopAt,ctx.Config.StateSyncHeight)

	cfg := ctx.Config

	dbProvider := node.DefaultDBProvider
	db, err := dbProvider(&node.DBContext{"application", cfg})
	if err != nil {
		return err
	}

	bncApp := newBinanceChain(ctx.Logger,db,nil,stopAt)

	nodeKey, err := p2p.LoadOrGenNodeKey(cfg.NodeKeyFile())
	if err != nil {
		return err
	}

	cliCreator := concurrent.NewAsyncLocalClientCreator(bncApp,
		ctx.Logger.With("module", "abciCli"))


	genesisDocProvider := func() (*tmTypes.GenesisDoc, error){
		genesisDoc,err := tmTypes.GenesisDocFromJSON([]byte(GenesisJson))
		if err != nil {
			return nil,err
		}
		return genesisDoc,nil
	}

	// create & start tendermint node
	tmNode, err := node.NewNode(
		cfg,
		pvm.LoadOrGenFilePV(cfg.PrivValidatorKeyFile(), cfg.PrivValidatorStateFile()),
		nodeKey,
		cliCreator,
		genesisDocProvider,
		//node.DefaultGenesisDocProviderFunc(cfg),
		dbProvider,
		node.DefaultMetricsProvider(cfg.Instrumentation),
		ctx.Logger.With("module", "node"),
	)

	if err != nil {
		return err
	}

	err = tmNode.Start()
	if err != nil {
		return err
	}

	log.Printf("===>node started from height = %d\n",tmNode.BlockStore().Height())
	log.Println("===>syncing......")

	server.TrapSignal(func() {
		if tmNode.IsRunning() {
			_ = tmNode.Stop()
		}
	})


	select {
		case <- bncApp.stopSignal:
			log.Printf("===>node catches up the target height %d, terminal the node\n", stopAt)
			if tmNode.IsRunning() {
				err = tmNode.Stop()
				close(bncApp.stopSignal)
				if err != nil {
					return err
				}
				time.Sleep(1 * time.Second)
				db.Close()
			}
	}

	return nil
}







