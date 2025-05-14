package variable

import (
	"github.com/spf13/cobra"
)

var description = `
This command consists of multiple subcommands to interact with variables.
It can be used to create, update, delete and list variables within a workspace.
`

var Cmd = &cobra.Command{
	Use:     "variable create|update|delete|list [ARGS]",
	Short:   "create, update, delete and list variables",
	Long:    description,
	Aliases: []string{"var"},
}

func init() {
	Cmd.AddCommand(createVariableCmd)
	createVariableCmd.Flags().StringVarP(&createKey, "key", "k", "", "Key of the new variable (required)")
	_ = createVariableCmd.MarkFlagRequired("key")
	createVariableCmd.Flags().StringVarP(&createOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = createVariableCmd.MarkFlagRequired("organization-id")
	createVariableCmd.Flags().StringVarP(&createValue, "value", "v", "", "Value for the new variable")
	_ = createVariableCmd.MarkFlagRequired("value")
	createVariableCmd.Flags().StringVarP(&createDescription, "description", "d", "", "Description of the new variable")
	createVariableCmd.Flags().StringVarP(&createCategory, "category", "c", "", "Category of the new variable. Valid values are TERRAFORM or ENV")
	_ = createVariableCmd.MarkFlagRequired("category")
	createVariableCmd.Flags().BoolVarP(&createSensitive, "sensitive", "s", false, "Whether the value is sensitive. If true then the variable is written once and not visible thereafter.")
	_ = createVariableCmd.MarkFlagRequired("sensitive")
	createVariableCmd.Flags().BoolVarP(&createHcl, "hcl", "", false, "Whether to evaluate the value of the variable as a string of HCL code. Has no effect for environment variables.")
	_ = createVariableCmd.MarkFlagRequired("hcl")
	createVariableCmd.Flags().StringVarP(&createWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = createVariableCmd.MarkFlagRequired("workspace-id")

	Cmd.AddCommand(deleteVariableCmd)
	deleteVariableCmd.Flags().StringVarP(&deleteOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = deleteVariableCmd.MarkFlagRequired("organization-id")
	deleteVariableCmd.Flags().StringVarP(&deleteId, "id", "", "", "Id of the variable (required)")
	_ = deleteVariableCmd.MarkFlagRequired("id")
	deleteVariableCmd.Flags().StringVarP(&deleteWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = deleteVariableCmd.MarkFlagRequired("workspace-id")

	Cmd.AddCommand(listVariablesCmd)
	listVariablesCmd.Flags().StringVarP(&listFilter, "filter", "f", "", "Filter")
	listVariablesCmd.Flags().StringVarP(&listOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = listVariablesCmd.MarkFlagRequired("organization-id")
	listVariablesCmd.Flags().StringVarP(&listWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = listVariablesCmd.MarkFlagRequired("workspace-id")

	Cmd.AddCommand(updateVariableCmd)
	//updateVariableCmd.AddCommand(organization.UpdateOrganizationCmd)
	updateVariableCmd.Flags().StringVarP(&updateId, "id", "", "", "Id of the variable (required)")
	_ = updateVariableCmd.MarkFlagRequired("id")
	updateVariableCmd.Flags().StringVarP(&updateKey, "key", "k", "", "Key of the variable")
	updateVariableCmd.Flags().StringVarP(&updateValue, "value", "v", "", "Value of the variable")
	updateVariableCmd.Flags().StringVarP(&updateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = updateVariableCmd.MarkFlagRequired("organization-id")
	updateVariableCmd.Flags().StringVarP(&updateWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = updateVariableCmd.MarkFlagRequired("workspace-id")
	updateVariableCmd.Flags().StringVarP(&updateDescription, "description", "d", "", "Description of the variable")
	updateVariableCmd.Flags().StringVarP(&updateCategory, "category", "c", "", "Category of the variable. Valid values are TERRAFORM or ENV")
	updateVariableCmd.Flags().BoolVarP(&updateSensitive, "sensitive", "s", false, "Whether the value is sensitive. If true then the variable is written once and not visible thereafter.")
	updateVariableCmd.Flags().BoolVarP(&updateHcl, "hcl", "", false, "Whether to evaluate the value of the variable as a string of HCL code. Has no effect for environment variables.")

}
