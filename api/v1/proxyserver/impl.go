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

//Implementation response defines an error code with the associated body
type ImplResponse struct {
	Code    int
	Headers map[string][]string
	Body    interface{}
}