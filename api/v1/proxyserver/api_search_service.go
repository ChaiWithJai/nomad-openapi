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
	"context"
	"errors"
	"net/http"
)

// SearchApiService is a service that implents the logic for the SearchApiServicer
// This service should implement the business logic for every endpoint for the SearchApi API.
// Include any external packages or services that will be required by this service.
type SearchApiService struct {
}

// NewSearchApiService creates a default api service
func NewSearchApiService() SearchApiServicer {
	return &SearchApiService{}
}

// GetFuzzySearch -
func (s *SearchApiService) GetFuzzySearch(ctx context.Context, fuzzySearchRequest FuzzySearchRequest, region string, namespace string, index int32, wait string, stale string, prefix string, xNomadToken string, perPage int32, nextToken string) (ImplResponse, error) {
	// TODO - update GetFuzzySearch with the required logic for this service method.
	// Add api_search_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, FuzzySearchResponse{}) or use other options such as http.Ok ...
	//return Response(200, FuzzySearchResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(403, {}) or use other options such as http.Ok ...
	//return Response(403, nil),nil

	//TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	//return Response(405, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetFuzzySearch method not implemented")
}

// GetSearch -
func (s *SearchApiService) GetSearch(ctx context.Context, searchRequest SearchRequest, region string, namespace string, index int32, wait string, stale string, prefix string, xNomadToken string, perPage int32, nextToken string) (ImplResponse, error) {
	// TODO - update GetSearch with the required logic for this service method.
	// Add api_search_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, SearchResponse{}) or use other options such as http.Ok ...
	//return Response(200, SearchResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(403, {}) or use other options such as http.Ok ...
	//return Response(403, nil),nil

	//TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	//return Response(405, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetSearch method not implemented")
}