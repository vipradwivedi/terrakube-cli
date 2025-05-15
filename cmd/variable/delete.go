package variable

import (
	"fmt"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var deleteExample string = `Delete a variable
    %[1]v workspace variable delete --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb  -w 38b6635a-d38e-46f2-a95e-d00a416de4fd --id 38b6635a-d38e-46f2-a95e-d00a416de4fd `

var deleteId string
var deleteOrgId string
var deleteWorkspaceId string

var deleteVariableCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a variable",
	Run: func(cmd *cobra.Command, args []string) {
		deleteVariable()
	},
	Example: fmt.Sprintf(deleteExample, config.CliConfig.CommandName),
}

func deleteVariable() {
	client := utils.NewClient()

	err := client.Variable.Delete(deleteOrgId, deleteWorkspaceId, deleteId)

	if err != nil {
		fmt.Println(err)
		return
	}
}
