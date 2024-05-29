package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/criteo/varnishncsa-exporter/internal/config"
	"github.com/criteo/varnishncsa-exporter/lib/command"
	"github.com/criteo/varnishncsa-exporter/lib/prometheus"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"

	log "github.com/sirupsen/logrus"
)

var (
	version = "unknown"
	commit  = "unknown"
	date    = "unknown"
)

func versionGet() string {
	return fmt.Sprintf("Version: %s - Commit: %s - Date: %s", version, commit, date)
}

func main() {
	app := &cli.App{
		Name:  "varnishncsa_exporter",
		Usage: "Exposes prometheus metrics for Varnish by parsing structured logs output from varnishncsa.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases:     []string{"b"},
				Category:    "Piped daemon:",
				DefaultText: config.Binary,
				Name:        "binary",
				Required:    false,
				Usage:       "Binary to run and pipe output",
				Value:       config.Binary,
			},
			&cli.BoolFlag{
				Aliases:  []string{"D"},
				Category: "Miscellaneous:",
				Name:     "debug",
				Usage:    "show debug output",
			},
			&cli.StringFlag{
				Aliases:     []string{"n"},
				Category:    "Piped daemon:",
				DefaultText: config.WorkingDir,
				Name:        "directory",
				Required:    false,
				Usage:       "Varnishd working directory",
				Value:       config.WorkingDir,
			},
			&cli.StringFlag{
				Aliases:     []string{"F"},
				Category:    "Piped daemon:",
				DefaultText: config.Format,
				Name:        "format",
				Required:    false,
				Usage:       "Set the output log format string",
				Value:       config.Format,
			},
			&cli.StringFlag{
				Aliases:     []string{"L"},
				Category:    "Piped daemon:",
				DefaultText: config.PrometheusLabels,
				Name:        "labels",
				Required:    false,
				Usage:       "Prometheus labels mapping key, value represented  in json",
				Value:       config.PrometheusLabels,
			},
			&cli.StringFlag{
				Action: func(ctx *cli.Context, address string) error {
					if address == "localhost" {
						return nil
					} else if net.ParseIP(address) == nil {
						return fmt.Errorf(fmt.Sprintf("IP Address: %s - Invalid", address))
					}
					return nil
				},
				Aliases:     []string{"a"},
				Category:    "Prometheus server:",
				DefaultText: config.HttpdServerAddress,
				Name:        "httpd_address",
				Required:    false,
				Usage:       "Prometheus HTTP server address",
				Value:       config.HttpdServerAddress,
			},
			&cli.IntFlag{
				Action: func(ctx *cli.Context, port int) error {
					if port < 1024 || port > 49151 {
						return fmt.Errorf(fmt.Sprintf("Port: %d - Invalid - 1024 <= port <= 49151", port))
					}
					return nil
				},
				Aliases:     []string{"p"},
				Category:    "Prometheus server:",
				DefaultText: fmt.Sprintf("%v", config.HttpdServerPort),
				Name:        "httpd_port",
				Required:    false,
				Usage:       "Prometheus HTTP server port",
				Value:       config.HttpdServerPort,
			},
			&cli.BoolFlag{
				Name:     "version",
				Category: "Miscellaneous:",
				Usage:    "Print version",
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.Bool("version") {
				return cli.Exit(versionGet(), 0)
			}
			if ctx.Bool("debug") {
				log.SetLevel(log.DebugLevel)
			}
			prometheusLabelsMapping := ctx.String("labels")
			prometheusLabels := make(map[string]string)
			if err := json.Unmarshal([]byte(prometheusLabelsMapping), &prometheusLabels); err != nil {
				message := fmt.Sprintf("Failed to unmarshal prometheusLabels %s", err)
				log.Error(message)
				return cli.Exit(message, 1)
			}
			fields := []string{}
			for _, v := range prometheusLabels {
				fields = append(fields, v)
			}
			counters := prometheus.InitPrometheusCounters(true, fields)

			var wg sync.WaitGroup
			wg.Add(1)
			log.Info(versionGet())
			go command.RunCommand(ctx.String("binary"), []string{"-n", ctx.String("directory"), "-F", ctx.String("format")}, counters, prometheusLabels, true, &wg)

			httpAddressPort := fmt.Sprintf("%s:%d", ctx.String("httpd_address"), ctx.Int("httpd_port"))
			log.Info("Starting prometheus server on ", httpAddressPort)
			router := httprouter.New()
			router.Handler("GET", "/metrics", promhttp.Handler())
			router.GET("/health", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
				w.WriteHeader(http.StatusOK)
			})
			server := &http.Server{
				Addr:              httpAddressPort,
				ReadHeaderTimeout: 20 * time.Second,
				Handler:           router,
			}
			if err := server.ListenAndServe(); err != nil {
				return cli.Exit(fmt.Sprintf("Could not listen on %s: %v", httpAddressPort, err), 1)
			}

			return cli.Exit(nil, 0)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
