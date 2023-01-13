package cmd

import (
	cdktfutil "github.com/algolucky/algocdk/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var outputsCmd = &cobra.Command{
	Use:   "outputs",
	Short: "Print stack outputs",
	RunE:  outputs,
}

func outputs(cmd *cobra.Command, args []string) (err error) {
	return cdktfutil.ExecCdktf([]string{"deploy", viper.GetString("stack")})
}
