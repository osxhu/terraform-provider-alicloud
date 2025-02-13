package gpdb

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

// DescribeHealthStatus invokes the gpdb.DescribeHealthStatus API synchronously
func (client *Client) DescribeHealthStatus(request *DescribeHealthStatusRequest) (response *DescribeHealthStatusResponse, err error) {
	response = CreateDescribeHealthStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHealthStatusWithChan invokes the gpdb.DescribeHealthStatus API asynchronously
func (client *Client) DescribeHealthStatusWithChan(request *DescribeHealthStatusRequest) (<-chan *DescribeHealthStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeHealthStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHealthStatus(request)
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

// DescribeHealthStatusWithCallback invokes the gpdb.DescribeHealthStatus API asynchronously
func (client *Client) DescribeHealthStatusWithCallback(request *DescribeHealthStatusRequest, callback func(response *DescribeHealthStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHealthStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeHealthStatus(request)
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

// DescribeHealthStatusRequest is the request struct for api DescribeHealthStatus
type DescribeHealthStatusRequest struct {
	*requests.RpcRequest
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	Key          string `position:"Query" name:"Key"`
}

// DescribeHealthStatusResponse is the response struct for api DescribeHealthStatus
type DescribeHealthStatusResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DBClusterId string `json:"DBClusterId" xml:"DBClusterId"`
	Status      Status `json:"Status" xml:"Status"`
}

// CreateDescribeHealthStatusRequest creates a request to invoke DescribeHealthStatus API
func CreateDescribeHealthStatusRequest() (request *DescribeHealthStatusRequest) {
	request = &DescribeHealthStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeHealthStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeHealthStatusResponse creates a response to parse from DescribeHealthStatus response
func CreateDescribeHealthStatusResponse() (response *DescribeHealthStatusResponse) {
	response = &DescribeHealthStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
