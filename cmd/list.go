/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list k8s resources",
	Long: `list k8s resources
For example:

minikubectl list deployment.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Required nested subcommand.")
	},
}

// listDeploymentCmd represents the list command
var listDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "list deployments",
	Long: `list your deployments on k8s cluster.
For example:

minikubectl list deployment.`,
	Run: func(cmd *cobra.Command, args []string) {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err)
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}
		deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
		list, err := deploymentsClient.List(metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("🍺 There are %d deployments in the cluster\n", len(list.Items))
		for _, d := range list.Items {
			fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
		}

	},
}

// listPodCmd represents the list command
var listPodCmd = &cobra.Command{
	Use:   "pod",
	Short: "list pods",
	Long: `list your pods on k8s cluster.
For example:

minikubectl list pod.`,
	Run: func(cmd *cobra.Command, args []string) {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err)
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}

		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("🍉 There are %d pods in the cluster\n", len(pods.Items))
		for _, d := range pods.Items {
			fmt.Printf(" * %s\n", d.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listDeploymentCmd)
	listCmd.AddCommand(listPodCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listDeployments(clientset *kubernetes.Clientset) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("🍺 There are %d deployments in the cluster\n", len(list.Items))
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
}

func listPods(clientset *kubernetes.Clientset) {
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("🍉 There are %d pods in the cluster\n", len(pods.Items))
	for _, d := range pods.Items {
		fmt.Printf(" * %s\n", d.Name)
	}
}
