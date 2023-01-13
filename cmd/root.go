/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/algolucky/algocdk/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "algocdk",
	Short: "Algorand Cloud Development Kit",
	Long:  `Quickly spin up Algorand infrastructure in the cloud.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(synthCmd)
	rootCmd.AddCommand(outputsCmd)
	rootCmd.AddCommand(destroyCmd)

	rootCmd.PersistentFlags().String("config", "", "config file (default is ./algocdk.yaml")
	rootCmd.PersistentFlags().String("stack", "", "stack name")

	viper.BindPFlag("stack", rootCmd.Flags().Lookup("stack"))
}

func initConfig() {
	config.InitConfig(configFile)
	err := config.Unmarshal()
	if err != nil {
		fmt.Println(err)
	}
}
