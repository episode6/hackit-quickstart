package main

import (
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

const defaultProjectVersion = "0.0.1-SNAPSHOT"
const defaultLicenseName = "The MIT License (MIT)"

func defaultConfigFilePath() string {
	userdir, err := homedir.Dir()
	if err != nil {
		return ""
	}
	configFilePath, err := homedir.Expand(filepath.Join(userdir, ".hackit-quickstart"))
	if err != nil {
		return ""
	}
	fileInfo, err := os.Stat(configFilePath)
	if err != nil {
		return ""
	}
	if fileInfo.Mode().IsRegular() {
		return configFilePath
	}
	return ""
}
