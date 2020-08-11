package onode

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTokenInfo(t *testing.T) {
	token, err := GetTokenInfo("BNB", "")
	assert.NoError(t, err)
	bz, err := json.Marshal(token)
	fmt.Println(string(bz))
}

func TestGetBlock(t *testing.T) {
	height := int64(70337812)
	block, err := GetBlock(&height, "")
	assert.NoError(t, err)
	bk, err := json.Marshal(block)
	fmt.Println(string(bk))
}
