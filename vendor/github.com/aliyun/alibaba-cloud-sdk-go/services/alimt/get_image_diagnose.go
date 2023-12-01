package alimt

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

// GetImageDiagnose invokes the alimt.GetImageDiagnose API synchronously
func (client *Client) GetImageDiagnose(request *GetImageDiagnoseRequest) (response *GetImageDiagnoseResponse, err error) {
	response = CreateGetImageDiagnoseResponse()
	err = client.DoAction(request, response)
	return
}

// GetImageDiagnoseWithChan invokes the alimt.GetImageDiagnose API asynchronously
func (client *Client) GetImageDiagnoseWithChan(request *GetImageDiagnoseRequest) (<-chan *GetImageDiagnoseResponse, <-chan error) {
	responseChan := make(chan *GetImageDiagnoseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetImageDiagnose(request)
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

// GetImageDiagnoseWithCallback invokes the alimt.GetImageDiagnose API asynchronously
func (client *Client) GetImageDiagnoseWithCallback(request *GetImageDiagnoseRequest, callback func(response *GetImageDiagnoseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetImageDiagnoseResponse
		var err error
		defer close(result)
		response, err = client.GetImageDiagnose(request)
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

// GetImageDiagnoseRequest is the request struct for api GetImageDiagnose
type GetImageDiagnoseRequest struct {
	*requests.RpcRequest
	Url   string `position:"Body" name:"Url"`
	Extra string `position:"Body" name:"Extra"`
}

// GetImageDiagnoseResponse is the response struct for api GetImageDiagnose
type GetImageDiagnoseResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetImageDiagnoseRequest creates a request to invoke GetImageDiagnose API
func CreateGetImageDiagnoseRequest() (request *GetImageDiagnoseRequest) {
	request = &GetImageDiagnoseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alimt", "2018-10-12", "GetImageDiagnose", "", "")
	request.Method = requests.POST
	return
}

// CreateGetImageDiagnoseResponse creates a response to parse from GetImageDiagnose response
func CreateGetImageDiagnoseResponse() (response *GetImageDiagnoseResponse) {
	response = &GetImageDiagnoseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
