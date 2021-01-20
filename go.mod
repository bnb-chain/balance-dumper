module github.com/binance-chain/balance-dumper

go 1.12

require (
	github.com/binance-chain/go-sdk v1.2.3
	github.com/binance-chain/node v0.8.1
	github.com/cosmos/cosmos-sdk v0.25.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.4.0
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/iavl v0.12.4
	github.com/tendermint/tendermint v0.32.3
	google.golang.org/appengine v1.4.0 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/binance-chain/bnc-cosmos-sdk v0.25.0-binance.24
	github.com/tendermint/go-amino => github.com/binance-chain/bnc-go-amino v0.14.1-binance.2
	github.com/tendermint/iavl => github.com/binance-chain/bnc-tendermint-iavl v0.12.0-binance.3
	github.com/tendermint/tendermint => github.com/binance-chain/bnc-tendermint v0.32.3-binance.5
	github.com/zondax/ledger-cosmos-go => github.com/binance-chain/ledger-cosmos-go v0.9.9-binance.3
	golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20190823143015-45b1026d81ae
)
