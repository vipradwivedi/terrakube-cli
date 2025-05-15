package cmd

import (
	"fmt"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/pflag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"terrakube/cmd/job"
	"terrakube/cmd/module"
	"terrakube/cmd/organization"
	"terrakube/cmd/team"
	"terrakube/cmd/variable"
	"terrakube/cmd/workspace"
	"terrakube/config"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var output string
var envPrefix string = "TERRAKUBE"

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "terrakube",
	Short: "terrakube command line tool",
	Long: `
terrakube is a CLI to handle remote terraform workspace and modules in organizations 
and handle all the lifecycle (plan, apply, destroy).`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set the global config before any command runs
		config.CliConfig.ConfigFileLocation = cfgFile
		config.CliConfig.OutputFormat = output
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
	err := doc.GenMarkdownTree(RootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.terrakube-cli.yaml)")
	_ = viper.BindPFlag("root.config", RootCmd.Flags().Lookup("config"))

	RootCmd.PersistentFlags().StringVar(&output, "output", "json", "Use json, table, tsv, markdown or none to format CLI output")
	_ = viper.BindPFlag("root.output", RootCmd.Flags().Lookup("output"))

	_ = RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	RootCmd.AddCommand(job.Cmd)
	RootCmd.AddCommand(organization.Cmd)
	RootCmd.AddCommand(module.Cmd)
	RootCmd.AddCommand(team.Cmd)
	RootCmd.AddCommand(variable.Cmd)
	RootCmd.AddCommand(workspace.Cmd)

	RootCmd.AddCommand(&cobra.Command{
		Use:    "docs",
		Short:  "Generate documentation",
		Hidden: true, // This makes the command hidden
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generating documentation...")
			// Ensure the docs directory exists
			if err := os.MkdirAll("./docs", 0755); err != nil {
				fmt.Printf("Error creating docs directory: %v\n", err)
				os.Exit(1)
			}
			err := doc.GenMarkdownTree(RootCmd, "./docs")
			if err != nil {
				fmt.Printf("Error generating docs: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Documentation generated in ./docs")
		},
	})

	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgCyan).SprintFunc())
	usageTemplate := RootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Available Commands:`, `{{StyleHeading "Commands:"}}`,
		`Global Flags:`, `{{StyleHeading "Global Flags:"}}`,
		`Examples:`, `{{StyleHeading "Examples:"}}`,
	).Replace(usageTemplate)

	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{StyleHeading "Flags:"}}`)
	RootCmd.SetUsageTemplate(usageTemplate)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configFile := filepath.Join(home, ".terrakube-cli.yaml")
		viper.SetConfigFile(configFile)
	}

	viper.SetEnvPrefix(envPrefix)
	_ = viper.BindEnv("workspace-id", "TERRAKUBE_WORKSPACE_ID")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	postInitCommands(RootCmd.Commands())
}

func postInitCommands(commands []*cobra.Command) {
	for _, cmd := range commands {
		presetRequiredFlags(cmd)
		if cmd.HasSubCommands() {
			postInitCommands(cmd.Commands())
		}
	}
}

func presetRequiredFlags(cmd *cobra.Command) {
	_ = viper.BindPFlags(cmd.Flags())
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
			_ = cmd.Flags().Set(f.Name, viper.GetString(f.Name))
		}
	})
}
