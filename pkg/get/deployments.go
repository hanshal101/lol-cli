package get

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDeployments(namespace string) {
	if namespace == "" {
		namespace = v1.NamespaceDefault
	}
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Deployments in %s:\n", namespace)
	for _, deployment := range deployments.Items {
		fmt.Printf(" - %s\n", deployment.Name)
	}
}
