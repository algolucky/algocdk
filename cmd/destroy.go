package cmd

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

var destroyStackID string

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy stack",
	Run:   destroy,
}

func init() {
	rootCmd.AddCommand(destroyCmd)
	destroyCmd.Flags().StringVar(&destroyStackID, "id", "stack", "The ID (or Name) to give the stack")
}

func destroy(cmd *cobra.Command, args []string) {
	binary, err := exec.LookPath("cdktf")
	if err != nil {
		panic(err)
	}

	args = []string{
		"cdktf",
		"destroy",
		deployStackID,
	}

	err = syscall.Exec(binary, args, os.Environ())
	if err != nil {
		panic(err)
	}
}
