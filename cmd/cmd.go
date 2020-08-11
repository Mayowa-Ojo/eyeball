package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

// Stat - defines the stats for each file
type stat struct {
	name          string
	size          int64
	NumberOfLines int
}

var (
	// Stats - total stats for given directory
	stats     []stat
	ignoreDir = []string{
		".git",
		".idea",
		".vscode",
	}
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
						love by spf13 and friends in Go.
						Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute - serves as entry point to cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getNumberOfLines(filename string) int {
	content, err := ioutil.ReadFile(filename)
	NumLines := 0

	if err != nil {
		log.Fatal(err)
	}

	for _, byte := range content {
		if string(byte) == "\n" {
			NumLines++
		}
	}

	return NumLines
}

func generateTable(stats []stat) *simpletable.Table {
	table := simpletable.New()
	totalFileSize := int64(0)
	totalNumberOfFiles := 0
	totalNumberOfLines := 0

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "FILE NAME"},
			{Align: simpletable.AlignCenter, Text: "FILE SIZE (kb)"},
			{Align: simpletable.AlignCenter, Text: "NO OF LINES"},
		},
	}

	for i, stat := range stats {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", i)},
			{Align: simpletable.AlignRight, Text: stat.name},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", stat.size)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", stat.NumberOfLines)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
		totalFileSize += stat.size
		totalNumberOfFiles++
		totalNumberOfLines += stat.NumberOfLines
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", totalNumberOfFiles)},
			{},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", totalFileSize)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", totalNumberOfLines)},
		},
	}

	table.SetStyle(simpletable.StyleMarkdown)

	return table
}

func walkDirectories(stats []stat, root string) error {
	err := filepath.Walk(root, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() && fileInfo.Name() == ".git" { //TODO: handle ignored directories
			return filepath.SkipDir
		}

		if !fileInfo.IsDir() {
			n := getNumberOfLines(path)
			stats = append(stats, stat{fileInfo.Name(), fileInfo.Size(), n})
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
