/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.3
 * Contact: support@hashicorp.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

type MetricsSummary struct {
	Counters []SampledValue `json:"Counters,omitempty"`

	Gauges []GaugeValue `json:"Gauges,omitempty"`

	Points []PointValue `json:"Points,omitempty"`

	Samples []SampledValue `json:"Samples,omitempty"`

	Timestamp string `json:"Timestamp,omitempty"`
}