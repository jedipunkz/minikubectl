# Minikubectl

[![Build Status](https://travis-ci.org/jedipunkz/minikubectl.svg?branch=master)](https://travis-ci.org/jedipunkz/minikubectl)

## Description

🔥 Minimal 🔥 k8s Client CLI. Using [client-go](https://github.com/kubernetes/client-go) (as Kubernetes Library).

## Installation

### Go get

```bash
go get -u github.com/jedipunkz/minikubectl
```

### Build

```bash
git clone https://github.com/jedipunkz/minikubectl .
go build
```

## Pre-Requirements

* local or remote kubernetes environment

You need any k8s cluster on local (such as minikube) or remote environment.

* local kubectl command and $HOME/.kube/config file

You need $HOME/.kube/config which you use for connecting your k8s cluster.

## Usage

### List Deployments.

Options

| Option | Description | Default Value |
|--------|-------------|---------------|
| --namespace | namespace name | default |

Example

```bash
minikubectl list deployments [--namespace default]
🍺 There are 1 deployments in the cluster
 * nginx-deployment (2 replicas)
```

### List Pods.

Options

| Option | Description | Default Value |
|--------|-------------|---------------|
| --namespace | namespace name | default |

Example

```bash
minikubectl list pods [--namespace default]
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

### Create Deployment.

Options

| Option | Description | Default Value | Required |
|--------|-------------|---------------|----------|
| --name | Deployment Name | N/A | ✔ |
| --app | Application Name | app01 |  |
| --container | Container Name | container01 | |
| --image | Container Image Name and Tag | nginx:latest | ✔ |
| --port | Port Number | 0 | ✔ |
| --replica | Replica Number | 1 | |

Example

```bash
minikubectl create deployment --name demo --app demo --container demo --image nginx:1.12 --port 80 --replica 1
Creating deployment...
🍺 Created deployment "demo".
```

### Update Deployment.

Options

| Option | Description | Default Value | Required |
|--------|-------------|---------------|----------|
| --name | Deployment Name | dep01 | |
| --image| Container Image Name and Tag | N/A |
| --replica| Replica Number | N/A | |

Example

```bash
# update image tag
minikubectl update deployment --name demo --image nginx:1.11
Updating deployment...
🐙 Updated deployment...
# update replica number
minikubectl update --deployment demo --replica 10
Updating deployment...
🐙 Updated deployment...
```

### Delete Deployment.

Options

| Option | Description | Default Value | Required |
|--------|-------------|---------------|----------|
| --name | Deployment Name | N/A | ✔ |

Example

```bash
minikubectl delete deployment --name demo
Deleting deployment...
🍺 Deleted deployment.
```

## Author

Tomokazu HIRAI https://twitter.com/jedipunkz
