package cmd

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

var deployStackID string

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy stack",
	Run:   deploy,
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringVar(&deployStackID, "id", "stack", "The ID (or Name) to give the stack")
}

func deploy(cmd *cobra.Command, args []string) {
	binary, err := exec.LookPath("cdktf")
	if err != nil {
		panic(err)
	}

	args = []string{
		"cdktf",
		"deploy",
		deployStackID,
	}

	err = syscall.Exec(binary, args, os.Environ())
	if err != nil {
		panic(err)
	}
}
