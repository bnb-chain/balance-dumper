package account

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestGetLatestHeight(t *testing.T) {
	homePath := os.ExpandEnv("$HOME/.bdumper")
	height := GetLatestBlockHeight(homePath)
	fmt.Printf("latest height is %d", height)
}

func TestFetch(t *testing.T) {
	height := int64(68515735)
	asset := "bnb"
	homePath := os.ExpandEnv("$HOME/.bdumper")
	accounts, err := Fetch(height, asset, homePath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("got %d accounts", len(accounts))
}

func TestGetAccount(t *testing.T) {
	homePath := os.ExpandEnv("$HOME/.bdumper")
	acc, err := GetAccount(58571291, homePath, "bnb1ag3rpe9lten7fhyqg4cde9qusrv3dv67lsshup")
	if err != nil {
		t.Error(err)
	}
	jsonValue, _ := json.Marshal(acc)
	fmt.Printf("%s\n\n", string(jsonValue))

}
