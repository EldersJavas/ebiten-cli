package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init your ebiten game",
	Long: `Init your ebiten game`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().String("name","game-"+ strconv.FormatInt(time.Now().Unix(),10),"Game project name")
	initCmd.PersistentFlags().String("type","s","Project type")
	initCmd.PersistentFlags().String("module","","Go module name")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
