package cmd

import (
	"fmt"

	"github.com/benjic/marks/lib"
	"github.com/spf13/cobra"
)

const (
	addUse              = `add url`
	addShortDescription = `inserts a given URL into the bookmark collection`
)

func init() {
	RootCommand.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   addUse,
	Short: addShortDescription,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("A URL must be provided")
		}

		e, err := lib.NewEntry(args[0])

		if err != nil {
			return err
		}

		file, err := getBookmarksFile()
		if err != nil {
			fmt.Println("cant get file")
			return err
		}
		defer file.Close()

		provider, err := lib.NewProvider(file)

		if err := provider.Add(e); err != nil {
			return err
		}

		return nil
	},
}
