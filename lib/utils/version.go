package utils

import "github.com/criteo/varnishncsa-exporter/internal/config"

func VersionGet() string {
	return config.Version
}
