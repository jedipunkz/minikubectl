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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var name string

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a deployment.",
	Long: `delete a deployment of k8s
Allowed Arguments: deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Required nested subcommand.")
	},
}

// deleteDeploymentCmd represents the delete command
var deleteDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "delete a deployment.",
	Long: `delete a deployment of k8s
For example:

minikubectl delete deployment --deployment demo`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			fmt.Printf("%d", len(args))
			return errors.New("No need to have argument.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		deleteDeployment()
	},
}

// deleteNsCmd represents the delete command
var deleteNsCmd = &cobra.Command{
	Use:   "ns",
	Short: "delete a namesapce.",
	Long: `delete a namespace of k8s
For example:

minikubectl delete ns--name test`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteNs()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteDeploymentCmd)
	deleteDeploymentCmd.Flags().StringVarP(&name, "name", "n", "", "deployment name")
	deleteDeploymentCmd.MarkFlagRequired("name")
	deleteCmd.AddCommand(deleteNsCmd)
	deleteNsCmd.Flags().StringVarP(&name, "name", "n", "", "namespace name")
	deleteNsCmd.MarkFlagRequired("name")
}

func deleteDeployment() {
	config := loadConfig()

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	fmt.Println("Deleting deployment...")
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		fmt.Printf("☔ Fatal error: %s", err)
	} else {
		fmt.Println("🍺 Deleted deployment.")
	}
}

func deleteNs() {
	namespace := name
	config := loadConfig()

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	errc := clientset.CoreV1().Namespaces().Delete(namespace, &metav1.DeleteOptions{})
	if errc != nil {
		fmt.Printf("☔ Fatal error: %s", errc)
	}
}
