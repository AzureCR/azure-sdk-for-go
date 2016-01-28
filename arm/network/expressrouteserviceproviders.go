package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// ExpressRouteServiceProvidersClient is the the Microsoft Azure Network
// management API provides a RESTful set of web services that interact with
// Microsoft Azure Networks service to manage your network resrources. The
// API has entities that capture the relationship between an end user and the
// Microsoft Azure Networks service.
type ExpressRouteServiceProvidersClient struct {
	ManagementClient
}

// NewExpressRouteServiceProvidersClient creates an instance of the
// ExpressRouteServiceProvidersClient client.
func NewExpressRouteServiceProvidersClient(subscriptionID string) ExpressRouteServiceProvidersClient {
	return NewExpressRouteServiceProvidersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteServiceProvidersClientWithBaseURI creates an instance of the
// ExpressRouteServiceProvidersClient client.
func NewExpressRouteServiceProvidersClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteServiceProvidersClient {
	return ExpressRouteServiceProvidersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List the List ExpressRouteServiceProvider opertion retrieves all the
// available ExpressRouteServiceProviders.
func (client ExpressRouteServiceProvidersClient) List() (result ExpressRouteServiceProviderListResult, ae error) {
	req, err := client.ListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteServiceProvidersClient", "List", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteServiceProvidersClient", "List", resp.StatusCode, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteServiceProvidersClient", "List", resp.StatusCode, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteServiceProvidersClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteServiceProvidersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteServiceProvidersClient) ListResponder(resp *http.Response) (result ExpressRouteServiceProviderListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ExpressRouteServiceProvidersClient) ListNextResults(lastResults ExpressRouteServiceProviderListResult) (result ExpressRouteServiceProviderListResult, ae error) {
	req, err := lastResults.ExpressRouteServiceProviderListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteServiceProvidersClient", "List", autorest.UndefinedStatusCode, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteServiceProvidersClient", "List", resp.StatusCode, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteServiceProvidersClient", "List", resp.StatusCode, "Failure responding to next results request request")
	}

	return
}
