package organization

import (
	"fmt"
	"terrakube/client/models"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var createExample string = `Create a new organization
    %[1]v organization create -n myorg -d "org description" `

var createName string
var createDescription string
var createOrganizationCmd = &cobra.Command{
	Use:   "create",
	Short: "create an organization",
	Run: func(cmd *cobra.Command, args []string) {
		createOrganization()
	},
	Example: fmt.Sprintf(createExample, config.CliConfig.CommandName),
}

func createOrganization() {
	client := utils.NewClient()

	organization := models.Organization{
		Attributes: &models.OrganizationAttributes{
			Name:        createName,
			Description: createDescription,
		},
		Type: "organization",
	}
	resp, err := client.Organization.Create(organization)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
