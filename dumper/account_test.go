package dumper

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestAccExport(t *testing.T) {
	viper.Set("home", os.ExpandEnv("$HOME/.tbdumper"))
	viper.Set("height", 1000)
	viper.Set("asset", "HYN-F21")
	viper.Set("output", os.ExpandEnv("$HOME/.tbdumper"))
	err := AccExport()
	if err != nil {
		t.Error(err)
	}
}

func TestLatestBreatheHeight(t *testing.T) {
	targetHeight := int64(60909538)
	breatheBlockHeight, err := LatestBreatheBlockHeight(targetHeight)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("query count = %d\n", QC)
	fmt.Printf("result : %d\n", breatheBlockHeight)
}
