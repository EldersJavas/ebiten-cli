package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/EldersJavas/ebiten-cli/cmd/tool"
	"github.com/EldersJavas/ebiten-cli/model"
	"io/ioutil"
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
		if tool.VaildFile(rootpath + tool.DirFormat() + "config.json") {
			cmd.Println("Target flies: " + "./config.json")
			if InitJon.Game.Repo == "" {
				InitJon.Game.Repo = InitJon.Game.Name
			}
			switch InitJon.Game.Type {
			case "":

			}
			SaveFile("", InitJon)

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
	initCmd.Flags().StringVarP(&InitJon.Game.Name, "name", "n", "game-"+strconv.FormatInt(time.Now().Unix(), 10), "Game project name")
	initCmd.Flags().StringVarP(&InitJon.Game.Type, "type", "t", "classic", "Project type")
	initCmd.Flags().StringVarP(&InitJon.Game.Repo, "module", "m", "", "Go module name(blank same game name)")
	rootCmd.AddCommand(initCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func SaveFile(filename string, thing interface{}) error {
	saveData, _ := json.Marshal(thing)
	err := ioutil.WriteFile(filename, saveData, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}
