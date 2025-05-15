package team

import (
	"fmt"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var listFilter string
var orgId string
var listExample string = `List all existing teams
    %[1]v team list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb
List specific team organizations applying a filter
    %[1]v team list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb --filter name==myteam `

var listTeamsCmd = &cobra.Command{
	Use:   "list",
	Short: "list teams",
	Run: func(cmd *cobra.Command, args []string) {
		listTeams()
	},
	Example: fmt.Sprintf(listExample, config.CliConfig.CommandName),
}

func listTeams() {
	client := utils.NewClient()
	resp, err := client.Team.List(orgId, listFilter)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
