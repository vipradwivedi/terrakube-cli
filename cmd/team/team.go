package team

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var teamLong = `
This command consists of multiple subcommands to interact with modules.
It can be used to create, update, delete and list teams.
`

var Cmd = &cobra.Command{
	Use:   "team create|update|delete|list [ARGS]",
	Short: "create, update, delete and list teams",
	Long:  teamLong,
}

func init() {
	_ = viper.BindEnv("organization-id", "TERRAKUBE_ORGANIZATION_ID")

	Cmd.AddCommand(updateTeamCmd)
	//updateTeamCmd.AddCommand(organization.UpdateOrganizationCmd)
	updateTeamCmd.Flags().StringVarP(&updateId, "id", "", "", "Id of the Team (required)")
	_ = updateTeamCmd.MarkFlagRequired("id")
	updateTeamCmd.Flags().StringVarP(&updateName, "name", "n", "", "Name of the Team")
	updateTeamCmd.Flags().StringVarP(&updateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = updateTeamCmd.MarkFlagRequired("organization-id")
	updateTeamCmd.Flags().BoolVarP(&updateManageProvider, "manage-provider", "", false, "Manage Provider Permissions")
	updateTeamCmd.Flags().BoolVarP(&updateManageModule, "manage-module", "", false, "Manage Module Permissions")
	updateTeamCmd.Flags().BoolVarP(&updateManageWorkspace, "manage-workspace", "", false, "Manage Workspaces Permissions")

	Cmd.AddCommand(listTeamsCmd)
	listTeamsCmd.Flags().StringVarP(&listFilter, "filter", "f", "", "Filter")
	listTeamsCmd.Flags().StringVarP(&orgId, "organization-id", "", "", "Organization Id (required)")
	_ = listTeamsCmd.MarkFlagRequired("organization-id")

	Cmd.AddCommand(deleteTeamCmd)
	deleteTeamCmd.Flags().StringVarP(&deleteOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = deleteTeamCmd.MarkFlagRequired("organization-id")
	deleteTeamCmd.Flags().StringVarP(&deleteId, "id", "", "", "Id of the Team (required)")
	_ = deleteTeamCmd.MarkFlagRequired("id")

	Cmd.AddCommand(createTeamCmd)
	createTeamCmd.Flags().StringVarP(&createName, "name", "n", "", "Name of the new Team (required)")
	_ = createTeamCmd.MarkFlagRequired("name")
	createTeamCmd.Flags().StringVarP(&createOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = createTeamCmd.MarkFlagRequired("organization-id")
	createTeamCmd.Flags().BoolVarP(&createManageProvider, "manage-provider", "", false, "Manage Provider Permissions")
	_ = createTeamCmd.MarkFlagRequired("manage-provider")
	createTeamCmd.Flags().BoolVarP(&createManageModule, "manage-module", "", false, "Manage Module Permissions")
	_ = createTeamCmd.MarkFlagRequired("manage-module")
	createTeamCmd.Flags().BoolVarP(&createManageWorkspace, "manage-workspace", "", false, "Manage Workspaces Permissions")
	_ = createTeamCmd.MarkFlagRequired("manage-workspace")

}
