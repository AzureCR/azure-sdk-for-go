package media

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ProvidersClient is the client for the Providers methods of the Media service.
type ProvidersClient struct {
	BaseClient
}

// NewProvidersClient creates an instance of the ProvidersClient client.
func NewProvidersClient(subscriptionID string) ProvidersClient {
	return NewProvidersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProvidersClientWithBaseURI creates an instance of the ProvidersClient client.
func NewProvidersClientWithBaseURI(baseURI string, subscriptionID string) ProvidersClient {
	return ProvidersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability checks whether the Media Service resource name is available.
//
// parameters is the request parameters
func (client ProvidersClient) CheckNameAvailability(ctx context.Context, parameters CheckNameAvailabilityInput) (result EntityNameAvailabilityCheckOutput, err error) {
	req, err := client.CheckNameAvailabilityPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ProvidersClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.ProvidersClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ProvidersClient", "CheckNameAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client ProvidersClient) CheckNameAvailabilityPreparer(ctx context.Context, parameters CheckNameAvailabilityInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Media/checkNameAvailability", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client ProvidersClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client ProvidersClient) CheckNameAvailabilityResponder(resp *http.Response) (result EntityNameAvailabilityCheckOutput, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
