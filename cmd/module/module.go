package module

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var description = `
This command consists of multiple subcommands to interact with modules.
It can be used to create, update, delete and list modules.
`

var Cmd = &cobra.Command{
	Use:     "module create|update|delete|list [ARGS]",
	Short:   "create, update, delete and list modules",
	Long:    description,
	Aliases: []string{"mod"},
}

func init() {
	_ = viper.BindEnv("organization-id", "TERRAKUBE_ORGANIZATION_ID")

	Cmd.AddCommand(createModuleCmd)
	createModuleCmd.Flags().StringVarP(&createName, "name", "n", "", "Name of the new module (required)")
	_ = createModuleCmd.MarkFlagRequired("name")
	createModuleCmd.Flags().StringVarP(&createOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = createModuleCmd.MarkFlagRequired("organization-id")
	createModuleCmd.Flags().StringVarP(&createDescription, "description", "d", "", "Description of the new module")
	createModuleCmd.Flags().StringVarP(&createSource, "source", "s", "", "Source of the new module")
	createModuleCmd.Flags().StringVarP(&createProvider, "provider", "p", "", "Provider of the new module")

	Cmd.AddCommand(deleteModuleCmd)
	deleteModuleCmd.Flags().StringVarP(&deleteOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = deleteModuleCmd.MarkFlagRequired("organization-id")
	deleteModuleCmd.Flags().StringVarP(&deleteId, "id", "", "", "Id of the module (required)")
	_ = deleteModuleCmd.MarkFlagRequired("id")

	Cmd.AddCommand(listModulesCmd)
	listModulesCmd.Flags().StringVarP(&filter, "filter", "f", "", "Filter")
	listModulesCmd.Flags().StringVarP(&orgId, "organization-id", "", "", "Organization Id (required)")
	_ = listModulesCmd.MarkFlagRequired("organization-id")

	Cmd.AddCommand(updateModuleCmd)
	//updateModuleCmd.AddCommand(organization.UpdateOrganizationCmd)
	updateModuleCmd.Flags().StringVarP(&updateId, "id", "", "", "Id of the module (required)")
	_ = updateModuleCmd.MarkFlagRequired("id")
	updateModuleCmd.Flags().StringVarP(&updateName, "name", "n", "", "Name of the module")
	updateModuleCmd.Flags().StringVarP(&updateDescription, "description", "d", "", "Description of the module")
	updateModuleCmd.Flags().StringVarP(&updateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = updateModuleCmd.MarkFlagRequired("organization-id")
	updateModuleCmd.Flags().StringVarP(&updateSource, "source", "s", "", "Source of the module")
	updateModuleCmd.Flags().StringVarP(&updateProvider, "provider", "p", "", "Provider of the module")
}
