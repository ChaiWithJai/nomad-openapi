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

type ServiceCheck struct {
	AddressMode string `json:"AddressMode,omitempty"`

	Args []string `json:"Args,omitempty"`

	Body string `json:"Body,omitempty"`

	CheckRestart CheckRestart `json:"CheckRestart,omitempty"`

	Command string `json:"Command,omitempty"`

	Expose bool `json:"Expose,omitempty"`

	FailuresBeforeCritical int32 `json:"FailuresBeforeCritical,omitempty"`

	GRPCService string `json:"GRPCService,omitempty"`

	GRPCUseTLS bool `json:"GRPCUseTLS,omitempty"`

	Header map[string][]string `json:"Header,omitempty"`

	Id string `json:"Id,omitempty"`

	InitialStatus string `json:"InitialStatus,omitempty"`

	Interval int64 `json:"Interval,omitempty"`

	Method string `json:"Method,omitempty"`

	Name string `json:"Name,omitempty"`

	OnUpdate string `json:"OnUpdate,omitempty"`

	Path string `json:"Path,omitempty"`

	PortLabel string `json:"PortLabel,omitempty"`

	Protocol string `json:"Protocol,omitempty"`

	SuccessBeforePassing int32 `json:"SuccessBeforePassing,omitempty"`

	TLSSkipVerify bool `json:"TLSSkipVerify,omitempty"`

	TaskName string `json:"TaskName,omitempty"`

	Timeout int64 `json:"Timeout,omitempty"`

	Type string `json:"Type,omitempty"`
}