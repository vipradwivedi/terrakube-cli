package variable

import (
	"fmt"
	"terrakube/client/models"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var createExample string = `Create a new variable
    %[1]v workspace variable create --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd -k tag_name -v "Hola mundo" --hcl=false --sensitive=false --category TERRAFORM `

var createKey string
var createValue string
var createDescription string
var createCategory string
var createSensitive bool
var createHcl bool
var createOrgId string
var createWorkspaceId string

var createVariableCmd = &cobra.Command{
	Use:   "create",
	Short: "create a variable",
	Run: func(cmd *cobra.Command, args []string) {
		createVariable()
	},
	Example: fmt.Sprintf(createExample, config.CliConfig.CommandName),
}

func createVariable() {
	client := utils.NewClient()

	variable := models.Variable{
		Attributes: &models.VariableAttributes{
			Key:         createKey,
			Value:       createValue,
			Description: createDescription,
			Sensitive:   createSensitive,
			Hcl:         createHcl,
			Category:    createCategory,
		},
		Type: "variable",
	}

	resp, err := client.Variable.Create(createOrgId, createWorkspaceId, variable)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
