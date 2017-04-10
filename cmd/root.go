package cmd

import (
	"github.com/spf13/cobra"
)

const (
	rootUse = `marks`
)

// RootCommand is the entry command for the marks program.
var RootCommand = cobra.Command{
	Use: rootUse,
}
