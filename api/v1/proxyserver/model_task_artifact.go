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

type TaskArtifact struct {
	GetterHeaders map[string]string `json:"GetterHeaders,omitempty"`

	GetterMode string `json:"GetterMode,omitempty"`

	GetterOptions map[string]string `json:"GetterOptions,omitempty"`

	GetterSource string `json:"GetterSource,omitempty"`

	RelativeDest string `json:"RelativeDest,omitempty"`
}