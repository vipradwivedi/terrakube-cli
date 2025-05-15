package job

import (
	"fmt"
	"terrakube/cmd/utils"
	"terrakube/config"

	"github.com/spf13/cobra"
)

var filters string
var orgId string
var listExample string = `List all existing jobs
    %[1]v job list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb
List specific jobs applying a filter
    %[1]v job list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb --filter id==jobid `

var listJobsCmd = &cobra.Command{
	Use:   "list",
	Short: "list jobs",
	Run: func(cmd *cobra.Command, args []string) {
		listJobs()
	},
	Example: fmt.Sprintf(listExample, config.CliConfig.CommandName),
}

func listJobs() {
	client := utils.NewClient()
	resp, err := client.Job.List(orgId, filters)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.RenderOutput(resp)
}
