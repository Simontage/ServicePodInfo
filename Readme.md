# ServicePodInfo

## 仓库名称: ServicePodInfo

## 描述:
ServicePodInfo 是一个简单的 Go 应用程序，旨在用于 Kubernetes（K8s）环境的实验和学习。它提供有关正在运行的服务的基本信息，包括启动时间（UTC 和 CST），运行时长，服务名称，版本以及运行它的 Pod 的名称。该应用程序允许用户通过环境变量自定义服务的详细信息，并通过 JSON 格式的 HTTP 响应提供信息。您可以将此仓库作为起点，探索和了解应用程序与 Kubernetes 的交互方式，以及环境变量如何影响服务行为。

## 特性:
- 获取并以 JSON 格式显示服务详细信息。
- 通过环境变量自定义服务名称、版本、时区和端口。
- 实验 Kubernetes 部署和 Pod 行为。


## 构建镜像并运行容器
```shell
docker build -t service-pod-info .
```
```shell
docker run -d -p 32262:32262 --name service-pod-info service-pod-info
```

随意克隆并修改此仓库，以便在 Kubernetes 环境中进行自己的实验和学习。

## Repository Name：ServicePodInfo

## Description：
ServicePodInfo is a simple Go application designed for experimentation and learning in Kubernetes (K8s) environments. It provides basic information about a running service, including its start time (in UTC and CST), running duration, service name, version, and the pod's name where it's running. The application allows users to set environment variables to customize service details, and it serves the information via an HTTP server with a JSON response format. Use this repository as a starting point to explore and understand how applications interact with Kubernetes and how environment variables can impact service behavior.

## Features:
- Fetch and display service details in JSON format.
- Customize service name, version, timezone, and port through environment variables.
- Experiment with Kubernetes deployment and pod behavior.

## Build Image and Run Container
```shell
docker build -t service-pod-info .
```
```shell
docker run -d -p 32262:32262 --name service-pod-info service-pod-info
```

Feel free to clone and modify this repository for your own experimentation and learning in Kubernetes environments.