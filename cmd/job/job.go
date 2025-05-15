package job

import (
	"github.com/spf13/cobra"
)

var description = `
This command consists of multiple subcommands to interact with jobs.
It can be used to create, update, delete and list jobs.
`

var Cmd = &cobra.Command{
	Use:   "job create|list [ARGS]",
	Short: "create and list jobs",
	Long:  description,
}

func init() {
	Cmd.AddCommand(createJobCmd)
	createJobCmd.Flags().StringVarP(&jobCreateCommand, "command", "c", "", "Command to execute: plan,apply,destroy (required)")
	_ = createJobCmd.MarkFlagRequired("command")
	createJobCmd.Flags().StringVarP(&jobCreateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = createJobCmd.MarkFlagRequired("organization-id")
	createJobCmd.Flags().StringVarP(&workspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = createJobCmd.MarkFlagRequired("workspace-id")

	Cmd.AddCommand(listJobsCmd)
	listJobsCmd.Flags().StringVarP(&filters, "filter", "f", "", "Filter")
	listJobsCmd.Flags().StringVarP(&orgId, "organization-id", "", "", "Organization Id (required)")
	_ = listJobsCmd.MarkFlagRequired("organization-id")
}
