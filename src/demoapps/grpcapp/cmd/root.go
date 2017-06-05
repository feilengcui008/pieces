package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "grpcapp",
	Short: "a demo project for grpc-go",
	Long:  "a demo project for grpc-go",
	//Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
