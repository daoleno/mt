package cmd

import (
	"github.com/daoleno/mt/file"
	"github.com/spf13/cobra"
)

var flagEditor string

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a thought",
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]
		err := file.OpenFile(filename, flagEditor)
		if err != nil {
			return err
		}
		return nil
	},
	ValidArgs: bashCompleteFile(),
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	openCmd.Flags().StringVar(&flagEditor, "editor", "vim", "Open with editor.")

}
