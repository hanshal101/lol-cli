package cmd

import (
	"fmt"

	"github.com/hanshal101/client-go-cli/pkg/get"
	"github.com/spf13/cobra"
)

var getCMD = &cobra.Command{
	Use:   "get <resource>",
	Short: "Get Command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a resource")
			return
		}
		resource := args[0]
		namespace, _ := cmd.Flags().GetString("ns")
		if resource == "pods" || resource == "po" {
			get.GetAllPods(namespace)
		} else if resource == "deployments" {
			get.GetDeployments(namespace)
		} else if resource == "services" || resource == "svc" {
			get.GetServices(namespace)
		} else if resource == "namespace" || resource == "ns" {
			get.GetNamespaces()
		} else {
			fmt.Println("Give the correct resource name\nUse lol help or lol get -h for more info")
		}
	},
}

func init() {
	getCMD.Flags().String("ns", "", "Get Namespace")
	rootCMD.AddCommand(getCMD)
}
