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

type JobSummary struct {
	Children JobChildrenSummary `json:"Children,omitempty"`

	CreateIndex int32 `json:"CreateIndex,omitempty"`

	JobID string `json:"JobID,omitempty"`

	ModifyIndex int32 `json:"ModifyIndex,omitempty"`

	Namespace string `json:"Namespace,omitempty"`

	Summary map[string]TaskGroupSummary `json:"Summary,omitempty"`
}