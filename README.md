# share-now

## How to run the application

```bash
# use kind to spawn a new cluster
kind create cluster

# apply k8s manifest to local kind cluster
kubectl apply -f k8s/

# admire beautiful json logs
kubectl logs -l app=share-now-evaluator -f
```

For a faster local development cycle, skip the Pod deployment and run application on the host machine:
```bash
go run main.go --kubeconfig ~/.kube/config
```

The Docker image is automatically built and pushed to `https://quay.io/repository/stephanlindauer/share-now`.

**Update**: I just found out, that `kind load docker-image my-custom-image` exists. :D

## How it went

I kept my code relatively simple and strict to just what the code challenge requested. Going further with this project, I would think about the following improvements:
- Tests
- Prometheus exporter (tracking things like uptime or ratio of uncompliant pods)
- For local development and easier deployment into Kubernetes cluster (without the detour through a third-party repository), I would the build image directly in the cluster (in Kaniko for example).
- Use `client.CoreV1().Pods("").Watch` instead of continuously polling for all Pods.
- The Problem Statement suggests that there could be a `yaml` file with configurations for the different evaluation rules. I think it would be cool, to make the application more configurable via this potential config file. I think of something more dynamic like this:
```
- name: image_prefix
  description: "ensure the pod only uses images prefixed with `bitnami/`"
  output: boolean
  path: ".spec.containers[].image"
  type: "RegEx"
  value: "^bitnami/.*"
```
- The config could also be created as a CRD, but with the current complexity of this tool it seems absurd. :D
- Properly configure the Pod this application is running in with health/liveness probes, priorities, resource limits and resource requests and so on.
- Think about if this could also be a CronJob resource. After I decided that it only runs every hour, it would be nice to free up those resources in the mean time.
- Refine what the evaluator can do. Does it need access to all namespaces?
- Should it run in the default namespace?

