package state_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/grafana/grafana/pkg/services/annotations"
	"github.com/grafana/grafana/pkg/services/ngalert/eval"
	"github.com/grafana/grafana/pkg/services/ngalert/models"
	"github.com/grafana/grafana/pkg/services/ngalert/state"
	"github.com/grafana/grafana/pkg/services/ngalert/state/historian"
	"github.com/stretchr/testify/mock"
)

func BenchmarkProcessEvalResults(b *testing.B) {
	as := annotations.FakeAnnotationsRepo{}
	as.On("SaveMany", mock.Anything, mock.Anything).Return(nil)
	hist := historian.NewAnnotationBackend(&as, nil, nil)
	cfg := state.ManagerCfg{
		Historian: hist,
	}
	sut := state.NewManager(cfg)
	now := time.Now().UTC()
	rule := makeBenchRule()
	results := makeBenchResults(100)
	labels := map[string]string{}

	var ans []state.StateTransition
	for i := 0; i < b.N; i++ {
		ans = sut.ProcessEvalResults(context.Background(), now, &rule, results, labels)
	}

	b.StopTimer()

	_ = fmt.Sprintf("%v", len(ans))
}

func makeBenchRule() models.AlertRule {
	dashUID := "my-dash"
	panelID := int64(14)
	return models.AlertRule{
		ID:              5,
		OrgID:           1,
		Title:           "some rule",
		Condition:       "A",
		Data:            []models.AlertQuery{},
		Updated:         time.Now().UTC(),
		IntervalSeconds: 60,
		Version:         2,
		UID:             "abcd-efg",
		NamespaceUID:    "my-folder",
		DashboardUID:    &dashUID,
		PanelID:         &panelID,
		RuleGroup:       "my-group",
		RuleGroupIndex:  2,
		NoDataState:     models.NoData,
		ExecErrState:    models.ErrorErrState,
		For:             5 * time.Minute,
		Annotations: map[string]string{
			"text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
			"url":  "https://grafana.com",
		},
		Labels: map[string]string{
			"alertname": "some rule",
			"a":         "b",
			"cluster":   "prod-eu-west-123",
			"namespace": "coolthings",
		},
	}
}

func makeBenchResults(count int) eval.Results {
	labels := map[string]string{
		"alertname": "some rule",
		"a":         "b",
		"cluster":   "prod-eu-west-123",
		"namespace": "coolthings",
	}
	one := 1.0
	results := make([]eval.Result, 0, count)
	for i := 0; i < count; i++ {
		results = append(results, eval.Result{
			Instance:           labels,
			State:              eval.Alerting,
			EvaluatedAt:        time.Now().UTC(),
			EvaluationDuration: 5 * time.Second,
			Values: map[string]eval.NumberValueCapture{
				"A": eval.NumberValueCapture{
					Var:   "A",
					Value: &one,
				},
				"B": eval.NumberValueCapture{
					Var:   "B",
					Value: &one,
				},
				"C": eval.NumberValueCapture{
					Var:   "C",
					Value: &one,
				},
			},
		})
	}
	return results
}
