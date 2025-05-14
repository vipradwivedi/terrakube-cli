package workspace

import (
	"github.com/spf13/cobra"
)

var description = `
This command consists of multiple subcommands to interact with workspaces.
It can be used to create, update, delete and list workspaces.
`

var Cmd = &cobra.Command{
	Use:     "workspace create|update|delete|list [ARGS]",
	Short:   "create, update, delete and list workspaces",
	Long:    description,
	Aliases: []string{"wrk"},
}

func init() {
	Cmd.AddCommand(createWorkspaceCmd)
	createWorkspaceCmd.Flags().StringVarP(&createName, "name", "n", "", "Name of the new workspace (required)")
	_ = createWorkspaceCmd.MarkFlagRequired("name")
	createWorkspaceCmd.Flags().StringVarP(&createOrgId, "organization-id", "", "", "Id of the organization (required)")
	_ = createWorkspaceCmd.MarkFlagRequired("organization-id")
	createWorkspaceCmd.Flags().StringVarP(&createBranch, "branch", "b", "", "Branch of the new workspace")
	createWorkspaceCmd.Flags().StringVarP(&createSource, "source", "s", "", "Source of the new workspace")
	createWorkspaceCmd.Flags().StringVarP(&createTerraformV, "terraform-version", "t", "", "Terraform Version use in the new workspace")

	Cmd.AddCommand(deleteWorkspaceCmd)
	deleteWorkspaceCmd.Flags().StringVarP(&deleteOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = deleteWorkspaceCmd.MarkFlagRequired("organization-id")
	deleteWorkspaceCmd.Flags().StringVarP(&deleteId, "id", "", "", "Id of the workspace (required)")
	_ = deleteWorkspaceCmd.MarkFlagRequired("id")

	Cmd.AddCommand(listWorkspacesCmd)
	listWorkspacesCmd.Flags().StringVarP(&listFilter, "filter", "f", "", "Filter")
	listWorkspacesCmd.Flags().StringVarP(&listOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = listWorkspacesCmd.MarkFlagRequired("organization-id")

	Cmd.AddCommand(updateWorkspaceCmd)
	updateWorkspaceCmd.Flags().StringVarP(&updateName, "name", "n", "", "Name of the workspace (required)")
	updateWorkspaceCmd.Flags().StringVarP(&updateOrgId, "organization-id", "", "", "Id of the organization (required)")
	_ = updateWorkspaceCmd.MarkFlagRequired("organization-id")
	updateWorkspaceCmd.Flags().StringVarP(&updateId, "id", "", "", "Id of the workspace (required)")
	_ = updateWorkspaceCmd.MarkFlagRequired("id")
	updateWorkspaceCmd.Flags().StringVarP(&updateBranch, "branch", "b", "", "Branch of the workspace")
	updateWorkspaceCmd.Flags().StringVarP(&updateSource, "source", "s", "", "Source of the workspace")
	updateWorkspaceCmd.Flags().StringVarP(&updateTerraformV, "terraform-version", "t", "", "Terraform Version use in the workspace")
}
