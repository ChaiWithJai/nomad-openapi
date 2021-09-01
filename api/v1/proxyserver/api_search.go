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
	"encoding/json"
	"net/http"
	"strings"
)

// A SearchApiController binds http requests to an api service and writes the service results to the http response
type SearchApiController struct {
	service SearchApiServicer
}

// NewSearchApiController creates a default api controller
func NewSearchApiController(s SearchApiServicer) Router {
	return &SearchApiController{service: s}
}

// Routes returns all of the api route for the SearchApiController
func (c *SearchApiController) Routes() Routes {
	return Routes{
		{
			"GetFuzzySearch",
			strings.ToUpper("Post"),
			"/v1/search/fuzzy",
			c.GetFuzzySearch,
		},
		{
			"GetSearch",
			strings.ToUpper("Post"),
			"/v1/search",
			c.GetSearch,
		},
	}
}

// GetFuzzySearch -
func (c *SearchApiController) GetFuzzySearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fuzzySearchRequest := &FuzzySearchRequest{}
	if err := json.NewDecoder(r.Body).Decode(&fuzzySearchRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	region := query.Get("region")
	namespace := query.Get("namespace")
	index, err := parseInt32Parameter(query.Get("index"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	wait := query.Get("wait")
	stale := query.Get("stale")
	prefix := query.Get("prefix")
	xNomadToken := r.Header.Get("X-Nomad-Token")
	perPage, err := parseInt32Parameter(query.Get("per_page"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	nextToken := query.Get("next_token")
	result, err := c.service.GetFuzzySearch(r.Context(), *fuzzySearchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, result.Headers, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetSearch -
func (c *SearchApiController) GetSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	searchRequest := &SearchRequest{}
	if err := json.NewDecoder(r.Body).Decode(&searchRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	region := query.Get("region")
	namespace := query.Get("namespace")
	index, err := parseInt32Parameter(query.Get("index"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	wait := query.Get("wait")
	stale := query.Get("stale")
	prefix := query.Get("prefix")
	xNomadToken := r.Header.Get("X-Nomad-Token")
	perPage, err := parseInt32Parameter(query.Get("per_page"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	nextToken := query.Get("next_token")
	result, err := c.service.GetSearch(r.Context(), *searchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, result.Headers, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}