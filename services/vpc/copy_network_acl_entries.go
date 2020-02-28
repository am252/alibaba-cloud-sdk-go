package vpc

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

// CopyNetworkAclEntries invokes the vpc.CopyNetworkAclEntries API synchronously
// api document: https://help.aliyun.com/api/vpc/copynetworkaclentries.html
func (client *Client) CopyNetworkAclEntries(request *CopyNetworkAclEntriesRequest) (response *CopyNetworkAclEntriesResponse, err error) {
	response = CreateCopyNetworkAclEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// CopyNetworkAclEntriesWithChan invokes the vpc.CopyNetworkAclEntries API asynchronously
// api document: https://help.aliyun.com/api/vpc/copynetworkaclentries.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyNetworkAclEntriesWithChan(request *CopyNetworkAclEntriesRequest) (<-chan *CopyNetworkAclEntriesResponse, <-chan error) {
	responseChan := make(chan *CopyNetworkAclEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyNetworkAclEntries(request)
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

// CopyNetworkAclEntriesWithCallback invokes the vpc.CopyNetworkAclEntries API asynchronously
// api document: https://help.aliyun.com/api/vpc/copynetworkaclentries.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyNetworkAclEntriesWithCallback(request *CopyNetworkAclEntriesRequest, callback func(response *CopyNetworkAclEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyNetworkAclEntriesResponse
		var err error
		defer close(result)
		response, err = client.CopyNetworkAclEntries(request)
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

// CopyNetworkAclEntriesRequest is the request struct for api CopyNetworkAclEntries
type CopyNetworkAclEntriesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	NetworkAclId         string           `position:"Query" name:"NetworkAclId"`
	SourceNetworkAclId   string           `position:"Query" name:"SourceNetworkAclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CopyNetworkAclEntriesResponse is the response struct for api CopyNetworkAclEntries
type CopyNetworkAclEntriesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCopyNetworkAclEntriesRequest creates a request to invoke CopyNetworkAclEntries API
func CreateCopyNetworkAclEntriesRequest() (request *CopyNetworkAclEntriesRequest) {
	request = &CopyNetworkAclEntriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CopyNetworkAclEntries", "Vpc", "openAPI")
	return
}

// CreateCopyNetworkAclEntriesResponse creates a response to parse from CopyNetworkAclEntries response
func CreateCopyNetworkAclEntriesResponse() (response *CopyNetworkAclEntriesResponse) {
	response = &CopyNetworkAclEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
