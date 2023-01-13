package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/algolucky/algocdk/stacks/local"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/spf13/cobra"
)

var synthStackID string

var synthCmd = &cobra.Command{
	Use:   "synth",
	Short: "Synthesize a stack",
	Run:   synth,
}

func init() {
	rootCmd.AddCommand(synthCmd)
	synthCmd.Flags().StringVar(&synthStackID, "id", "stack", "The ID (or Name) to give the stack")
}

type cdktfJson struct {
	Language         string `json:"language"`
	App              string `json:"app"`
	SendCrashReports string `json:"sendCrashReports"`
}

func synth(cmd *cobra.Command, args []string) {
	app := cdktf.NewApp(nil)

	exe, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	cdktfJson := &cdktfJson{
		Language:         "go",
		App:              exe + " synth --id " + synthStackID,
		SendCrashReports: "false",
	}

	file, err := json.MarshalIndent(cdktfJson, "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("cdktf.json", file, 0644)

	local.LocalStack(app, synthStackID)

	app.Synth()
}
