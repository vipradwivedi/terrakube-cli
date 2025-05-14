package module

import (
	"fmt"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var filter string
var orgId string
var listExample string = `List all existing modules
    %[1]v module list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb
List specific organizations applying a filter
    %[1]v module list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb --filter name==mymodule `

var listModulesCmd = &cobra.Command{
	Use:   "list",
	Short: "list modules",
	Run: func(cmd *cobra.Command, args []string) {
		listModules()
	},
	Example: fmt.Sprintf(listExample, config.CliConfig.CommandName),
}

func listModules() {
	client := utils.NewClient()
	resp, err := client.Module.List(orgId, filter)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
