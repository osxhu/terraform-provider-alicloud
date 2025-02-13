package cdn

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

// RollbackStagingConfig invokes the cdn.RollbackStagingConfig API synchronously
func (client *Client) RollbackStagingConfig(request *RollbackStagingConfigRequest) (response *RollbackStagingConfigResponse, err error) {
	response = CreateRollbackStagingConfigResponse()
	err = client.DoAction(request, response)
	return
}

// RollbackStagingConfigWithChan invokes the cdn.RollbackStagingConfig API asynchronously
func (client *Client) RollbackStagingConfigWithChan(request *RollbackStagingConfigRequest) (<-chan *RollbackStagingConfigResponse, <-chan error) {
	responseChan := make(chan *RollbackStagingConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RollbackStagingConfig(request)
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

// RollbackStagingConfigWithCallback invokes the cdn.RollbackStagingConfig API asynchronously
func (client *Client) RollbackStagingConfigWithCallback(request *RollbackStagingConfigRequest, callback func(response *RollbackStagingConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RollbackStagingConfigResponse
		var err error
		defer close(result)
		response, err = client.RollbackStagingConfig(request)
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

// RollbackStagingConfigRequest is the request struct for api RollbackStagingConfig
type RollbackStagingConfigRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
}

// RollbackStagingConfigResponse is the response struct for api RollbackStagingConfig
type RollbackStagingConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRollbackStagingConfigRequest creates a request to invoke RollbackStagingConfig API
func CreateRollbackStagingConfigRequest() (request *RollbackStagingConfigRequest) {
	request = &RollbackStagingConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "RollbackStagingConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateRollbackStagingConfigResponse creates a response to parse from RollbackStagingConfig response
func CreateRollbackStagingConfigResponse() (response *RollbackStagingConfigResponse) {
	response = &RollbackStagingConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
