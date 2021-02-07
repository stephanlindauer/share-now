# Code Challenge Kubernetes Pod Evaluation Service

The main purpose of this code challenge is to understand how problems are approached by our potential candidates. This will also be a conversation starter for a potential follow-up interview session with the platform team.

It is completely fine if the problem is not solved entirely or the code quality is not perfect. We will take a closer look and dive into your code together with you in the next round.

Please send us the result within 7 days, we whill then look into it and contact you again.

## Problem Statement

At Share-Now we are running thousands of containers in 50+ Kubernetes clusters.
Considering this large scale, it can be difficult to do basic house-keeping on the pods in a cluster.
Therefore, we want to write a microservice, that evaluates the current status of a cluster, by watching all running pods in a cluster and evaluating some rules.

The following rules should be evaluated:

```yaml
rules:
- name: image_prefix
  description: "ensure the pod only uses images prefixed with `bitnami/`"
  output: boolean
- name: team_label_present
  description: "ensure the pod contains a label `team` with some value"
  output: boolean
- name: recent_start_time
  description: "ensure the pod has not been running for more than 7 days according to it's `startTime`"
  output: boolean
```

The service should evaluate these rules for all pods in the cluster, as well as output the results on stdout in json log format one line per pod.
The results on stdout should look similar to the following example:

```json
{"pod": "mytest", "rule_evaluation": [{"name": "image_prefix", "valid": true}, {"name": "team_label_present", "valid": true}, {"name": "recent_start_time", "valid": false}]}
{"pod": "another", "rule_evaluation": [{"name": "image_prefix", "valid": false}, {"name": "team_label_present", "valid": true}, {"name": "recent_start_time", "valid": false}]}
```

## Technical Requirements

- you are free to use any programming language of your choice, e.g. go, python, ruby, ...
- please commit your work to a public git repository and send us the link for evaluation (e.g. on GitHub, GitLab or similar)
- please provide a `README.md` file in your repository with some information on how to run your code

## Useful Links

- [Kubernetes Intro](https://kubernetes.io/docs/tutorials/kubernetes-basics/)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/) a tool to easily create a small Kubernetes cluster on your local machine using [docker](https://www.docker.com/get-started)
- [Kubernetes Client Libraries](https://kubernetes.io/docs/reference/using-api/client-libraries/) to interact with the Kubernetes API from your favorite programming language

