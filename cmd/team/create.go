package team

import (
	"fmt"
	"terrakube/client/models"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var createExample string = `Create a new Team
    %[1]v team create --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -n AZB_USER --manage-workspace=true --manage-module=true --manage-provider=true`

var createName string
var createOrgId string
var createManageProvider bool
var createManageModule bool
var createManageWorkspace bool

var createTeamCmd = &cobra.Command{
	Use:   "create",
	Short: "create a Team",
	Run: func(cmd *cobra.Command, args []string) {
		createTeam()
	},
	Example: fmt.Sprintf(createExample, config.CliConfig.CommandName),
}

func createTeam() {
	client := utils.NewClient()

	team := models.Team{
		Attributes: &models.TeamAttributes{
			Name:            createName,
			ManageWorkspace: createManageWorkspace,
			ManageModule:    createManageModule,
			ManageProvider:  createManageProvider,
		},
		Type: "team",
	}

	resp, err := client.Team.Create(createOrgId, team)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
