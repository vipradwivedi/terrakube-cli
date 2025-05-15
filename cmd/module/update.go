package module

import (
	"fmt"
	"terrakube/client/models"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var updateExample string = `Update the description of the module using id
    %[1]v module update --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb --id 38b6635a-d38e-46f2-a95e-d00a416de4fd -d "new description" `

var updateId string
var updateDescription string
var updateName string
var updateOrgId string
var updateSource string
var updateProvider string

var updateModuleCmd = &cobra.Command{
	Use:   "update",
	Short: "update a module",
	Run: func(cmd *cobra.Command, args []string) {
		updateModule()
	},
	Example: fmt.Sprintf(updateExample, config.CliConfig.CommandName),
}

func updateModule() {
	client := utils.NewClient()

	module := models.Module{
		Attributes: &models.ModuleAttributes{
			Name:        updateName,
			Description: updateDescription,
			Source:      updateSource,
			Provider:    updateProvider,
		},
		ID:   updateId,
		Type: "module",
	}
	err := client.Module.Update(updateOrgId, module)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Updated")

}
