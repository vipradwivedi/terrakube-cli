package workspace

import (
	"fmt"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var deleteExample string = `Delete a workspace
    %[1]v workspace delete --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb --id 38b6635a-d38e-46f2-a95e-d00a416de4fd `

var deleteId string
var deleteOrgId string
var deleteWorkspaceCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		deleteWorkspace()
	},
	Example: fmt.Sprintf(deleteExample, config.CliConfig.CommandName),
}

func deleteWorkspace() {
	client := utils.NewClient()

	err := client.Workspace.Delete(deleteOrgId, deleteId)

	if err != nil {
		fmt.Println(err)
		return
	}
}
