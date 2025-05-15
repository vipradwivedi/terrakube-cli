package job

import (
	"fmt"
	"terrakube/client/models"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var createExample string = `Create a new job
    %[1]v job create --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w e5ad0642-f9b3-48b3-9bf4-35997febe1fb  -c apply`

var workspaceId string
var jobCreateCommand string
var jobCreateOrgId string

var createJobCmd = &cobra.Command{
	Use:   "create",
	Short: "create a job",
	Run: func(cmd *cobra.Command, args []string) {
		createJob()
	},
	Example: fmt.Sprintf(createExample, config.CliConfig.CommandName),
}

func createJob() {
	client := utils.NewClient()

	job := models.Job{
		Attributes: &models.JobAttributes{
			Command: jobCreateCommand,
		},
		Type: "job",
		Relationships: &models.JobRelationships{
			Workspace: &models.JobRelationshipsWorkspace{
				Data: &models.JobRelationshipsWorkspaceData{
					Type: "workspace",
					ID:   workspaceId,
				},
			},
		},
	}

	resp, err := client.Job.Create(jobCreateOrgId, job)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
