package cloudapi

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

// DeleteBackend invokes the cloudapi.DeleteBackend API synchronously
func (client *Client) DeleteBackend(request *DeleteBackendRequest) (response *DeleteBackendResponse, err error) {
	response = CreateDeleteBackendResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBackendWithChan invokes the cloudapi.DeleteBackend API asynchronously
func (client *Client) DeleteBackendWithChan(request *DeleteBackendRequest) (<-chan *DeleteBackendResponse, <-chan error) {
	responseChan := make(chan *DeleteBackendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBackend(request)
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

// DeleteBackendWithCallback invokes the cloudapi.DeleteBackend API asynchronously
func (client *Client) DeleteBackendWithCallback(request *DeleteBackendRequest, callback func(response *DeleteBackendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBackendResponse
		var err error
		defer close(result)
		response, err = client.DeleteBackend(request)
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

// DeleteBackendRequest is the request struct for api DeleteBackend
type DeleteBackendRequest struct {
	*requests.RpcRequest
	BackendId     string `position:"Query" name:"BackendId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DeleteBackendResponse is the response struct for api DeleteBackend
type DeleteBackendResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteBackendRequest creates a request to invoke DeleteBackend API
func CreateDeleteBackendRequest() (request *DeleteBackendRequest) {
	request = &DeleteBackendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteBackend", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteBackendResponse creates a response to parse from DeleteBackend response
func CreateDeleteBackendResponse() (response *DeleteBackendResponse) {
	response = &DeleteBackendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
