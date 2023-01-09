module github.com/binance-chain/balance-dumper

go 1.12

require (
	github.com/binance-chain/go-sdk v1.2.9
	github.com/bnb-chain/node v0.10.6
	github.com/cosmos/cosmos-sdk v0.25.0
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/iavl v0.12.4
	github.com/tendermint/tendermint v0.32.3
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bnb-chain/bnc-cosmos-sdk v0.25.6
	github.com/tendermint/go-amino => github.com/bnb-chain/bnc-go-amino v0.14.1-binance.2
	github.com/tendermint/iavl => github.com/bnb-chain/bnc-tendermint-iavl v0.12.0-binance.4
	github.com/tendermint/tendermint => github.com/bnb-chain/bnc-tendermint v0.32.3-bc.8
	github.com/zondax/ledger-cosmos-go => github.com/bnb-chain/ledger-cosmos-go v0.9.9-binance.3
	golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20190823183015-45b1026d81ae
)
