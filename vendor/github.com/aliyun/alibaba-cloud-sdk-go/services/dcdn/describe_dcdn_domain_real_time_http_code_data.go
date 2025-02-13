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

// DescribeDcdnDomainRealTimeHttpCodeData invokes the dcdn.DescribeDcdnDomainRealTimeHttpCodeData API synchronously
func (client *Client) DescribeDcdnDomainRealTimeHttpCodeData(request *DescribeDcdnDomainRealTimeHttpCodeDataRequest) (response *DescribeDcdnDomainRealTimeHttpCodeDataResponse, err error) {
	response = CreateDescribeDcdnDomainRealTimeHttpCodeDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainRealTimeHttpCodeDataWithChan invokes the dcdn.DescribeDcdnDomainRealTimeHttpCodeData API asynchronously
func (client *Client) DescribeDcdnDomainRealTimeHttpCodeDataWithChan(request *DescribeDcdnDomainRealTimeHttpCodeDataRequest) (<-chan *DescribeDcdnDomainRealTimeHttpCodeDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainRealTimeHttpCodeDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainRealTimeHttpCodeData(request)
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

// DescribeDcdnDomainRealTimeHttpCodeDataWithCallback invokes the dcdn.DescribeDcdnDomainRealTimeHttpCodeData API asynchronously
func (client *Client) DescribeDcdnDomainRealTimeHttpCodeDataWithCallback(request *DescribeDcdnDomainRealTimeHttpCodeDataRequest, callback func(response *DescribeDcdnDomainRealTimeHttpCodeDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainRealTimeHttpCodeDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainRealTimeHttpCodeData(request)
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

// DescribeDcdnDomainRealTimeHttpCodeDataRequest is the request struct for api DescribeDcdnDomainRealTimeHttpCodeData
type DescribeDcdnDomainRealTimeHttpCodeDataRequest struct {
	*requests.RpcRequest
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
}

// DescribeDcdnDomainRealTimeHttpCodeDataResponse is the response struct for api DescribeDcdnDomainRealTimeHttpCodeData
type DescribeDcdnDomainRealTimeHttpCodeDataResponse struct {
	*responses.BaseResponse
	EndTime              string               `json:"EndTime" xml:"EndTime"`
	StartTime            string               `json:"StartTime" xml:"StartTime"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	DomainName           string               `json:"DomainName" xml:"DomainName"`
	DataInterval         string               `json:"DataInterval" xml:"DataInterval"`
	RealTimeHttpCodeData RealTimeHttpCodeData `json:"RealTimeHttpCodeData" xml:"RealTimeHttpCodeData"`
}

// CreateDescribeDcdnDomainRealTimeHttpCodeDataRequest creates a request to invoke DescribeDcdnDomainRealTimeHttpCodeData API
func CreateDescribeDcdnDomainRealTimeHttpCodeDataRequest() (request *DescribeDcdnDomainRealTimeHttpCodeDataRequest) {
	request = &DescribeDcdnDomainRealTimeHttpCodeDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainRealTimeHttpCodeData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainRealTimeHttpCodeDataResponse creates a response to parse from DescribeDcdnDomainRealTimeHttpCodeData response
func CreateDescribeDcdnDomainRealTimeHttpCodeDataResponse() (response *DescribeDcdnDomainRealTimeHttpCodeDataResponse) {
	response = &DescribeDcdnDomainRealTimeHttpCodeDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
