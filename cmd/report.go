/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/daoleno/mt/utils"
	"github.com/daoleno/mt/vcs"
	"github.com/spf13/cobra"
)

var (
	start, end string
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report 3.weeks.ago 1.weeks.ago",
	Short: "Weekly report",
	RunE: func(cmd *cobra.Command, args []string) error {
		git := vcs.ByCmd("git")
		out, err := git.DiffStat(utils.DataDir(), start, end)
		if err != nil {
			return err
		}
		fmt.Printf("%s", out)

		return nil

	},
}

func init() {
	rootCmd.AddCommand(reportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	reportCmd.Flags().StringVar(&start, "start", "1.weeks.ago", "Start date")
	reportCmd.Flags().StringVar(&end, "end", "", "End date")
}
