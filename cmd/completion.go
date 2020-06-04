/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/daoleno/mt/file"
	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash/zsh completion scripts",
	Long: `Generated script will get you completions of subcommands and flags. 
Bash: Copy completion_bash.sh to /etc/bash_completion.d/ and reset your terminal to use autocompletion. 
Zsh: Copy the content of completion_zsh.sh to ~/.zshrc and source .zshrc to use autocompletion.
`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletionFile("completion_bash.sh")
		rootCmd.GenZshCompletionFile("completion_zsh.sh")
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func bashCompleteFile() []string {
	files, err := file.ListDataFile()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var fName []string
	for _, f := range files {
		fName = append(fName, f.Name())
	}
	return fName
}
