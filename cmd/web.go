/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/daoleno/mt/utils"
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Server static html page",
	Run: func(cmd *cobra.Command, args []string) {
		serveStatic()
	},
}

func init() {
	rootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ServeStatic - server static html file
func serveStatic() {
	http.Handle("/", http.FileServer(http.Dir(utils.RenderDir())))
	fmt.Println("All your thoughs is here: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
