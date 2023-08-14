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
	ctx := context.Background()
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

	// pods list
	pods, err := clientSet.CoreV1().Pods("kube-system").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error listing pods: %v", err)
	}
	fmt.Println("printing pods list .....")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	// deployments list
	deployments, err := clientSet.AppsV1().Deployments("kube-system").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error listing deployments: %v", err)
	}
	fmt.Println("printing deployment list .....")
	for _, d := range deployments.Items {
		fmt.Println(d.Name)
	}
}
