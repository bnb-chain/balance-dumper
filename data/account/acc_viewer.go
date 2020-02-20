package account

import (
	"github.com/binance-chain/node/common"
	ntypes "github.com/binance-chain/node/common/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	"github.com/tendermint/tendermint/libs/bech32"
	"github.com/tendermint/tendermint/libs/db"
	dbm "github.com/tendermint/tendermint/libs/db"
	"log"
	"path"
	"sort"
	"strings"
)

const (
	latestVersionKey = "s/latest"
)

var codec = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(codec)
	ntypes.RegisterWire(codec)

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("bnb", "bnbp")
}

func openDB(root, dbName string) *db.GoLevelDB {
	ldb, err := db.NewGoLevelDB(dbName, path.Join(root, "data"))
	if err != nil {
		log.Fatalf("new levelDb err in path %s\n", path.Join(root, "data"))
	}
	return ldb
}

func openAppDB(root string) *db.GoLevelDB {
	return openDB(root, "application")
}

func prepareCms(root string, appDB *db.GoLevelDB, height int64) (cms sdk.CommitMultiStore,err error) {
	keys := common.GetNonTransientStoreKeys()

	cms = store.NewCommitMultiStore(appDB)
	for _, key := range keys {
		cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, nil)
	}
	err = cms.LoadVersion(height)
	if err != nil {
		log.Printf("data does not exist in %s at height = %d\n", root,height)
	}

	return cms,err
}

func accountValueDecoder(value []byte) interface{} {
	acc := ntypes.AppAccount{}
	_ = codec.UnmarshalBinaryBare(value, &acc)
	return acc
}


type Balance struct {
	Address string
	Asset string
	Quantity int64
}

func getLatestVersion(db dbm.DB) int64 {
	var latest int64
	latestBytes := db.Get([]byte(latestVersionKey))
	if latestBytes == nil {
		return 0
	}
	err := codec.UnmarshalBinaryLengthPrefixed(latestBytes, &latest)
	if err != nil {
		panic(err)
	}
	return latest
}

func GetLatestBlockHeight(homePath string) int64{
	ldb := openAppDB(homePath)
	defer ldb.Close()
	return getLatestVersion(ldb)
}

func Fetch(height int64,asset string,homePath string) ([]Balance,error){

	log.Printf("===>start to fetch at height = %d\n",height)
	ldb := openAppDB(homePath)
	defer ldb.Close()

	cms,err := prepareCms(homePath,ldb,height)
	if err != nil {
		return nil,err
	}
	tree := cms.GetCommitStore(common.AccountStoreKey).(store.TreeStore).GetImmutableTree()

	var matchedAccounts []Balance
	tree.Iterate(func(key []byte, value []byte) bool {
		appAcc := accountValueDecoder(value).(ntypes.AppAccount)
		for _,coin := range appAcc.BaseAccount.Coins {
			if strings.Compare(coin.Denom,asset) ==  0 {
				if coin.Amount == 0 {
					break
				}
				bech32Addr,_ := bech32.ConvertAndEncode("bnb",appAcc.Address)
				matchedAccounts = append(matchedAccounts,Balance{bech32Addr,asset,coin.Amount})
				break
			}
		}
		return false
	})

	sort.Slice(matchedAccounts, func(i, j int) bool {
		return matchedAccounts[j].Quantity < matchedAccounts[i].Quantity
	})

	return matchedAccounts,nil

}
