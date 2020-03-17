package dumper

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccExport(t *testing.T) {
	viper.Set("home", os.ExpandEnv("$HOME/.bdumper"))
	viper.Set("height", 1000)
	viper.Set("asset", "HYN-F21")
	viper.Set("output", os.ExpandEnv("$HOME/.bdumper"))
	err := AccExport()
	assert.NoError(t, err)

}

func TestLatestBreatheHeight(t *testing.T) {
	targetHeight := int64(60909538)
	breatheBlockHeight, err := LatestBreatheBlockHeight(targetHeight)
	assert.NoError(t, err)
	fmt.Printf("query count = %d\n", QC)
	fmt.Printf("result : %d\n", breatheBlockHeight)
}
