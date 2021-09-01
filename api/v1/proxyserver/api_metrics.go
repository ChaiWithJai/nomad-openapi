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
	"net/http"
	"strings"
)

// A MetricsApiController binds http requests to an api service and writes the service results to the http response
type MetricsApiController struct {
	service MetricsApiServicer
}

// NewMetricsApiController creates a default api controller
func NewMetricsApiController(s MetricsApiServicer) Router {
	return &MetricsApiController{service: s}
}

// Routes returns all of the api route for the MetricsApiController
func (c *MetricsApiController) Routes() Routes {
	return Routes{
		{
			"GetMetricsSummary",
			strings.ToUpper("Get"),
			"/v1/metrics",
			c.GetMetricsSummary,
		},
	}
}

// GetMetricsSummary -
func (c *MetricsApiController) GetMetricsSummary(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	format := query.Get("format")
	result, err := c.service.GetMetricsSummary(r.Context(), format)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, result.Headers, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}