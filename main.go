package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error %s building config from flags\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error Creating client: %s", err.Error())
	}
	pods, err := clientset.CoreV1().Pods("cafe").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	for _, pod := range pods.Items {
		fmt.Printf("pod name %s\n", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("cafe").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	for _, deps := range deployments.Items {
		fmt.Printf("Deployments %s \n", deps.Name)
	}
}
