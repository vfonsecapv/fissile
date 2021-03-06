package cmd

import (
	"github.com/spf13/cobra"
)

// showReleaseCmd represents the release command
var showReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Displays information about BOSH releases.",
	Long: `
Displays a report of all jobs and packages in all referenced releases.
The report contains the name, version, description and counts of jobs and packages.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Show job information

		if err := fissile.LoadReleases(
			flagRelease,
			flagReleaseName,
			flagReleaseVersion,
			flagCacheDir,
		); err != nil {
			return err
		}

		if err := fissile.ListJobs(); err != nil {
			return err
		}

		return fissile.ListPackages()
	},
}

func init() {
	showCmd.AddCommand(showReleaseCmd)
}
