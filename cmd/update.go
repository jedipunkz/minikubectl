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
	"flag"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
)

type OptionsUpdate struct {
	deployment string
	app        string
	container  string
	image      string
	port       int32
	replica    int32
}

var (
	ou = &OptionsUpdate{}
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a deployment",
	Long: `update a deployment.
For example:

minikubectl update --deployment deployment01 --app app01 --container web01 --image nginx:latest --port 80`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			fmt.Printf("%d", len(args))
			return errors.New("have to no argument.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		namespace := "default"

		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err)
		}

		client, err := dynamic.NewForConfig(config)
		if err != nil {
			panic(err)
		}

		deploymentsRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

		fmt.Printf("%d", ou.replica)
		// Update Deployment
		fmt.Println("Updating deployment...")
		retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			result, getErr := client.Resource(deploymentsRes).Namespace(namespace).Get(ou.deployment, metav1.GetOptions{})
			if getErr != nil {
				panic(getErr)
			}

			// update replica number
			if ou.replica != -99 {
				if err := unstructured.SetNestedField(result.Object, int64(ou.replica), "spec", "replicas"); err != nil {
					panic(err)
				}
			}

			// update image
			if ou.image != "" {
				containers, found, err := unstructured.NestedSlice(result.Object, "spec", "template", "spec", "containers")
				if err != nil || !found || containers == nil {
					panic(err)
				}
				if err := unstructured.SetNestedField(containers[0].(map[string]interface{}), ou.image, "image"); err != nil {
					panic(err)
				}
				if err := unstructured.SetNestedField(result.Object, containers, "spec", "template", "spec", "containers"); err != nil {
					panic(err)
				}
			}

			_, updateErr := client.Resource(deploymentsRes).Namespace(namespace).Update(result, metav1.UpdateOptions{})
			return updateErr
		})
		if retryErr != nil {
			panic(retryErr)
		}
		fmt.Println("🐙 Updated deployment...")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&ou.deployment, "deployment", "d", "dep01", "deployment name")
	updateCmd.MarkFlagRequired("deployment")
	updateCmd.Flags().StringVarP(&ou.image, "image", "i", "", "image name")
	updateCmd.MarkFlagRequired("image")
	updateCmd.Flags().Int32VarP(&ou.replica, "replica", "r", -99, "replicas number")
}