package dcdn

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

// DeleteRoutine invokes the dcdn.DeleteRoutine API synchronously
func (client *Client) DeleteRoutine(request *DeleteRoutineRequest) (response *DeleteRoutineResponse, err error) {
	response = CreateDeleteRoutineResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRoutineWithChan invokes the dcdn.DeleteRoutine API asynchronously
func (client *Client) DeleteRoutineWithChan(request *DeleteRoutineRequest) (<-chan *DeleteRoutineResponse, <-chan error) {
	responseChan := make(chan *DeleteRoutineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRoutine(request)
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

// DeleteRoutineWithCallback invokes the dcdn.DeleteRoutine API asynchronously
func (client *Client) DeleteRoutineWithCallback(request *DeleteRoutineRequest, callback func(response *DeleteRoutineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRoutineResponse
		var err error
		defer close(result)
		response, err = client.DeleteRoutine(request)
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

// DeleteRoutineRequest is the request struct for api DeleteRoutine
type DeleteRoutineRequest struct {
	*requests.RpcRequest
	Name string `position:"Body" name:"Name"`
}

// DeleteRoutineResponse is the response struct for api DeleteRoutine
type DeleteRoutineResponse struct {
	*responses.BaseResponse
	Content   map[string]interface{} `json:"Content" xml:"Content"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRoutineRequest creates a request to invoke DeleteRoutine API
func CreateDeleteRoutineRequest() (request *DeleteRoutineRequest) {
	request = &DeleteRoutineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DeleteRoutine", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteRoutineResponse creates a response to parse from DeleteRoutine response
func CreateDeleteRoutineResponse() (response *DeleteRoutineResponse) {
	response = &DeleteRoutineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
