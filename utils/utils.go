package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/alexeyco/simpletable"
)

// Stat - defines the stats for each file
type Stat struct {
	name          string
	size          int64
	NumberOfLines int
}

var (
	// Stats - total stats for given directory
	ignoreDir = []string{
		".git",
		".idea",
		".vscode",
	}
)

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

// GenerateTable - builds a table from given stats
func GenerateTable(stats []*Stat) *simpletable.Table {
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

// WalkDirectories - recurrsively traverses specified directory to get each file
func WalkDirectories(stats []*Stat, root string, excludeDir []string) ([]*Stat, error) {
	err := filepath.Walk(root, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() && fileInfo.Name() == ".git" { //TODO: handle ignored directories
			return filepath.SkipDir
		}

		if !fileInfo.IsDir() {
			n := getNumberOfLines(path)
			stats = append(stats, &Stat{fileInfo.Name(), fileInfo.Size(), n})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return stats, nil
}
