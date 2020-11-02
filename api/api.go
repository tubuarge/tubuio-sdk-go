/*
Copyright 2020 TUBU ARGE
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package api

import (
	"../util"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BASE_URL = "https://prodservice-dot-dynamic-sun-260208.appspot.com"
	TIMEOUT  = 10000
)

//ApiStruct constructor struct of the api instance.
type ApiStruct struct {
	ApiKey string
}

//NewApiStruct creates new api struct.
func NewApiStruct(apiKey string) *ApiStruct {
	return &ApiStruct{ApiKey: apiKey}
}

func (a *ApiStruct) IntegrationCall(shortId, method, tag, account string, args []string) ([]byte, error) {
	url := util.GetHttpGetUrl(BASE_URL, shortId, method, tag, account, args)

	//TODO: check args type.
	resp, err := a.doGet(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApiStruct) IntegrationSend(shortId, method, tag string, args []string, account string) ([]byte, error) {
	url := util.GetHttpPostUrl(BASE_URL, shortId, method, tag)

	requestBody, err := util.GetBodyRequest(account, args)

	resp, err := a.doPost(url, requestBody)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//doGet handles get request.
func (a *ApiStruct) doGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body to byte slice: %v", err)
	}
	return body, nil
}

//doPost handles post request.
func (a *ApiStruct) doPost(url string, data []byte) ([]byte, error) {
	//create client
	client := http.Client{
		Timeout: TIMEOUT,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error while creating request: %v", err)
	}

	//set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("ApiKey", a.ApiKey)

	//do request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body to byte slice: %v", err)
	}

	return body, nil
}
