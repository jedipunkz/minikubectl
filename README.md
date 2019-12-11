# Minikubectl

🔥 Minimal 🔥 k8s Client CLI. Using [client-go](https://github.com/kubernetes/client-go) (as Kubernetes Library).

[![Build Status](https://travis-ci.org/jedipunkz/minikubectl.svg?branch=master)](https://travis-ci.org/jedipunkz/minikubectl)

## Installation

```bash
go get -u github.com/jedipunkz/minikubectl
```

## Pre-Requirements

* local or remote kubernetes environment
* local kubectl command and .kube/config file

## Usage

List Deployments.

```bash
minikubectl list deployments
🍺 There are 1 deployments in the cluster
 * nginx-deployment (2 replicas)
```

List Pods.

```bash
minikubectl list pods
🍉 There are 11 pods in the cluster
 * nginx-deployment-54f57cf6bf-65k86
 * nginx-deployment-54f57cf6bf-6lj2s
 * coredns-5644d7b6d9-kd5z4
 * coredns-5644d7b6d9-ndg4t
 * etcd-minikube
 * kube-addon-manager-minikube
 * kube-apiserver-minikube
 * kube-controller-manager-minikube
 * kube-proxy-xxgmt
 * kube-scheduler-minikube
 * storage-provisioner
```

Create Deployment.

* --deployment: Deployment Name (Required Argument)
* --app: Application Name
* --container: Container Name
* --image: Container Image Name and Tag (Required Argument)
* --port: Port Number (Required Argument)
* --replica: Replica Number

```bash
minikubectl create --deployment demo --app demo --container demo --image nginx:1.12 --port 80 --replica 1
Creating deployment...
🍺 Created deployment "demo".
```

Update Deployment

* --deployment: Deployment Name (Required Argument)
* --image: Container Image Name and Tag
* --replica: Replica Number

```bash
# update image tag
minikubectl update --deployment demo --image nginx:1.11
Updating deployment...
🐙 Updated deployment...
# update replica number
minikubectl update --deployment demo --replica 10
Updating deployment...
🐙 Updated deployment...

Delete Deployment.

* --deployment: Deployment Name (Required Argument)

```bash
minikubectl delete --deployment demo
Deleting deployment...
🍺 Deleted deployment.
```
