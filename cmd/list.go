// Package cmd /*
package cmd

import (
	"github.com/EldersJavas/ebiten-cli/tmpl"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `list`,
	Run: func(cmd *cobra.Command, args []string) {
		t := tmpl.Gobal
		for a, b := range t.AllTmpl {
			cmd.Printf("%v:%v", a, b.Name)

		}
	},
}
var IsTmpl = true

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&IsTmpl, "tmpl", "t", true, "Show Tmpl")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
