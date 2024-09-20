/*
Copyright © 2024 beyondlex <beyondsnk@163.com>
*/
package cmd

import (
	"errors"
	"os"

	"github.com/beyondlex/ftree/util"

	"github.com/spf13/cobra"
)

var Lines int
var Depth int

const defaultLines = 50
const defaultDepth = 8

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ftree",
	Short: "Print specific directory as a tree",
	Long: `Examples:
	ftree
	ftree ..
	ftree ~/Downloads
	ftree ~/Downloads -l=10 -d=1
	ftree .. -l 10 -d 1
	ftree -l 10 -d 1
`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Validate args
		if len(args) > 1 {
			return errors.New("too many arguments")
		}
		if len(args) == 0 {
			return nil // No path provided; assume the current directory
		}

		path := args[0]
		existed, isDir, err := util.IsDirOrFile(path)
		if err != nil {
			return errors.New("error checking path: " + err.Error())
		}
		if !existed {
			return errors.New("invalid path: " + path)
		}
		if !isDir {
			return errors.New("expected a directory, but a file was given: " + path)
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		printer := MyPrinter{}
		param := PrintDirParam{maxLines: Lines, maxDepth: Depth}
		// defaults to current dir
		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		printer.printDir(path, 0, param)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&Lines, "lines", "l", defaultLines, "Limit lines for printing")
	rootCmd.PersistentFlags().IntVarP(&Depth, "depth", "d", defaultDepth, "Limit depth for printing")
}
