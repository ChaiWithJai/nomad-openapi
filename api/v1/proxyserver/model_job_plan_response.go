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

import (
	"time"
)

type JobPlanResponse struct {
	Annotations PlanAnnotations `json:"Annotations,omitempty"`

	CreatedEvals []Evaluation `json:"CreatedEvals,omitempty"`

	Diff JobDiff `json:"Diff,omitempty"`

	FailedTGAllocs map[string]AllocationMetric `json:"FailedTGAllocs,omitempty"`

	JobModifyIndex int32 `json:"JobModifyIndex,omitempty"`

	NextPeriodicLaunch time.Time `json:"NextPeriodicLaunch,omitempty"`

	Warnings string `json:"Warnings,omitempty"`
}