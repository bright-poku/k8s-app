package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// specify kubeconfig files
	kubeConfig := flag.String("kubeConfig", "/Users/bap/.kube/config", "location to the file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		fmt.Printf("error creating kubeconfig: %v", err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error creating config: %v", err)
	}

	pods, err := clientSet.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error listing pods: %v", err)
	}
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
}
