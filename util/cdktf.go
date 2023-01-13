package util

import (
	"encoding/json"
	"os"
)

type CdktfJson struct {
	Language         string `json:"language"`
	App              string `json:"app"`
	SendCrashReports string `json:"sendCrashReports"`
}

func WriteCdktfJson(cdktfJson *CdktfJson) (err error) {
	file, err := json.MarshalIndent(cdktfJson, "", "  ")
	if err != nil {
		return err
	}
	os.WriteFile("cdktf.json", file, 0644)
	return
}
