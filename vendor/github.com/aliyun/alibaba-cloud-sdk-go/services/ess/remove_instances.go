package ess

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

// RemoveInstances invokes the ess.RemoveInstances API synchronously
func (client *Client) RemoveInstances(request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
	response = CreateRemoveInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveInstancesWithChan invokes the ess.RemoveInstances API asynchronously
func (client *Client) RemoveInstancesWithChan(request *RemoveInstancesRequest) (<-chan *RemoveInstancesResponse, <-chan error) {
	responseChan := make(chan *RemoveInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveInstances(request)
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

// RemoveInstancesWithCallback invokes the ess.RemoveInstances API asynchronously
func (client *Client) RemoveInstancesWithCallback(request *RemoveInstancesRequest, callback func(response *RemoveInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveInstancesResponse
		var err error
		defer close(result)
		response, err = client.RemoveInstances(request)
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

// RemoveInstancesRequest is the request struct for api RemoveInstances
type RemoveInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken             string           `position:"Query" name:"ClientToken"`
	ScalingGroupId          string           `position:"Query" name:"ScalingGroupId"`
	DecreaseDesiredCapacity requests.Boolean `position:"Query" name:"DecreaseDesiredCapacity"`
	RemovePolicy            string           `position:"Query" name:"RemovePolicy"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId              *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
}

// RemoveInstancesResponse is the response struct for api RemoveInstances
type RemoveInstancesResponse struct {
	*responses.BaseResponse
	ScalingActivityId string `json:"ScalingActivityId" xml:"ScalingActivityId"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveInstancesRequest creates a request to invoke RemoveInstances API
func CreateRemoveInstancesRequest() (request *RemoveInstancesRequest) {
	request = &RemoveInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "RemoveInstances", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveInstancesResponse creates a response to parse from RemoveInstances response
func CreateRemoveInstancesResponse() (response *RemoveInstancesResponse) {
	response = &RemoveInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
