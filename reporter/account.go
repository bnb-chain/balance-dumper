package reporter

import (
	"fmt"
	"github.com/binance-chain/acc-tool/data/account"
	"github.com/binance-chain/acc-tool/node"
	"github.com/binance-chain/acc-tool/remote/explorer"
	"github.com/binance-chain/acc-tool/remote/onode"
	"github.com/binance-chain/acc-tool/utils"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
	"time"
)

const NoData = "failed to get rootMultiStore: no data"
const LogName  = "reporter.log"

func AccExport() (err error){

	targetHeight := viper.GetInt64("height")
	home := viper.GetString("home")
	asset := viper.GetString("asset")


	if targetHeight <= 0 {
		return fmt.Errorf("error height, must above 0")
	}

	if asset == "" {
		return fmt.Errorf("asset can not be empty")
	}

	_,err = onode.GetTokenInfo(asset)
	if err != nil {
		return err
	}

	latestHeight := account.GetLatestBlockHeight(home)
	var breatheHeight int64

	var accs []account.Balance
	if targetHeight > latestHeight {
		breatheHeight = latestBreatheBlockHeight(targetHeight)
		if breatheHeight - latestHeight > 15000 { // state sync mode
			launchStateSync(breatheHeight,home)
		} else { // fast sync
			viper.Set("state_sync_height",-1)
		}
		err = node.Start(targetHeight)
	} else {
		accs,err = account.Fetch(targetHeight,asset,home)
		if err == nil {
			export(accs)
			return nil
		}

		if err.Error() == NoData { // data does not exist in DB
			breatheHeight = latestBreatheBlockHeight(targetHeight)
			launchStateSync(breatheHeight,home)
			err = node.Start(targetHeight)
		}
	}

	if err == nil {
		accs,err = account.Fetch(targetHeight,asset,home)
		export(accs)
	}

	return err

}

func launchStateSync(stateSyncHeight int64,home string) {
	if _,err:= os.Stat(home); !os.IsNotExist(err) { // 存在
		err = utils.RemoveAllExcept(home,LogName)
		if err != nil {
			panic(err)
		}
	}

	viper.Set("state_sync_height",stateSyncHeight)
	viper.Set("state_sync_reactor",true)
	viper.Set("recv_rate",102428800)
	viper.Set("ping_interval","10m30s")
	viper.Set("pong_timeout","450s")
}

func latestBreatheBlockHeight(target int64) int64 {
	block,_ := explorer.QueryBlock(target)
	ts := block.TimeStamp


	startOfDay := utils.StartOfUTCTime(utils.Milli2Time(ts))
	at := startOfDay.Add(1 * time.Second)

	blocks,_ := explorer.QueryBlocks(utils.Time2Milli(startOfDay),utils.Time2Milli(at),1)
	breatheBlock := blocks.BlockArray[len(blocks.BlockArray)-1]
	log.Printf("===>got the latest breathe block height = %d\n",breatheBlock.BlockHeight)
	return breatheBlock.BlockHeight
}


func export(accs []account.Balance) {
	data := make([][]string,len(accs))
	if accs != nil && len(accs) > 0 {
		header := []string{"address","asset","balance"}
		for index,acc := range accs {
			row := []string{acc.Address,acc.Asset,strconv.FormatInt(acc.Quantity,10)}
			data[index] = row
		}
		err := csvExport(header,data,viper.GetString("output"),viper.GetString("asset") + "_" + viper.GetString("height") + ".csv")
		if err != nil {
			panic(err)
		}
		log.Printf("===>finish fetching,got %d matched accounts\n\n", len(accs))
		fmt.Printf("===>finish fetching,got %d matched accounts\n\n", len(accs))
	} else {
		log.Printf("===>finish fetching,no matching records\n\n")
		fmt.Printf("===>finish fetching,No matching records\n\n")
	}

}
