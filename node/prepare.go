package node

import (
	"fmt"
	"github.com/binance-chain/node/app/config"
	bnclog "github.com/binance-chain/node/common/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
	tmcfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	"os"
	"path"
	"path/filepath"
)

const DefaultDirPerm = 0700

var (
	defaultConfigDir     = "config"
 	defaultDataDir       = "data"
 	defaultConfigFileName  = "config.toml"
 	defaultConfigFilePath   = filepath.Join(defaultConfigDir, defaultConfigFileName)
)

func prepare(ctx *config.BinanceChainContext) (err error) {

	err = interceptLoadConfigInPlace(ctx)

	if err != nil {
		return err
	}

	sdkConfig := sdk.GetConfig()
	sdkConfig.SetBech32PrefixForAccount("tbnb", ctx.Bech32PrefixAccPub)
	sdkConfig.SetBech32PrefixForValidator(ctx.Bech32PrefixValAddr, ctx.Bech32PrefixValPub)
	sdkConfig.SetBech32PrefixForConsensusNode(ctx.Bech32PrefixConsAddr, ctx.Bech32PrefixConsPub)
	sdkConfig.Seal()

	logger := newLogger(ctx)
	logger, err = tmflags.ParseLogLevel(ctx.Config.LogLevel, logger, tmcfg.DefaultLogLevel())
	if err != nil {
		return err
	}
	if viper.GetBool(cli.TraceFlag) {
		logger = log.NewTracingLogger(logger)
	}
	logger = logger.With("module", "main")
	bnclog.InitLogger(logger)

	ctx.Logger = logger

	return nil

}

func interceptLoadConfigInPlace(ctx *config.BinanceChainContext) (err error) {
	ctx.Config,err = parseConfig() //new config,read viper to config. create config file
	if err != nil {
		return err
	}

	appConfigFilePath := filepath.Join(ctx.Config.RootDir, "config/", config.AppConfigFileName+".toml")
	if _, err := os.Stat(appConfigFilePath); os.IsNotExist(err) {

		customizedAppFile(ctx)

		config.WriteConfigFile(appConfigFilePath, ctx.BinanceChainConfig)
	} else {
		err = ctx.ParseAppConfigInPlace()
		if err != nil {
			return err
		}
	}
	return nil
}

func parseConfig() (*tmcfg.Config, error) {
	conf := tmcfg.DefaultConfig()
	err := viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	customizedConfigFile(conf)

	conf.SetRoot(conf.RootDir)
	ensureRoot(conf.RootDir,conf)
	if err = conf.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("Error in config file: %v", err)
	}
	return conf, err
}

// EnsureRoot creates the root, config, and data directories if they don't exist,
// and panics if it fails.
func ensureRoot(rootDir string, conf *tmcfg.Config) {
	if err := cmn.EnsureDir(rootDir, DefaultDirPerm); err != nil {
		panic(err.Error())
	}
	if err := cmn.EnsureDir(filepath.Join(rootDir, defaultConfigDir), DefaultDirPerm); err != nil {
		panic(err.Error())
	}
	if err := cmn.EnsureDir(filepath.Join(rootDir, defaultDataDir), DefaultDirPerm); err != nil {
		panic(err.Error())
	}

	configFilePath := filepath.Join(rootDir, defaultConfigFilePath)

	tmcfg.WriteConfigFile(configFilePath,conf)

}



func customizedAppFile(ctx *config.BinanceChainContext) {
	ctx.BinanceChainConfig.BEP6Height = 24020000
	ctx.BinanceChainConfig.BEP9Height = 24020000
	ctx.BinanceChainConfig.BEP10Height = 24020000
	ctx.BinanceChainConfig.BEP19Height = 24020000
	ctx.BinanceChainConfig.BEP12Height = 29794000
	ctx.BinanceChainConfig.BEP3Height = 39581000
	ctx.BinanceChainConfig.FixSignBytesOverflowHeight = 49721000
	ctx.BinanceChainConfig.LotSizeUpgradeHeight = 49721000
	ctx.BinanceChainConfig.ListingRuleUpgradeHeight = 49721000
	ctx.BinanceChainConfig.FixZeroBalanceHeight = 49721000
	ctx.BinanceChainConfig.LogToConsole = false
}

func customizedConfigFile(conf *tmcfg.Config) {
	conf.ProxyApp = "tcp://127.0.0.1:28658"
	conf.HotSync = true
	conf.HotSyncReactor = true
	conf.ProfListenAddress = ":9060"

	conf.RPC.ListenAddress = "tcp://0.0.0.0:27147"

	conf.P2P.ListenAddress = "tcp://0.0.0.0:27146"
	conf.P2P.Seeds = "2726550182cbc5f4618c27e49c730752a96901e8@a41086771245011e988520ad55ba7f5a-5f7331395e69b0f3.elb.us-east-1.amazonaws.com:27146,34ac6eb6cd914014995b5929be8d7bc9c16f724d@aa13359cd244f11e988520ad55ba7f5a-c3963b80c9b991b7.elb.us-east-1.amazonaws.com:27146,fe5eb5a945598476abe4826a8d31b9f8da7b1a54@aa35ed7c1244f11e988520ad55ba7f5a-bbfb4fe79dee5d7e.elb.us-east-1.amazonaws.com:27146,8825b32e3abec71d772abf009ba1956d452be1fa@aa58a7e44244f11e988520ad55ba7f5a-45d504e63bacb8dd.elb.us-east-1.amazonaws.com:27146"
	conf.P2P.AddrBookStrict = false
	conf.P2P.MaxPacketMsgPayloadSize = 10485760
	conf.P2P.KeysPerRequest = 1500
	conf.P2P.RecvRate = 102428800

	conf.DBCache.BlockCacheCapacity = 1073741824
	conf.DBCache.WriteBuffer = 67108864

	conf.Consensus.SkipTimeoutCommit = true

	conf.Instrumentation.PrometheusListenAddr = ":28660"
}

func newLogger(ctx *config.BinanceChainContext) log.Logger {
	if ctx.LogConfig.LogToConsole {
		return bnclog.NewConsoleLogger()
	} else {
		logFilePath := ""
		if ctx.LogConfig.LogFileRoot == "" {
			logFilePath = path.Join(ctx.Config.RootDir, ctx.LogConfig.LogFilePath)
		} else {
			logFilePath = path.Join(ctx.LogConfig.LogFileRoot, ctx.LogConfig.LogFilePath)
		}
		err := cmn.EnsureDir(path.Dir(logFilePath), 0755)
		if err != nil {
			panic(fmt.Sprintf("create log dir failed, err=%s", err.Error()))
		}
		return bnclog.NewAsyncFileLogger(logFilePath, ctx.LogConfig.LogBuffSize)
	}
}
