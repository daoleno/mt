package cmd

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/daoleno/mt/file"
	"github.com/daoleno/mt/utils"

	"github.com/spf13/cobra"
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render all markdown to beautiful html",
	RunE: func(cmd *cobra.Command, args []string) error {
		files, err := file.ListDataFile()
		if err != nil {
			return err
		}

		for _, fileinfo := range files {
			if fileinfo.Mode().IsRegular() {
				plaintext, err := ioutil.ReadFile(utils.DataDir() + "/" + fileinfo.Name())
				if err != nil {
					return err
				}
				htmltext := utils.Render(plaintext)
				utils.MkDir(utils.RenderDir())
				err = ioutil.WriteFile(utils.RenderDir()+"/"+strings.TrimSuffix(fileinfo.Name(), path.Ext(fileinfo.Name()))+".html", htmltext, 0644)
				if err != nil {
					return err
				}
			}
		}
		fmt.Printf("All your thoughts is generated in %s\n", utils.RenderDir())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
