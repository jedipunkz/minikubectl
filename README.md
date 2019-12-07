# Minikubectl

🔥 Minimal 🔥 k8s Client CLI. Using [client-go](https://github.com/kubernetes/client-go) (as Kubernetes Library).

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
