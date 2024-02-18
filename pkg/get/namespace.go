package get

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNamespaces() {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Namespaces :\n")
	for _, namespace := range namespaces.Items {
		fmt.Printf(" - %s\n", namespace.Name)
	}
}
