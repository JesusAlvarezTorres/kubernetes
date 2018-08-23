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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// LoadBalancerNetworkInterfacesClient is the network Client
type LoadBalancerNetworkInterfacesClient struct {
	ManagementClient
}

// NewLoadBalancerNetworkInterfacesClient creates an instance of the LoadBalancerNetworkInterfacesClient client.
func NewLoadBalancerNetworkInterfacesClient(subscriptionID string) LoadBalancerNetworkInterfacesClient {
	return NewLoadBalancerNetworkInterfacesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLoadBalancerNetworkInterfacesClientWithBaseURI creates an instance of the LoadBalancerNetworkInterfacesClient
// client.
func NewLoadBalancerNetworkInterfacesClientWithBaseURI(baseURI string, subscriptionID string) LoadBalancerNetworkInterfacesClient {
	return LoadBalancerNetworkInterfacesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets associated load balancer network interfaces.
//
// resourceGroupName is the name of the resource group. loadBalancerName is the name of the load balancer.
func (client LoadBalancerNetworkInterfacesClient) List(resourceGroupName string, loadBalancerName string) (result InterfaceListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName, loadBalancerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client LoadBalancerNetworkInterfacesClient) ListPreparer(resourceGroupName string, loadBalancerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"loadBalancerName":  autorest.Encode("path", loadBalancerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/networkInterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerNetworkInterfacesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LoadBalancerNetworkInterfacesClient) ListResponder(resp *http.Response) (result InterfaceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client LoadBalancerNetworkInterfacesClient) ListNextResults(lastResults InterfaceListResult) (result InterfaceListResult, err error) {
	req, err := lastResults.InterfaceListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client LoadBalancerNetworkInterfacesClient) ListComplete(resourceGroupName string, loadBalancerName string, cancel <-chan struct{}) (<-chan Interface, <-chan error) {
	resultChan := make(chan Interface)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName, loadBalancerName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
