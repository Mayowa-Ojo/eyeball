package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Mayowa-Ojo/eyeball/utils"
	"github.com/spf13/cobra"
)

var (
	rootDir    string
	excludeDir []string
	stats      = make([]*utils.Stat, 0)

	rootCmd = &cobra.Command{
		Use:   "eyeball",
		Short: "Eyeball is a file stats generator for project directories",
		Long: `A Simple and Concise file stats generator with tabular display built in Go
				Eyeball is inspired by tokei (a rust project). Full documentation can be
				found here https://github.com/Mayowa-Ojo/eyeball`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			stats, err := utils.WalkDirectories(stats, rootDir, excludeDir)
			if err != nil {
				log.Fatal(err)
			}

			table := utils.GenerateTable(stats)

			fmt.Println(table.String())
		},
	}
)

// Execute - serves as entry point to cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&rootDir, "root", "r", "", "Project root directory relative to current path (default is '.')")
	rootCmd.PersistentFlags().StringSliceVarP(&excludeDir, "exclude", "e", []string{}, "Directories to exlude from stats")
}
