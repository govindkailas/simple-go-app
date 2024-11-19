
# Pod Info Service

A simple Go microservice that displays pod information when deployed to Kubernetes.

## Overview

This microservice prints:
- Pod Name
- Pod Namespace  
- Node Name

It helps demonstrate basic Kubernetes concepts like deployments, pods and services.

## Components

- `simple-app.go` - Go service that exposes pod info on HTTP endpoint
- `Dockerfile` - For building the container image
- `simple-app-deploy-svc.yaml` - Kubernetes deployment manifest, also include service manifest to expose the pods as `LoadBalancer`

## Building

Build the Docker image:

``` bash
docker build -t pod-info-service:latest .
```

## Deployment

Deploy to Kubernetes:

``` bash
kubectl apply -f simple-app-deploy-svc.yaml
```


## Usage

Once deployed, you can access the service via:

``` bash
kubectl get svc simple-go-app-svc
NAME                TYPE           CLUSTER-IP      EXTERNAL-IP    PORT(S)        AGE
simple-go-app-svc   LoadBalancer   10.95.168.183   10.125.0.146   80:30985/TCP   29m
```

# If using NodePort
`curl http://<node-ip>:<node-port>`

# If using LoadBalancer
`curl http://<load-balancer-ip>`



The response will show the pod information in JSON format:

``` json
{
  "podName": "pod-info-6f7d9b5f45-abc12",
  "podNamespace": "default", 
  "nodeName": "worker-1"
}
```



## Local Development

Run locally:


go run main.go


## License

MIT License
