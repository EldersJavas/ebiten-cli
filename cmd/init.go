package cmd

import (
	"fmt"
	"github.com/EldersJavas/ebiten-cli/cmd/tool"
	"github.com/EldersJavas/ebiten-cli/model"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init your ebiten game",
	Long:  `Init your ebiten game`,
	Run: func(cmd *cobra.Command, args []string) {
		rootpath, _ := os.Getwd()
		fmt.Println(args)
		cmd.Println("Target folder: " + tool.WinDir(rootpath))
		if tool.VaildFile("config.json") {
			cmd.Println("Target flies: " + "./config.json")
			if InitJon.Game.Repo == "" {
				InitJon.Game.Repo = InitJon.Game.Name
			}

		} else {
			cmd.Println("./config.json already exists. Please delete it and again.")
			return
		}

	},
}

var InitJon model.ECSTProject
var rootpath string

// init
func init() {
	initCmd.PersistentFlags().StringVar(&InitJon.Game.Name, "name", "game-"+strconv.FormatInt(time.Now().Unix(), 10), "Game project name")
	initCmd.PersistentFlags().StringVar(&InitJon.Game.Type, "type", "classic", "Project type")
	initCmd.PersistentFlags().StringVar(&InitJon.Game.Repo, "module", "", "Go module name")
	rootCmd.AddCommand(initCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
