package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the root cmd of an app
var RootCmd = &cobra.Command{
	Use:   "grpcapp",
	Short: "a demo project for grpc-go",
	Long:  "a demo project for grpc-go",
	//Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute execute the root cmd
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
