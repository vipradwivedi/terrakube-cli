package workspace

import (
	"fmt"
	"terrakube/client/models"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var updateExample string = `Update Terraform version in workspace
    %[1]v workspace update --organization-id 312b4415-806b-47a9-9452-b71f0753136e --id 38b6635a-d38e-46f2-a95e-d00a416de4fd -t 0.14.0 `

var updateName string
var updateSource string
var updateBranch string
var updateTerraformV string
var updateOrgId string
var updateId string

var updateWorkspaceCmd = &cobra.Command{
	Use:   "update",
	Short: "update a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		updateWorkspace()
	},
	Example: fmt.Sprintf(updateExample, config.CliConfig.CommandName),
}

func updateWorkspace() {
	client := utils.NewClient()

	workspace := models.Workspace{
		Attributes: &models.WorkspaceAttributes{
			Name:             updateName,
			Branch:           updateBranch,
			Source:           updateSource,
			TerraformVersion: updateTerraformV,
		},
		Type: "workspace",
	}

	err := client.Workspace.Update(updateOrgId, workspace)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Updated")
}
