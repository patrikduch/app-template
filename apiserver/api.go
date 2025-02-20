/*
 * App template API
 *
 * API to access and configure the app template
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

import (
	"context"
	"net/http"
)



// ConfigurationApiRouter defines the required methods for binding the api requests to a responses for the ConfigurationApi
// The ConfigurationApiRouter implementation should parse necessary information from the http request,
// pass the data to a ConfigurationApiServicer to perform the required actions, then write the service results to the http response.
type ConfigurationApiRouter interface { 
	GetBTC(http.ResponseWriter, *http.Request)
	GetExamples(http.ResponseWriter, *http.Request)
	PostExample(http.ResponseWriter, *http.Request)
}


// ConfigurationApiServicer defines the api actions for the ConfigurationApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ConfigurationApiServicer interface { 
	GetBTC(context.Context) (ImplResponse, error)
	GetExamples(context.Context) (ImplResponse, error)
	PostExample(context.Context, Example) (ImplResponse, error)
}
