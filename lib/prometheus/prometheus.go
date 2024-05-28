package prometheus

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type VarnishMetric struct {
	HitStatus string         `mapstructure:"Handling"`
	Other     map[string]any `mapstructure:",remain"`
}

type PrometheusCounters struct {
	HitCounter     prometheus.CounterVec
	HitMissCounter prometheus.CounterVec
	MissCounter    prometheus.CounterVec
	PassCounter    prometheus.CounterVec
	PipeCounter    prometheus.CounterVec
	SynthCounter   prometheus.CounterVec
}

func _initPrometheusCounter(name string, help string, fields []string, panicOnError bool) prometheus.CounterVec {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
		fields,
	)
	err := prometheus.Register(counter)
	if err != nil && panicOnError {
		panic(err)
	}
	return *counter
}

func InitPrometheusCounters(panicOnError bool, fields []string) PrometheusCounters {
	return PrometheusCounters{
		HitCounter:     _initPrometheusCounter("hit_count_total", "Hit count.", fields, panicOnError),
		HitMissCounter: _initPrometheusCounter("hit_miss_count_total", "Hit Miss count.", fields, panicOnError),
		MissCounter:    _initPrometheusCounter("miss_count_total", "Miss count.", fields, panicOnError),
		PassCounter:    _initPrometheusCounter("pass_count_total", "Pass count.", fields, panicOnError),
		PipeCounter:    _initPrometheusCounter("pipe_count_total", "Pipe count.", fields, panicOnError),
		SynthCounter:   _initPrometheusCounter("synth_count_total", "Synth count.", fields, panicOnError),
	}
}

func ProcessMetric(rawMetric []byte, counters PrometheusCounters, labelsMapping map[string]string) {
	var varnishMetric VarnishMetric
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(rawMetric, &jsonMap); err != nil {
		log.Error(fmt.Printf("Failed to unmarshal rawMetric %s", err))
		return
	}
	if err := mapstructure.Decode(jsonMap, &varnishMetric); err != nil {
		log.Error(fmt.Printf("Failed to decode json %s: %s", jsonMap, err))
		return
	}
	log.Debug(varnishMetric)
	labelsStr := make(map[string]string)
	var ok bool
	for k, v := range labelsMapping {
		if labelsStr[v], ok = varnishMetric.Other[k].(string); !ok {
			log.Error(fmt.Printf("Failed to create labels for %v: conversion to string failed", varnishMetric.Other[k]))
			return
		}
	}
	labels := prometheus.Labels(labelsStr)

	switch varnishMetric.HitStatus {
	case "hit":
		counters.HitCounter.With(labels).Inc()
	case "hitmiss":
		counters.HitMissCounter.With(labels).Inc()
	case "miss":
		counters.MissCounter.With(labels).Inc()
	case "pass":
		counters.PassCounter.With(labels).Inc()
	case "pipe":
		counters.PipeCounter.With(labels).Inc()
	case "synth":
		counters.SynthCounter.With(labels).Inc()
	default:
		log.Warning(fmt.Printf("Status type %s not supported", varnishMetric.HitStatus))
	}
}
