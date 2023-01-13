package cmd

import (
	cdktfutil "github.com/algolucky/algocdk/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy stack",
	RunE:  destroy,
}

func destroy(cmd *cobra.Command, args []string) (err error) {
	return cdktfutil.ExecCdktf([]string{"destroy", viper.GetString("stack")})
}
