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
```

List Pods.

```bash
minikubectl list pods
```
