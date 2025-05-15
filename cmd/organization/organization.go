package organization

import (
	"github.com/spf13/cobra"
)

var organizationLong = `
This command consists of multiple subcommands to interact with organizations.
It can be used to create, update, delete and list organizations.
`

var Cmd = &cobra.Command{
	Use:     "organization create|update|delete|list [ARGS]",
	Short:   "create, update, delete and list organizations",
	Long:    organizationLong,
	Aliases: []string{"org"},
}

func init() {
	Cmd.AddCommand(createOrganizationCmd)
	createOrganizationCmd.Flags().StringVarP(&createName, "name", "n", "", "Name of the new organization (required)")
	_ = createOrganizationCmd.MarkFlagRequired("name")
	createOrganizationCmd.Flags().StringVarP(&createDescription, "description", "d", "", "Description of the new organization")

	Cmd.AddCommand(UpdateOrganizationCmd)
	UpdateOrganizationCmd.Flags().StringVarP(&updateId, "id", "", "", "Id of the organization (required)")
	_ = UpdateOrganizationCmd.MarkFlagRequired("id")
	UpdateOrganizationCmd.Flags().StringVarP(&updateName, "name", "n", "", "Name of the organization")
	UpdateOrganizationCmd.Flags().StringVarP(&updateDescription, "description", "d", "", "Description of the organization")

	Cmd.AddCommand(deleteOrganizationCmd)
	deleteOrganizationCmd.Flags().StringVarP(&deleteId, "id", "", "", "Id of the organization (required)")
	_ = deleteOrganizationCmd.MarkFlagRequired("id")

	Cmd.AddCommand(listOrganizationsCmd)
	listOrganizationsCmd.Flags().StringVarP(&listFilter, "filter", "f", "", "Filter")
}
