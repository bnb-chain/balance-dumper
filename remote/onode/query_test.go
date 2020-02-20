package onode

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTokenInfo(t *testing.T) {
	token,err := GetTokenInfo("BNB2")
	assert.NoError(t, err)
	bz, err := json.Marshal(token)
	fmt.Println(string(bz))
}
