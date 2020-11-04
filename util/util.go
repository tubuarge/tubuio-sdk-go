package util

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

//requestBody used for marshalling JSON data.
type requestBody struct {
	Args    []interface{} `json:"args"`
	Account string        `json:"account"`
}

//GetHttpGetUrl returns Http GET url for contract call.
func GetHttpGetUrl(baseUrl, shortId, method, tag, account string, args []interface{}) string {
	reqUrl := baseUrl + "/int/" + strings.TrimSpace(shortId) + "/" + strings.TrimSpace(method)
	//tag is not empty, include tag in the reqUrl
	if strings.Compare(tag, "") != 0 {
		reqUrl += "/" + strings.TrimSpace(tag)
	}

	params := url.Values{}

	//TODO: check if account is required.
	if strings.Compare(account, "") != 0 {
		reqUrl += "?"
		params.Add("account", account)
	}
	//args doesn't empty
	if len(args) > 0 {
		if strings.Compare(account, "") == 0 {
			reqUrl += "?"
		}
		for _, elem := range args {
			params.Add("args", fmt.Sprintf("%v", elem))
		}
	}
	return reqUrl + params.Encode()
}

//GetHttpPostUrl returns Http POST url for contract send.
func GetHttpPostUrl(baseUrl, shortId, method, tag string) string {
	url := baseUrl + "/int/" + strings.TrimSpace(shortId) + "/" + strings.TrimSpace(method)

	if strings.Compare(tag, "") != 0 {
		url += "/" + strings.TrimSpace(tag)
	}

	return url
}

//GetBodyRequest returns a JSON object that contains account and args.
func GetBodyRequest(account string, args []interface{}) ([]byte, error) {
	requestBodyStruct := requestBody{}
	if args != nil {
		//add args to struct slice
		for _, elem := range args {
			requestBodyStruct.Args = append(requestBodyStruct.Args, elem)
		}
	}

	requestBodyStruct.Account = account

	//parse struct to JSON
	requestBody, err := json.Marshal(requestBodyStruct)
	if err != nil {
		return nil, fmt.Errorf("error while creating request body: %v", err)
	}
	return requestBody, nil
}
