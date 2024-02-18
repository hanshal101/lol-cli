package get

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetServices(namespace string) {
	if namespace == "" {
		namespace = v1.NamespaceDefault
	}
	services, err := clientset.CoreV1().Services(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Services in %s:\n", namespace)
	for _, service := range services.Items {
		fmt.Printf(" - %s\n", service.Name)
	}
}
