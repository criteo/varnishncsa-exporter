package prometheus

import (
	"fmt"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestProcessMetric(t *testing.T) {
	assert := assert.New(t)
	counters := InitPrometheusCounters(false, []string{"frontend", "host"})
	labelsMapping := make(map[string]string)
	labelsMapping["X-Frontend-Id"] = "frontend"
	labelsMapping["X-Real-Host"] = "host"

	host1 := "lb-control-plane.preprod.crto.in"
	frontendID1 := "zeroconf"
	rawMetric := "{\"Handling\": \"hit\", \"X-Real-Host\": \"" + host1 + "\", \"X-Frontend-Id\": \"" + frontendID1 + "\"}"
	i := 1
	for i <= 5 {
		ProcessMetric([]byte(rawMetric), counters, labelsMapping)
		assert.Equal(1, testutil.CollectAndCount(counters.HitCounter.With(prometheus.Labels{"host": host1, "frontend": frontendID1})), fmt.Sprintf("Count - Iteration %d", i))
		assert.Equal(float64(i), testutil.ToFloat64(counters.HitCounter), fmt.Sprintf("Value - Iteration %d", i))
		assert.Equal(0, testutil.CollectAndCount(counters.HitMissCounter), fmt.Sprintf("Iteration %d", i))
		assert.Equal(0, testutil.CollectAndCount(counters.MissCounter), fmt.Sprintf("Iteration %d", i))
		assert.Equal(0, testutil.CollectAndCount(counters.PassCounter), fmt.Sprintf("Iteration %d", i))
		assert.Equal(0, testutil.CollectAndCount(counters.PipeCounter), fmt.Sprintf("Iteration %d", i))
		i++
	}

	host2 := "static.criteo.net"
	frontendID2 := "static.crto.in"
	rawMetric = "{\"Handling\": \"hit\", \"X-Real-Host\": \"" + host2 + "\", \"X-Frontend-Id\": \"" + frontendID2 + "\"}"
	ProcessMetric([]byte(rawMetric), counters, labelsMapping)
	assert.Equal(1, testutil.CollectAndCount(counters.HitCounter.With(prometheus.Labels{"host": host1, "frontend": frontendID1})))
	assert.Equal(float64(5), testutil.ToFloat64(counters.HitCounter.With(prometheus.Labels{"host": host1, "frontend": frontendID1})))
	assert.Equal(1, testutil.CollectAndCount(counters.HitCounter.With(prometheus.Labels{"host": host2, "frontend": frontendID2})))
	assert.Equal(float64(1), testutil.ToFloat64(counters.HitCounter.With(prometheus.Labels{"host": host2, "frontend": frontendID2})))
	assert.Equal(0, testutil.CollectAndCount(counters.HitMissCounter))
	assert.Equal(0, testutil.CollectAndCount(counters.MissCounter))
	assert.Equal(0, testutil.CollectAndCount(counters.PassCounter))
	assert.Equal(0, testutil.CollectAndCount(counters.PipeCounter))

	rawMetric = "{\"Handling\": \"miss\", \"X-Real-Host\": \"" + host1 + "\", \"X-Frontend-Id\": \"" + frontendID1 + "\"}"
	ProcessMetric([]byte(rawMetric), counters, labelsMapping)
	assert.Equal(1, testutil.CollectAndCount(counters.HitCounter.With(prometheus.Labels{"host": host1, "frontend": frontendID1})))
	assert.Equal(float64(5), testutil.ToFloat64(counters.HitCounter.With(prometheus.Labels{"host": host1, "frontend": frontendID1})))
	assert.Equal(1, testutil.CollectAndCount(counters.HitCounter.With(prometheus.Labels{"host": host2, "frontend": frontendID2})))
	assert.Equal(float64(1), testutil.ToFloat64(counters.HitCounter.With(prometheus.Labels{"host": host2, "frontend": frontendID2})))
	assert.Equal(0, testutil.CollectAndCount(counters.HitMissCounter))
	assert.Equal(1, testutil.CollectAndCount(counters.MissCounter))
	assert.Equal(0, testutil.CollectAndCount(counters.PassCounter))
	assert.Equal(0, testutil.CollectAndCount(counters.PipeCounter))

	expectation := strings.NewReader(`
    # HELP hit_count_total Hit count.
    # TYPE hit_count_total counter
    hit_count_total{frontend="zeroconf",host="lb-control-plane.preprod.crto.in"} 5
	`)
	assert.NoError(testutil.CollectAndCompare(counters.HitCounter.With(prometheus.Labels{"host": host1, "frontend": frontendID1}), expectation, "hit_count_total"))

	expectation = strings.NewReader(`
    # HELP miss_count_total Miss count.
    # TYPE miss_count_total counter
    miss_count_total{frontend="zeroconf",host="lb-control-plane.preprod.crto.in"} 1
	`)
	assert.NoError(testutil.CollectAndCompare(counters.MissCounter, expectation, "miss_count_total"))
}
