package util

import (
	"encoding/json"
	"os"
	"os/exec"
	"syscall"
)

type CdktfJson struct {
	Language         string `json:"language"`
	App              string `json:"app"`
	SendCrashReports string `json:"sendCrashReports"`
}

func ExecCdktf(args []string) (err error) {
	binary, err := exec.LookPath("cdktf")
	if err != nil {
		return err
	}

	cdktfArgs := append([]string{binary}, args...)

	err = syscall.Exec(binary, cdktfArgs, os.Environ())
	if err != nil {
		return err
	}
	return
}

func WriteCdktfJson(stack string) (err error) {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	cdktfJson := &CdktfJson{
		Language:         "go",
		App:              exe + " synth --stack " + stack,
		SendCrashReports: "false",
	}

	file, err := json.MarshalIndent(cdktfJson, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("cdktf.json", file, 0644)
	if err != nil {
		return err
	}

	return
}
