package ddoscoo

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

// ConfigLayer4RealLimit invokes the ddoscoo.ConfigLayer4RealLimit API synchronously
func (client *Client) ConfigLayer4RealLimit(request *ConfigLayer4RealLimitRequest) (response *ConfigLayer4RealLimitResponse, err error) {
	response = CreateConfigLayer4RealLimitResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigLayer4RealLimitWithChan invokes the ddoscoo.ConfigLayer4RealLimit API asynchronously
func (client *Client) ConfigLayer4RealLimitWithChan(request *ConfigLayer4RealLimitRequest) (<-chan *ConfigLayer4RealLimitResponse, <-chan error) {
	responseChan := make(chan *ConfigLayer4RealLimitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigLayer4RealLimit(request)
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

// ConfigLayer4RealLimitWithCallback invokes the ddoscoo.ConfigLayer4RealLimit API asynchronously
func (client *Client) ConfigLayer4RealLimitWithCallback(request *ConfigLayer4RealLimitRequest, callback func(response *ConfigLayer4RealLimitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigLayer4RealLimitResponse
		var err error
		defer close(result)
		response, err = client.ConfigLayer4RealLimit(request)
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

// ConfigLayer4RealLimitRequest is the request struct for api ConfigLayer4RealLimit
type ConfigLayer4RealLimitRequest struct {
	*requests.RpcRequest
	SourceIp   string           `position:"Query" name:"SourceIp"`
	LimitValue requests.Integer `position:"Query" name:"LimitValue"`
	InstanceId string           `position:"Query" name:"InstanceId"`
}

// ConfigLayer4RealLimitResponse is the response struct for api ConfigLayer4RealLimit
type ConfigLayer4RealLimitResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigLayer4RealLimitRequest creates a request to invoke ConfigLayer4RealLimit API
func CreateConfigLayer4RealLimitRequest() (request *ConfigLayer4RealLimitRequest) {
	request = &ConfigLayer4RealLimitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ConfigLayer4RealLimit", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigLayer4RealLimitResponse creates a response to parse from ConfigLayer4RealLimit response
func CreateConfigLayer4RealLimitResponse() (response *ConfigLayer4RealLimitResponse) {
	response = &ConfigLayer4RealLimitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
