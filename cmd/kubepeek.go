/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// kubepeekCmd represents the kubepeek command
var kubepeekCmd = &cobra.Command{
	Use:   "kubepeek",
	Short: "kubepeek provides a CLI-based, easy-to-use overview of all resources in Kubernetes. It offers various views, allowing users to select and understand Kubernetes resources according to their preference.",
	Long:  `"kubepeek" is a tool that enables users to have a comprehensive and simplified view of all resources in a Kubernetes cluster through the command-line interface (CLI). It offers a range of different views, each presenting the Kubernetes resources in a distinct format. Users can choose the view that best suits their needs, allowing them to easily understand and analyze the resources present in the Kubernetes cluster. With "kubepeek," users gain a convenient and efficient way to gain insights into the various aspects of their Kubernetes environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		kubeconfigPath := "./.kube/config"
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			panic(err.Error())
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("kubepeek called")
		pods, err := clientset.CoreV1().Pods("namespace").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		for _, pod := range pods.Items {
			fmt.Println(pod.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(kubepeekCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kubepeekCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kubepeekCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
