package main

import (
	"context"

	"github.com/spf13/cobra"
)

// rootCMD is the base command when called without any subcommands.
var rootCMD = &cobra.Command{
	Short:             "This cli tool is provided to interact with the whole application.",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

// ctx holds application's context.
var ctx context.Context

// init will be called when this package is imported.
func init() {
	ctx = context.Background()
}

// main is the application's kernel.
func main() {
	cobra.CheckErr(rootCMD.Execute())
}
