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

	// "k8s.io/client-go/1.5/pkg/api"
	// "k8s.io/client-go/1.5/pkg/api/unversioned"
	// "k8s.io/client-go/1.5/pkg/api/v1"
	"github.com/arekkas/kubernetes/pkg/api"
	"github.com/arekkas/kubernetes/pkg/api/v1"
	"github.com/arekkas/kubernetes/pkg/client/unversioned"
	"k8s.io/client-go/kubernetes"
)

type Options struct {
	deployment string
	app        string
	container  string
	image      string
	port       int32
	replica    int32
}

var (
	o = &Options{}
)

type OptionsNs struct {
	name string
}

var (
	on = &OptionsNs{}
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create root command",
	Long: `create command
Allowed Arguments: deploymet`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Required nested subcommand.")
	},
}

// createDeploymentCmd represents the create command
var createDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "create a deployment",
	Long: `create a deployment.
For example:

minikubectl create deployment --deployment deployment01 --app app01 --container web01 --image nginx:latest --port 80`,
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

		deployment := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "apps/v1",
				"kind":       "Deployment",
				"metadata": map[string]interface{}{
					"name": o.deployment,
				},
				"spec": map[string]interface{}{
					"replicas": o.replica,
					"selector": map[string]interface{}{
						"matchLabels": map[string]interface{}{
							"app": o.app,
						},
					},
					"template": map[string]interface{}{
						"metadata": map[string]interface{}{
							"labels": map[string]interface{}{
								"app": o.app,
							},
						},

						"spec": map[string]interface{}{
							"containers": []map[string]interface{}{
								{
									"name":  o.container,
									"image": o.image,
									"ports": []map[string]interface{}{
										{
											"name":          "http",
											"protocol":      "TCP",
											"containerPort": o.port,
										},
									},
								},
							},
						},
					},
				},
			},
		}

		// Create Deployment
		fmt.Println("Creating deployment...")
		result, err := client.Resource(deploymentsRes).Namespace(namespace).Create(deployment, metav1.CreateOptions{})
		if err != nil {
			fmt.Printf("☔ Fatal error: %s", err)
		} else {
			fmt.Printf("🍺 Created deployment %q.\n", result.GetName())
		}
	},
}

// createDeploymentCmd represents the create command
var createNamespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "create a namespace",
	Long: `create a namespace.
For example:

minikubectl create namespace --name demo`,
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

		ns := new(v1.Namespace)
		// ns.TypeMeta = unversioned.TypeMeta{Kind: "Namespace", APIVersion: "v1"}
		ns.ObjectMeta = v1.ObjectMeta{Name: on.name}
		ns.Spec = v1.NamespaceSpec{}
		// nsname, err := clientset.Core().Namespace().Create(ns)
		nsname, err := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}
		if err != nil {
			panic(err)
		}

		fmt.Printf("namespace: %s have created\n", nsname.ObjectMeta.Name)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createDeploymentCmd)
	createDeploymentCmd.Flags().StringVarP(&o.deployment, "deployment", "d", "dep01", "deployment name")
	createDeploymentCmd.MarkFlagRequired("deployment")
	createDeploymentCmd.Flags().StringVarP(&o.app, "app", "a", "app01", "app name")
	createDeploymentCmd.Flags().StringVarP(&o.container, "container", "c", "container01", "container name")
	createDeploymentCmd.Flags().StringVarP(&o.image, "image", "i", "nginx:latest", "image name")
	createDeploymentCmd.MarkFlagRequired("image")
	createDeploymentCmd.Flags().Int32VarP(&o.port, "port", "p", 0, "port name")
	createDeploymentCmd.MarkFlagRequired("port")
	createDeploymentCmd.Flags().Int32VarP(&o.replica, "replica", "r", 1, "replicas number")
	createCmd.AddCommand(createNamespaceCmd)
	createNamespaceCmd.Flags().StringVarP(&on.name, "name", "n", "", "namespace name")
	createNamespaceCmd.MarkFlagRequired("name")
}
