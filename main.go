/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/algolucky/algocdk/cmd"
)

func main() {
	// cdktf doesn't need cli
	if _, present := os.LookupEnv("CDKTF_OUTDIR"); !present {
		cmd.Execute()
	}
}
