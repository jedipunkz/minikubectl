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
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Options struct {
	deployment string
	app string
	container string
	image string
	port int32
	replica int32
}

var (
	o = &Options{}
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a deployment",
	Long: `create a deployment.
For example:

minikubectl create --deployment deployment01 --app app01 --container web01 --image nginx:latest --port 80`,
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

        config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
        if err != nil {
            panic(err)
        }

        clientset, err := kubernetes.NewForConfig(config)
        if err != nil {
            panic(err)
        }

		deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

		deployment := &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name: o.deployment,
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: int32Ptr(o.replica),
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"app": o.app,
					},
				},
				Template: apiv1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app": o.app,
						},
					},
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							{
								Name:  o.container,
								Image: o.image,
								Ports: []apiv1.ContainerPort{
									{
										Name:          "http",
										Protocol:      apiv1.ProtocolTCP,
										ContainerPort: o.port,
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
		result, err := deploymentsClient.Create(deployment)
		if err != nil {
			fmt.Printf("Fatal error: %s", err)
		}
		fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&o.deployment, "deployment", "d", "dep01", "deployment name")
	createCmd.MarkFlagRequired("deployment")
	createCmd.Flags().StringVarP(&o.app, "app", "a", "app01", "app name")
	createCmd.Flags().StringVarP(&o.container, "container", "c", "container01", "container name")
	createCmd.Flags().StringVarP(&o.image, "image", "i", "nginx:latest", "image name")
	createCmd.MarkFlagRequired("image")
	createCmd.Flags().Int32VarP(&o.port, "port", "p", 0, "port name")
	createCmd.MarkFlagRequired("port")
	createCmd.Flags().Int32VarP(&o.replica, "replica", "r", 1, "replicas number")
}

func int32Ptr(i int32) *int32 { return &i }

