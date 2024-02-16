package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "lol",
	Short: "This is a sample kubernetes communicator (like kubectl)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is kubectl")
	},
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println("Error in running lol")
	}
}
