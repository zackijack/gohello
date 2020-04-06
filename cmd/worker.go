package cmd

import (
	"fmt"
	"time"

	"github.com/zackijack/gohello/config"

	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Running worker",
	Long:  `Running worker`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Config()
		fmt.Println("Worker starting")
		for {
			fmt.Println(cfg.GetString("worker.name"))
			time.Sleep(time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
