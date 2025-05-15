package organization

import (
	"fmt"
	"terrakube/client/models"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var updateExample string = `Update the description of the organization using id
    %[1]v organization update --id 38b6635a-d38e-46f2-a95e-d00a416de4fd -d "new description" `

var UpdateOrganizationCmd = &cobra.Command{
	Use:   "update",
	Short: "update an organization",
	Run: func(cmd *cobra.Command, args []string) {
		updateOrganization()
	},
	Example: fmt.Sprintf(updateExample, config.CliConfig.CommandName),
}

var updateId string
var updateDescription string
var updateName string

func updateOrganization() {
	client := utils.NewClient()

	organization := models.Organization{
		Attributes: &models.OrganizationAttributes{
			Name:        updateName,
			Description: updateDescription,
		},
		Type: "organization",
		ID:   updateId,
	}
	err := client.Organization.Update(organization)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Updated")

}
