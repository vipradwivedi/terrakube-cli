package variable

import (
	"fmt"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var listFilter string
var listOrgId string
var listWorkspaceId string
var listListExample string = `List all existing variables for a workspace
    %[1]v workspace variable list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd
List specific variable applying a filter
    %[1]v workspace variable list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd --filter key==myvariable `

var listVariablesCmd = &cobra.Command{
	Use:   "list",
	Short: "list variables",
	Run: func(cmd *cobra.Command, args []string) {
		listVariables()
	},
	Example: fmt.Sprintf(listListExample, config.CliConfig.CommandName),
}

func listVariables() {
	client := utils.NewClient()
	resp, err := client.Variable.List(listOrgId, listWorkspaceId, listFilter)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
