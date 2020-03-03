package common

import (
	"github.com/spf13/cobra"
	"os"
)

const (
	FlagFile   = "file"
	FlagHome   = "home"
	FlagOutput = "output"
)

// Executor wraps the cobra Command with a nicer Execute method
type Executor struct {
	*cobra.Command
	Exit func(int) // this is os.Exit by default, override in tests
}

func NewExecutor(cmd *cobra.Command) Executor {
	return Executor{cmd, os.Exit}
}
