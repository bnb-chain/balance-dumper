package dumper

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestAccExport(t *testing.T) {
	viper.Set("home",os.ExpandEnv("$HOME/.bnbaccr"))
	viper.Set("height",68515742)
	viper.Set("asset","HYN-F21")
	viper.Set("output",os.ExpandEnv("$HOME/.bnbaccr"))
	err := AccExport()
	if err != nil {
		t.Error(err)
	}
}

func TestLatestBreatheHeight(t *testing.T){
	targetHeight := int64(68515745)
	rs := latestBreatheBlockHeight(targetHeight)
	fmt.Printf("result : %d\n",rs)
}
