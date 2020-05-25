package cmd

import (
	"fmt"
	"sort"

	"github.com/daoleno/mt/file"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all thoughts",
	RunE: func(cmd *cobra.Command, args []string) error {
		files, err := file.ListDataFile()
		if err != nil {
			return err
		}
		// Sort file by file modtime
		sort.Slice(files, func(i, j int) bool {
			return files[i].ModTime().After(files[j].ModTime())
		})

		// Print file
		for _, f := range files {
			// TODO: Chinese and English is not aligned correctly.
			fmt.Printf("%-30s\t%s\n", f.Name(), f.ModTime().Format("2006-01-02 15:04:05"))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
