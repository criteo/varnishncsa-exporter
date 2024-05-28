package command

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"

	"github.com/criteo/varnishncsa-exporter/lib/prometheus"
	log "github.com/sirupsen/logrus"
)

var execCommand = exec.Command

func RunCommand(binary string, options []string, counters prometheus.PrometheusCounters, labelsMapping map[string]string, exitOnLookPathFailure bool, wg *sync.WaitGroup) {
	path, err := exec.LookPath(binary)
	if err != nil && exitOnLookPathFailure {
		log.Fatal(fmt.Sprintf("An error occured while looking for the binary: %s", err))
	}
	if wg == nil {
		log.Fatal(fmt.Print("passed waitgroup object is nil"))
	}
	log.Info("Executing command:\n\t", binary, " ", strings.Join(options, " "))
	cmd := execCommand(path, options...)

	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(fmt.Sprintf("An error occured while fetching stdout pipe: %s", err))
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(fmt.Sprintf("An error occured while starting command: %s", err))
	}

	go func() {
		readPipe(cmdOut, counters, labelsMapping)
		if err := cmd.Wait(); err != nil {
			log.Fatal(fmt.Sprintf("An error occured while executing the command: %s", err))
		}
		defer wg.Done()
	}()
}

func readPipe(reader io.Reader, counters prometheus.PrometheusCounters, labelsMapping map[string]string) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		prometheus.ProcessMetric([]byte(line), counters, labelsMapping)
		log.Debug(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(fmt.Sprintf("Error reading from pipe: %v", err))
	}
}
