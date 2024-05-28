package command

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"testing"

	"github.com/criteo/varnishncsa-exporter/lib/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...) // #nosec G204
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	fmt.Println(`{"Timestamp": "[24/May/2024:07:49:57 +0000]", "Handling": "hit", "Request": "GET http://netsvc-probe-target-all.da1.preprod.crto.in/cache/test-lb-consul-cache?cache-control=stale-while-revalidate=5,max-age=10 HTTP/1.1", "X-Real-Host": "netsvc-probe-target-all.da1.preprod.crto.in", "X-Frontend-Id": "zeroconf"}`)
}

func TestRunCommand(t *testing.T) {
	assert := assert.New(t)
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	counters := prometheus.InitPrometheusCounters(false, []string{"frontend", "host"})
	labelsMapping := make(map[string]string)
	labelsMapping["X-Frontend-Id"] = "frontend"
	labelsMapping["X-Real-Host"] = "host"

	var wg sync.WaitGroup
	wg.Add(1)
	RunCommand("varnishncsa", []string{}, counters, labelsMapping, false, &wg)
	wg.Wait()
	assert.Equal(1, testutil.CollectAndCount(counters.HitCounter))
	assert.Equal(float64(1), testutil.ToFloat64(counters.HitCounter))
}
