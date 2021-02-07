package evaluator

import (
	"regexp"
	"time"

	"github.com/jonboulle/clockwork"
	v1 "k8s.io/api/core/v1"
)

var clock clockwork.Clock

type RuleEvaluation struct {
	Name  string `json:"name"`
	Valid bool   `json:"valid"`
}

func EvaluateRecentStartTime(pod v1.Pod) RuleEvaluation {
	var (
		maxAge  = 7 * 24 * time.Hour // 7 days
		created = pod.ObjectMeta.CreationTimestamp.Time
		age     = time.Since(created)
		valid   = true
	)

	if age > maxAge {
		valid = false
	}

	return RuleEvaluation{
		Name:  "recent_start_time",
		Valid: valid,
	}
}

func EvaluateTeamLabelPresent(pod v1.Pod) RuleEvaluation {
	var (
		keyToLookFor = "team"
		valid        = false
	)

	if _, ok := pod.Labels[keyToLookFor]; ok {
		if pod.Labels[keyToLookFor] != "" {
			valid = true
		}
	}

	return RuleEvaluation{
		Name:  "team_label_present",
		Valid: valid,
	}
}

func EvaluateImagePrefix(pod v1.Pod) RuleEvaluation {
	var (
		regExpToLookFor    = "^bitnami/.*"
		relevantContainers = append(pod.Spec.Containers, pod.Spec.InitContainers...)
		valid              = true
	)

	for _, container := range relevantContainers {
		found, err := regexp.MatchString(regExpToLookFor, container.Image)
		if err != nil {
			panic(err.Error())
		}
		if !found {
			valid = false
			break
		}
	}

	return RuleEvaluation{
		Name:  "image_prefix",
		Valid: valid,
	}
}
