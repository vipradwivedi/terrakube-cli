package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type Config struct {
	CommandName        string
	ConfigFileLocation string
	OutputFormat       string
	DefaultApiVersion  string
}

var CliConfig Config

func init() {
	CliConfig.CommandName = "terrakube"

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	CliConfig.ConfigFileLocation = filepath.Join(home, ".terrakube-cli.yaml")

	CliConfig.OutputFormat = "json"
	CliConfig.DefaultApiVersion = "/api/v1/"
}
