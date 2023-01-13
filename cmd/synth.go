package cmd

import (
	"fmt"
	"os"

	"github.com/algolucky/algocdk/stacks/local"
	"github.com/algolucky/algocdk/util"
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

func synth(cmd *cobra.Command, args []string) {
	app := cdktf.NewApp(nil)

	exe, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	util.WriteCdktfJson(&util.CdktfJson{
		Language:         "go",
		App:              exe + " synth --id " + synthStackID,
		SendCrashReports: "false",
	})

	local.LocalStack(app, synthStackID)

	app.Synth()
}
