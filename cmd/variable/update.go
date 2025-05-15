package variable

import (
	"fmt"
	"terrakube/client/models"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var updateExample string = `Update the value of the variable using id
    %[1]v workspace variable update --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd --id 38b6635a-d38e-46f2-a95e-d00a416de4fd -v "new value" `

var updateId string
var updateKey string
var updateValue string
var updateDescription string
var updateCategory string
var updateSensitive bool
var updateHcl bool
var updateOrgId string
var updateWorkspaceId string

var updateVariableCmd = &cobra.Command{
	Use:   "update",
	Short: "update a variable",
	Run: func(cmd *cobra.Command, args []string) {
		updateVariable()
	},
	Example: fmt.Sprintf(updateExample, config.CliConfig.CommandName),
}

func updateVariable() {
	client := utils.NewClient()

	variable := models.Variable{
		Attributes: &models.VariableAttributes{
			Key:         updateKey,
			Value:       updateValue,
			Description: updateDescription,
			Sensitive:   updateSensitive,
			Hcl:         updateHcl,
			Category:    updateCategory,
		},
		ID:   updateId,
		Type: "variable",
	}
	err := client.Variable.Update(updateOrgId, updateWorkspaceId, variable)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Updated")

}
