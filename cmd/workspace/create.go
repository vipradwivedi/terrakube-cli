package workspace

import (
	"fmt"
	"terrakube/client/models"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var createExample string = `Create a new workspace
    %[1]v workspace create --organization-id 312b4415-806b-47a9-9452-b71f0753136e -n myWorkspace -s https://github.com/AzBuilder/terraform-sample-repository.git -b master -t 0.15.0`

var createName string
var createSource string
var createBranch string
var createTerraformV string
var createOrgId string
var createWorkspaceCmd = &cobra.Command{
	Use:   "create",
	Short: "create a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		createWorkspace()
	},
	Example: fmt.Sprintf(createExample, config.CliConfig.CommandName),
}

func createWorkspace() {
	client := utils.NewClient()

	workspace := models.Workspace{
		Attributes: &models.WorkspaceAttributes{
			Name:             createName,
			Source:           createSource,
			Branch:           createBranch,
			TerraformVersion: createTerraformV,
		},
		Type: "workspace",
	}

	resp, err := client.Workspace.Create(createOrgId, workspace)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
