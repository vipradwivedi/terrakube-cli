package organization

import (
	"fmt"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var listFilter string

var listExample string = `List all existing organizations
    %[1]v organization list
List specific organizations applying a filter
    %[1]v organization list --filter name==myorg `

var listOrganizationsCmd = &cobra.Command{
	Use:   "list",
	Short: "list organizations",
	Run: func(cmd *cobra.Command, args []string) {
		listOrganizations()
	},
	Example: fmt.Sprintf(listExample, config.CliConfig.CommandName),
}

func listOrganizations() {
	client := utils.NewClient()
	resp, err := client.Organization.List(listFilter)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
