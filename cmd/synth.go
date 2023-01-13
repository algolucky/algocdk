package cmd

import (
	cdktfutil "github.com/algolucky/algocdk/util"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var synthCmd = &cobra.Command{
	Use:   "synth",
	Short: "Synthesize a stack",
	RunE:  synth,
}

func synth(cmd *cobra.Command, args []string) (err error) {
	app := cdktf.NewApp(nil)

	err = cdktfutil.WriteCdktfJson(viper.GetString("stack"))
	if err != nil {
		return err
	}

	// stack := cdktf.NewTerraformStack(app, &stackID)
	// stacks.ContainerStackSimple(stack, config.C.Container)

	app.Synth()

	return
}
