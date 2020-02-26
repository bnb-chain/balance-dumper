package main

import (
	"fmt"
	"github.com/binance-chain/balance-dumper/dumper"
	"github.com/binance-chain/balance-dumper/node"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

// Executor wraps the cobra Command with a nicer Execute method
type Executor struct {
	*cobra.Command
	Exit func(int) // this is os.Exit by default, override in tests
}


func main(){
	rootCmd := &cobra.Command{
		Use:               "bdumper",
		Short:             "Balance Dumper",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dumper.AccExport()
		},
		PersistentPreRunE: globalConfig,
	}

	rootCmd.PersistentFlags().String("home", os.ExpandEnv("$HOME/.bdumper"), "directory for config and data")
	rootCmd.PersistentFlags().Int64("height",0,"query height ")
	rootCmd.PersistentFlags().String("asset","","query asset ")
	rootCmd.PersistentFlags().StringP("output","o",os.ExpandEnv("$HOME/.bdumper"),"directory for storing the csv file of balance result")

	executor := Executor{rootCmd, os.Exit}
	err := executor.Execute()
	if err != nil {
		fmt.Println(err)
	}
}

// 1.Bind all flags and read the config into viper
// 2.Configure log file
func globalConfig(cmd *cobra.Command, args []string) error {
	// cmd.Flags() includes flags from this command and all persistent flags from the parent
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	homeDir := viper.GetString("home")
	viper.Set("home", homeDir)
	viper.SetConfigName("config")                         // name of config file (without extension)
	viper.AddConfigPath(homeDir)                          // search root directory
	viper.AddConfigPath(filepath.Join(homeDir, "config")) // search root directory /config

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// stderr, so if we redirect output to json file, this doesn't appear
		// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		// ignore not found error, return other errors
		return err
	}

	if _,err := os.Stat(homeDir); err != nil && os.IsNotExist(err) {
		err = os.Mkdir(homeDir,node.DefaultDirPerm)
		if err != nil {
			return err
		}
	}
	logf,err := os.OpenFile(filepath.Join(homeDir, dumper.LogName),os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644);if err != nil {
		return err
	}
	log.SetOutput(logf)
	log.SetPrefix("[Dumper]")
	return nil
}

