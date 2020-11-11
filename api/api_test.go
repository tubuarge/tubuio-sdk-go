package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/tubuarge/tubuio-sdk-go/util"
	"github.com/stretchr/testify/assert"
)

var (
	apiStruct *ApiStruct

	getDoFunc func(req *http.Request) (*http.Response, error)
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return getDoFunc(req)
}

func init() {
	Client = &MockClient{}
}

func TestContractSend(t *testing.T) {
	tables := []struct {
		apiKey  string
		shortId string
		method  string
		tag     string
		account string
		args    []interface{}
	}{
		{"Your-Api-Key", "de5baba74567442b", "addItem", "", "", []interface{}{"item", 123, true}},
		{"a877-352ca6df3bab9", "de5baba74567442b", "addItem", "", "", nil},
	}
	for _, table := range tables {
		apiStruct = NewApiStruct(table.apiKey)
		assert.NotNil(t, apiStruct)

		successRespJSON := `{"data": "", "message": "ok"}`
		r := ioutil.NopCloser(bytes.NewBuffer([]byte(successRespJSON)))
		getDoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r}, nil
		}
		testUrl := util.GetHttpPostUrl(BaseUrl, table.shortId, table.method, table.tag)

		bodyData, err := util.GetBodyRequest(table.account, table.args)
		assert.Nil(t, err)
		assert.NotNil(t, bodyData)

		request, err := apiStruct.createSendRequest(testUrl, bodyData)
		assert.Nil(t, err)
		assert.NotNil(t, request)

		resp, err := apiStruct.Send(table.shortId, table.method, table.tag,table.account, table.args)

		assert.NotNil(t, resp)
		assert.Nil(t, err)

		//check request headers
		assert.EqualValues(t, table.apiKey, request.Header.Get("ApiKey"))
		assert.EqualValues(t, "application/json", request.Header.Get("Content-Type"))
		assert.EqualValues(t, "application/json", request.Header.Get("accept"))

		//check request method
		assert.EqualValues(t, "POST", request.Method)

	}
}

func TestContractCall(t *testing.T) {
	tables := []struct {
		apiKey  string
		shortId string
		method  string
		tag     string
		account string
		args    []interface{}
	}{
		{"32eb414e-a046-4bf4-a877-352ca6d3bab9", "de5baba74567442b", "getItems", "", "", nil},
		{"a877-352ca6df3bab9", "de5baba74567442b", "getLastItem", "", "", nil},
	}

	for _, table := range tables {
		apiStruct = NewApiStruct(table.apiKey)
		assert.NotNil(t, apiStruct)

		successRespJSON := `{"data": "", "message": "ok"}`
		r := ioutil.NopCloser(bytes.NewBuffer([]byte(successRespJSON)))
		getDoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r}, nil
		}

		testUrl := util.GetHttpGetUrl(BaseUrl, table.shortId, "getItems", "", "", nil)

		request, err := apiStruct.createCallRequest(testUrl)
		assert.Nil(t, err)
		assert.NotNil(t, request)

		resp, err := apiStruct.Call(table.shortId, table.method, table.tag,table.account, table.args)

		assert.NotNil(t, resp)
		assert.Nil(t, err)

		//check request headers
		assert.EqualValues(t, table.apiKey, request.Header.Get("ApiKey"))
		assert.EqualValues(t, "application/json", request.Header.Get("accept"))

		//check request method
		assert.EqualValues(t, "GET", request.Method)
	}
}
