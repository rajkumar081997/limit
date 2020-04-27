package cmd

import (
	"fmt"
	"os"

	"github.com/m/v2/client/cmd/core"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	add := core.Additem()
	RootCmd.AddCommand(add)
	show := core.Showitem()
	RootCmd.AddCommand(show)
	list := core.List()
	RootCmd.AddCommand(list)
	rem := core.Remove()
	RootCmd.AddCommand(rem)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
