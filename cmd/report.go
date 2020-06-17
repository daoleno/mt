/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/daoleno/mt/utils"
	"github.com/daoleno/mt/vcs"
	"github.com/daoleno/tgraph"
	"github.com/spf13/cobra"
)

var (
	start, end string
	raw        bool
)

const (
	// the length of `git diff --numstat`
	numStatLength = 3
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:     "report <start date> <end date>",
	Short:   "Writing report",
	Example: "mt report 3.weeks.ago 1.weeks.ago",
	RunE: func(cmd *cobra.Command, args []string) error {
		git := vcs.ByCmd("git")
		out, err := git.DiffStat(utils.DataDir(), start, end)
		if err != nil {
			return err
		}
		if raw {
			fmt.Printf("%s", out)
			return nil
		}

		// Parse git diff stat output
		labels, data, err := parseGitDiffStat(out)
		if err != nil {
			return err
		}
		colors := []string{"green", "red"}

		tgraph.Chart("Writing Report\n", labels, data, colors, 50, true, "")
		return nil
	},
}

func parseGitDiffStat(buf []byte) (labels []string, data [][]float64, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(buf))

	for scanner.Scan() {
		stat := strings.Split(scanner.Text(), "\t")
		// Skip invalid output
		if len(stat) != numStatLength || stat[0] == "-" || stat[1] == "-" {
			continue
		}

		// Convert git diff --numstat output to standerd data
		stat0, err := strconv.ParseFloat(stat[0], 64)
		if err != nil {
			return nil, nil, err
		}
		stat1, err := strconv.ParseFloat(stat[1], 64)
		if err != nil {
			return nil, nil, err
		}

		d := []float64{stat0, stat1}
		l := stat[2]

		data = append(data, d)
		labels = append(labels, l)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return labels, data, nil
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
	reportCmd.Flags().BoolVar(&raw, "raw", false, "Raw report")
}
