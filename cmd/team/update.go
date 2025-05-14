package team

import (
	"fmt"
	"terrakube/client/models"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var updateExample string = `Update the permissions of the Team using id
    %[1]v team update --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb --id 38b6635a-d38e-46f2-a95e-d00a416de4fd --manage-workspace true --manage-module true --manage-provider true" `

var updateId string
var updateName string
var updateOrgId string
var updateManageProvider bool
var updateManageModule bool
var updateManageWorkspace bool

var updateTeamCmd = &cobra.Command{
	Use:   "update",
	Short: "update a Team",
	Run: func(cmd *cobra.Command, args []string) {
		updateTeam()
	},
	Example: fmt.Sprintf(updateExample, config.CliConfig.CommandName),
}

func updateTeam() {
	client := utils.NewClient()

	team := models.Team{
		Attributes: &models.TeamAttributes{
			Name:            updateName,
			ManageWorkspace: updateManageWorkspace,
			ManageModule:    updateManageModule,
			ManageProvider:  updateManageProvider,
		},
		ID:   updateId,
		Type: "Team",
	}
	err := client.Team.Update(updateOrgId, team)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Updated")

}
