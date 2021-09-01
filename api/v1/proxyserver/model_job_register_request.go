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

type JobRegisterRequest struct {
	EnforceIndex bool `json:"EnforceIndex,omitempty"`

	Job Job `json:"Job,omitempty"`

	JobModifyIndex int32 `json:"JobModifyIndex,omitempty"`

	Namespace string `json:"Namespace,omitempty"`

	PolicyOverride bool `json:"PolicyOverride,omitempty"`

	PreserveCounts bool `json:"PreserveCounts,omitempty"`

	Region string `json:"Region,omitempty"`

	SecretID string `json:"SecretID,omitempty"`
}