package workspace

import (
	"fmt"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var listFilter string
var listOrgId string
var listExample string = `List all existing workspaces
    %[1]v workspace list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb
List specific organizations applying a filter
    %[1]v workspace list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb  --filter name==mymodule `
var listWorkspacesCmd = &cobra.Command{
	Use:   "list",
	Short: "list workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		listWorkspaces()
	},
	Example: fmt.Sprintf(listExample, config.CliConfig.CommandName),
}

func listWorkspaces() {
	client := utils.NewClient()
	resp, err := client.Workspace.List(listOrgId, listFilter)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
