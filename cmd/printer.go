package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"os"
)

type DirPrinter interface {
	linesPrinted() int
}

type MyPrinter struct {
	lines int
}

func (p *MyPrinter) linesPrinted() int {
	return p.lines
}

type PrintDirParam struct {
	maxDepth int
	maxLines int
	showAll  bool
}

func isHiddenFile(name string) bool {
	return strings.HasPrefix(name, ".")
}

var prefix = "- "

func (p *MyPrinter) printDir(path string, depth int, param PrintDirParam) error {
	if depth == 0 {
		// print root
		fileInfo, _ := os.Stat(path)
		if fileInfo != nil {
			fmt.Printf("%s%s%s%s\n", "", prefix, fileInfo.Name(), "/")
		}
		depth++
	}
	if depth > param.maxDepth {
		return nil
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if isHiddenFile(entry.Name()) && !param.showAll {
			continue
		}

		p.lines++
		if p.linesPrinted() > param.maxLines {
			return nil
		}

		indent := strings.Repeat("  ", depth)
		if entry.IsDir() {
			fmt.Printf("%s%s%s%s\n", indent, prefix, entry.Name(), "/")
			subPath := filepath.Join(path, entry.Name())
			err := p.printDir(subPath, depth+1, param)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Printf("%s%s%s\n", indent, prefix, entry.Name())
		}
	}
	return nil
}
