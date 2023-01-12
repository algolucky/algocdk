package cmd

import (
	"github.com/algolucky/algocdk/stacks"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/spf13/cobra"
)

var stackID string

var synthCmd = &cobra.Command{
	Use:   "synth",
	Short: "",
	Long:  ``,
	Run:   synth,
}

func init() {
	rootCmd.AddCommand(synthCmd)
	synthCmd.Flags().StringVar(&stackID, "id", "stack", "The ID (or Name) to give the stack")
}

func synth(cmd *cobra.Command, args []string) {
	app := cdktf.NewApp(nil)

	stacks.LocalStack(app, stackID)

	app.Synth()
}
