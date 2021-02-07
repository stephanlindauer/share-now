package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	evaluator "github.com/stephanlindauer/share-now/evaluator"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	SyncInterval = 1 * time.Hour
)

var (
	kubeConfigPath string
	client         *kubernetes.Clientset
	ctx            context.Context
)

type PodEvaluationResult struct {
	Pod            string                     `json:"pod"`
	RuleEvaluation []evaluator.RuleEvaluation `json:"rule_evaluation"`
}

func init() {
	log.SetOutput(os.Stdout)

	flag.StringVar(&kubeConfigPath, "kubeconfig", "", "(optional) path to the kubeconfig file. Will fall back to InCluster config.")
	flag.Parse()

	ctx = context.Background()

	config, err := clientcmd.BuildConfigFromFlags("", *&kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}

	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	for {
		pods, err := client.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		for _, pod := range pods.Items {
			currentPodEvaluationResult := PodEvaluationResult{
				Pod: pod.Name,
			}

			currentPodEvaluationResult.RuleEvaluation = append(currentPodEvaluationResult.RuleEvaluation,
				evaluator.EvaluateImagePrefix(pod),
				evaluator.EvaluateTeamLabelPresent(pod),
				evaluator.EvaluateRecentStartTime(pod),
			)

			jsonOutput, err := json.Marshal(currentPodEvaluationResult)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(string(jsonOutput))
		}

		time.Sleep(SyncInterval)
	}
}
