/*
Copyright © 2019 Tomokazu HIRAI <tomokazu.hirai@gmail.com>

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
	"fmt"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var namespace string

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
		listDeployments()
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
		listPods()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listDeploymentCmd)
	listCmd.AddCommand(listPodCmd)
	listPodCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace name")
	listDeploymentCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace name")
}

func listDeployments() {
	config := loadConfig()

	clientset, err := kubernetes.NewForConfig(config)

	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("🍺 There are %d deployments in the cluster\n", len(list.Items))
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
}

func listPods() {
	config := loadConfig()

	clientset, err := kubernetes.NewForConfig(config)

	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("🍉 There are %d pods in the cluster\n", len(pods.Items))
	for _, d := range pods.Items {
		fmt.Printf(" * %s\n", d.Name)
	}
}
