package organization

import (
	"fmt"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var deleteId string
var deleteExample string = `Delete an organization using id
    %[1]v organization delete --id 38b6635a-d38e-46f2-a95e-d00a416de4fd`

var deleteOrganizationCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an organization",
	Run: func(cmd *cobra.Command, args []string) {
		deleteOrganization()
	},
	Example: fmt.Sprintf(deleteExample, config.CliConfig.CommandName),
}

func deleteOrganization() {
	client := utils.NewClient()

	err := client.Organization.Delete(deleteId)

	if err != nil {
		fmt.Println(err)
		return
	}
}
