package cbn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateCenInterRegionTrafficQosPolicyAttribute invokes the cbn.UpdateCenInterRegionTrafficQosPolicyAttribute API synchronously
func (client *Client) UpdateCenInterRegionTrafficQosPolicyAttribute(request *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) (response *UpdateCenInterRegionTrafficQosPolicyAttributeResponse, err error) {
	response = CreateUpdateCenInterRegionTrafficQosPolicyAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCenInterRegionTrafficQosPolicyAttributeWithChan invokes the cbn.UpdateCenInterRegionTrafficQosPolicyAttribute API asynchronously
func (client *Client) UpdateCenInterRegionTrafficQosPolicyAttributeWithChan(request *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) (<-chan *UpdateCenInterRegionTrafficQosPolicyAttributeResponse, <-chan error) {
	responseChan := make(chan *UpdateCenInterRegionTrafficQosPolicyAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCenInterRegionTrafficQosPolicyAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateCenInterRegionTrafficQosPolicyAttributeWithCallback invokes the cbn.UpdateCenInterRegionTrafficQosPolicyAttribute API asynchronously
func (client *Client) UpdateCenInterRegionTrafficQosPolicyAttributeWithCallback(request *UpdateCenInterRegionTrafficQosPolicyAttributeRequest, callback func(response *UpdateCenInterRegionTrafficQosPolicyAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCenInterRegionTrafficQosPolicyAttributeResponse
		var err error
		defer close(result)
		response, err = client.UpdateCenInterRegionTrafficQosPolicyAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateCenInterRegionTrafficQosPolicyAttributeRequest is the request struct for api UpdateCenInterRegionTrafficQosPolicyAttribute
type UpdateCenInterRegionTrafficQosPolicyAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId             requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                 string           `position:"Query" name:"ClientToken"`
	TrafficQosPolicyName        string           `position:"Query" name:"TrafficQosPolicyName"`
	DryRun                      requests.Boolean `position:"Query" name:"DryRun"`
	TrafficQosPolicyId          string           `position:"Query" name:"TrafficQosPolicyId"`
	ResourceOwnerAccount        string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string           `position:"Query" name:"OwnerAccount"`
	TrafficQosPolicyDescription string           `position:"Query" name:"TrafficQosPolicyDescription"`
	OwnerId                     requests.Integer `position:"Query" name:"OwnerId"`
	Version                     string           `position:"Query" name:"Version"`
}

// UpdateCenInterRegionTrafficQosPolicyAttributeResponse is the response struct for api UpdateCenInterRegionTrafficQosPolicyAttribute
type UpdateCenInterRegionTrafficQosPolicyAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateCenInterRegionTrafficQosPolicyAttributeRequest creates a request to invoke UpdateCenInterRegionTrafficQosPolicyAttribute API
func CreateUpdateCenInterRegionTrafficQosPolicyAttributeRequest() (request *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) {
	request = &UpdateCenInterRegionTrafficQosPolicyAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "UpdateCenInterRegionTrafficQosPolicyAttribute", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateCenInterRegionTrafficQosPolicyAttributeResponse creates a response to parse from UpdateCenInterRegionTrafficQosPolicyAttribute response
func CreateUpdateCenInterRegionTrafficQosPolicyAttributeResponse() (response *UpdateCenInterRegionTrafficQosPolicyAttributeResponse) {
	response = &UpdateCenInterRegionTrafficQosPolicyAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
