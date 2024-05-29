package config

var (
	Binary             = "/usr/bin/varnishncsa"
	WorkingDir         = "/run/varnish/"
	PrometheusLabels   = "{\"X-Real-Host\": \"host\", \"X-Frontend-Id\": \"frontend\"}"
	Format             = "{\"Timestamp\": \"%t\", \"Handling\": \"%{Varnish:handling}x\", \"Bytes\": \"%b\", \"X-Real-Host\": \"%{x-real-host}i\", \"X-Frontend-Id\": \"%{x-frontend-id}i\"}"
	HttpdServerAddress = "127.0.0.1"
	HttpdServerPort    = 8080
)
