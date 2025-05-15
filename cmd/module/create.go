package module

import (
	"fmt"
	"terrakube/client/models"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var createExample string = `Create a new module
    %[1]v module create --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -n myModule -d "module description" -p azurerm -s https://github.com/AzBuilder/terraform-sample-repository.git `

var createName string
var createDescription string
var createOrgId string
var createSource string
var createProvider string

var createModuleCmd = &cobra.Command{
	Use:   "create",
	Short: "create a module",
	Run: func(cmd *cobra.Command, args []string) {
		createModule()
	},
	Example: fmt.Sprintf(createExample, config.CliConfig.CommandName),
}

func createModule() {
	client := utils.NewClient()

	module := models.Module{
		Attributes: &models.ModuleAttributes{
			Description:  createDescription,
			Name:         createName,
			Source:       createSource,
			SourceSample: createSource,
			Provider:     createProvider,
		},
		Type: "module",
	}

	resp, err := client.Module.Create(createOrgId, module)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
