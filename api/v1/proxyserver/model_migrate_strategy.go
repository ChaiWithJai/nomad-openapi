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

type MigrateStrategy struct {
	HealthCheck string `json:"HealthCheck,omitempty"`

	HealthyDeadline int64 `json:"HealthyDeadline,omitempty"`

	MaxParallel int32 `json:"MaxParallel,omitempty"`

	MinHealthyTime int64 `json:"MinHealthyTime,omitempty"`
}