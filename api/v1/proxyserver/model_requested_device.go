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

type RequestedDevice struct {
	Affinities []Affinity `json:"Affinities,omitempty"`

	Constraints []Constraint `json:"Constraints,omitempty"`

	Count int32 `json:"Count,omitempty"`

	Name string `json:"Name,omitempty"`
}