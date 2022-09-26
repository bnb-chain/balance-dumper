package dumper

import (
	"fmt"
	"github.com/binance-chain/balance-dumper/data/account"
	"github.com/binance-chain/balance-dumper/remote/explorer"
	"github.com/binance-chain/balance-dumper/remote/onode"
	"github.com/binance-chain/balance-dumper/utils"
	"github.com/spf13/viper"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

const NoData = "failed to get rootMultiStore: no data"
const KnownErrPrefix = "failed to load rootMultiStore: wanted to load target"
const LogName = "dumper.log"

func AccExport() (err error) {

	targetHeight := viper.GetInt64("height")
	home := viper.GetString("home")
	asset := viper.GetString("asset")
	output := viper.GetString("output")
	dataseed := viper.GetString("dataseed")
	if len(output) == 0 {
		output = home
	}

	if targetHeight <= 0 {
		return fmt.Errorf("error height, must above 0")
	}

	if asset == "" {
		return fmt.Errorf("asset can not be empty")
	}

	_, err = onode.GetTokenInfo(asset, dataseed)
	if err != nil {
		return err
	}

	accs, err := account.Fetch(targetHeight, asset, home)
	export(accs, output)
	return err

}

func launchStateSync(stateSyncHeight int64, home string) {
	if _, err := os.Stat(home); !os.IsNotExist(err) { // 存在
		err = utils.RemoveAllExcept(home, LogName)
		if err != nil {
			panic(err)
		}
	}

	if stateSyncHeight == 0 {
		viper.Set("state_sync_height", -1)
	}

	viper.Set("state_sync_height", stateSyncHeight)
	viper.Set("state_sync_reactor", true)
	viper.Set("recv_rate", 102428800)
	viper.Set("ping_interval", "10m30s")
	viper.Set("pong_timeout", "450s")
}

func LatestBreatheBlockHeight(target int64) (int64, error) {
	if target < 165865 { // first breathe block height
		return 0, nil
	}
	block, err := queryBlock(target)
	if err != nil {
		return -1, err
	}
	bh, ok := breatheBlockHeightFromExplorer(block)

	if ok {
		log.Printf("===>got the block height at 00:00 UTC of the day, %d\n", bh)
		fmt.Printf("===>got the block height at 00:00 UTC of the day, %d\n", bh)
		return bh, nil
	}

	startOfDay := utils.StartOfUTCTime(utils.Milli2Time(block.TimeStamp))
	block, isGreater := loopQuery(block, &startOfDay, 0)

	var curHeight int64
	lastBlock := block
	if isGreater {
		curHeight = lastBlock.BlockHeight - 1
		for {
			curBlock, err := queryBlock(curHeight)
			if err != nil {
				return -1, err
			}
			if !utils.IsSameUTCDay(utils.Milli2Time(lastBlock.TimeStamp), utils.Milli2Time(curBlock.TimeStamp)) {
				bh = lastBlock.BlockHeight
				break
			}
			lastBlock = curBlock
			curHeight--
		}
	} else {
		curHeight = lastBlock.BlockHeight + 1
		for {
			curBlock, err := queryBlock(curHeight)
			if err != nil {
				return -1, err
			}
			if !utils.IsSameUTCDay(utils.Milli2Time(lastBlock.TimeStamp), utils.Milli2Time(curBlock.TimeStamp)) {
				bh = curBlock.BlockHeight
				break
			}
			lastBlock = curBlock
			curHeight++
		}
	}

	log.Printf("===>got the block height at 00:00 UTC of the day, %d\n", bh)
	fmt.Printf("===>got the block height at 00:00 UTC of the day, %d\n", bh)
	return bh, nil

}

var QC int

func queryBlock(height int64) (*explorer.Block, error) {
	QC++
	return explorer.QueryBlock(height)
}

func loopQuery(block *explorer.Block, startOfDay *time.Time, lastEstimatedGapHeight int64) (*explorer.Block, bool) {

	gapMillisecond := float64(block.TimeStamp - utils.Time2Milli(startOfDay))
	if math.Abs(gapMillisecond) <= 1000 {
		return block, gapMillisecond > 0
	}

	estimatedGapHeight := int64(float32(gapMillisecond) * 0.0025)
	if estimatedGapHeight+lastEstimatedGapHeight == 0 {
		estimatedGapHeight = int64(float32(gapMillisecond) * 0.001)
	}
	block, err := queryBlock(block.BlockHeight - estimatedGapHeight)
	if err != nil {
		panic(err)
	}

	return loopQuery(block, startOfDay, estimatedGapHeight)

}

func breatheBlockHeightFromExplorer(block *explorer.Block) (int64, bool) {
	ts := block.TimeStamp

	startOfDay := utils.StartOfUTCTime(utils.Milli2Time(ts))
	at := startOfDay.Add(1 * time.Second)

	blocks, _ := explorer.QueryBlocks(utils.Time2Milli(&startOfDay), utils.Time2Milli(&at), 1)
	if blocks == nil || len(blocks.BlockArray) == 0 {
		return 0, false
	}
	breatheBlock := blocks.BlockArray[len(blocks.BlockArray)-1]
	return breatheBlock.BlockHeight, true
}

func export(accs []account.Balance, output string) {
	data := make([][]string, len(accs))
	if len(accs) > 0 {
		header := []string{"address", "available", "freeze", "in order", "total", "asset"}
		for index, acc := range accs {
			row := []string{acc.Address, strconv.FormatInt(acc.Available, 10),
				strconv.FormatInt(acc.Freeze, 10), strconv.FormatInt(acc.InOrder, 10),
				strconv.FormatInt(acc.Total, 10), acc.Asset}
			data[index] = row
		}
		err := utils.CsvExport(header, data, output, viper.GetString("asset")+"_"+viper.GetString("height")+".csv")
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
