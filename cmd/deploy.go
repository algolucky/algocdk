package cmd

import (
	cdktfutil "github.com/algolucky/algocdk/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy stack",
	RunE:  deploy,
}

func deploy(cmd *cobra.Command, args []string) (err error) {
	return cdktfutil.ExecCdktf([]string{"deploy", viper.GetString("stack")})
}
