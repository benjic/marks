package cmd

import (
	"fmt"

	"github.com/benjic/marks/lib"
	"github.com/spf13/cobra"
)

const (
	suggestUse              = `suggest {term}`
	suggestShortDescription = `returns a list of bookmarks for the given terms`
)

func init() {
	RootCommand.AddCommand(suggest)
}

var suggest = &cobra.Command{
	Use:   suggestUse,
	Short: suggestShortDescription,
	RunE: func(command *cobra.Command, args []string) error {
		file, err := getBookmarksFile()
		if err != nil {
			return err
		}
		defer file.Close()

		provider, err := lib.NewProvider(file)

		if err != nil {
			return err
		}

		for _, suggestion := range provider.Suggest(args) {
			fmt.Println(suggestion.URL())
		}

		return nil
	},
}
