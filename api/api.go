/*
Copyright 2020 TUBU ARGE
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package api

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/tubuarge/tubuio-sdk-go/util"
)

const (
	BaseUrl = "https://prodservice-dot-dynamic-sun-260208.appspot.com"
	Timeout = 10000 * time.Millisecond
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HttpClient
)

//ApiStruct constructor struct of the api instance.
type ApiStruct struct {
	ApiKey string
}

func init() {
	Client = &http.Client{
		Timeout: Timeout,
	}
}

//NewApiStruct creates new api struct.
func NewApiStruct(apiKey string) *ApiStruct {
	return &ApiStruct{ApiKey: apiKey}
}

//Call calls the given call method of the contract's given tag version with given args.
//returns response as http.Response pointer.
func (a *ApiStruct) Call(shortId, method, tag, account string, args ...interface{}) (*http.Response, error) {
	callUrl := util.GetHttpGetUrl(BaseUrl, shortId, method, tag, account, args)

	req, err := a.createCallRequest(callUrl)
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//Send calls the given send method of the contract's given tag version with given args.
//returns response as http.Response pointer.
func (a *ApiStruct) Send(shortId, method, tag, account string, args ...interface{}) (*http.Response, error) {
	sendUrl := util.GetHttpPostUrl(BaseUrl, shortId, method, tag)

	requestBody, err := util.GetBodyRequest(account, args)
	if err != nil {
		return nil, err
	}

	req, err := a.createSendRequest(sendUrl, requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//createCallRequest creates a "GET" http.Request for Contract Calls with the given callUrl.
func (a *ApiStruct) createCallRequest(callUrl string) (*http.Request, error) {
	//create get request
	req, err := http.NewRequest("GET", callUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating call request: %v", err)
	}

	//set request headers
	req.Header.Set("accept", "application/json")
	req.Header.Set("ApiKey", a.ApiKey)

	return req, nil
}

//createSendRequest creates a "POST" http.Request for Contract Sends with the given sendUrl
//and data.
func (a *ApiStruct) createSendRequest(sendUrl string, data []byte) (*http.Request, error) {
	//create post request
	req, err := http.NewRequest("POST", sendUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error while creating send request: %v", err)
	}

	//set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("ApiKey", a.ApiKey)

	return req, nil
}