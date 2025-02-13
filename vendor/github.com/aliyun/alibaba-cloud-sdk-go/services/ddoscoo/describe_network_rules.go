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

// DescribeNetworkRules invokes the ddoscoo.DescribeNetworkRules API synchronously
func (client *Client) DescribeNetworkRules(request *DescribeNetworkRulesRequest) (response *DescribeNetworkRulesResponse, err error) {
	response = CreateDescribeNetworkRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkRulesWithChan invokes the ddoscoo.DescribeNetworkRules API asynchronously
func (client *Client) DescribeNetworkRulesWithChan(request *DescribeNetworkRulesRequest) (<-chan *DescribeNetworkRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkRules(request)
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

// DescribeNetworkRulesWithCallback invokes the ddoscoo.DescribeNetworkRules API asynchronously
func (client *Client) DescribeNetworkRulesWithCallback(request *DescribeNetworkRulesRequest, callback func(response *DescribeNetworkRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkRules(request)
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

// DescribeNetworkRulesRequest is the request struct for api DescribeNetworkRules
type DescribeNetworkRulesRequest struct {
	*requests.RpcRequest
	IsOffset        requests.Boolean `position:"Query" name:"IsOffset"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	FrontendPort    requests.Integer `position:"Query" name:"FrontendPort"`
	ForwardProtocol string           `position:"Query" name:"ForwardProtocol"`
}

// DescribeNetworkRulesResponse is the response struct for api DescribeNetworkRules
type DescribeNetworkRulesResponse struct {
	*responses.BaseResponse
	TotalCount   int64         `json:"TotalCount" xml:"TotalCount"`
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	NetworkRules []NetworkRule `json:"NetworkRules" xml:"NetworkRules"`
}

// CreateDescribeNetworkRulesRequest creates a request to invoke DescribeNetworkRules API
func CreateDescribeNetworkRulesRequest() (request *DescribeNetworkRulesRequest) {
	request = &DescribeNetworkRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeNetworkRules", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNetworkRulesResponse creates a response to parse from DescribeNetworkRules response
func CreateDescribeNetworkRulesResponse() (response *DescribeNetworkRulesResponse) {
	response = &DescribeNetworkRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
