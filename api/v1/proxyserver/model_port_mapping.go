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

type PortMapping struct {
	HostIP string `json:"HostIP,omitempty"`

	Label string `json:"Label,omitempty"`

	To int32 `json:"To,omitempty"`

	Value int32 `json:"Value,omitempty"`
}