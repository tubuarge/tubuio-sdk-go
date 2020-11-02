package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

//requestBody used for marshalling JSON data.
type requestBody struct {
	args    []string `json:"args"`
	account string   `json:"account"`
}

//GetHttpGetUrl returns Http GET url for integration call.
func GetHttpGetUrl(baseUrl, shortId, method, tag, account string, args []string) string {
	url := baseUrl + "/int/" + strings.TrimSpace(shortId) + "/" + strings.TrimSpace(method)
	//tag is not empty, include tag in the url
	if strings.Compare(tag, "") != 0 {
		url += "/" + strings.TrimSpace(tag)
	}

	//TODO: check if account is required.
	if strings.Compare(account, "") != 0 {
		url += "?account=" + account
	}

	//args doesn't empty
	if len(args) > 0 {
		for _, elem := range args {
			//TODO: check if args could take any integer value.
			url += "&args=" + elem
		}
	}
	return url
}

//GetHttpPostUrl returns Http POST url for integration send.
func GetHttpPostUrl(baseUrl, shortId, method, tag string) string {
	url := baseUrl + "/int" + strings.TrimSpace(shortId) + "/" + strings.TrimSpace(method)

	if strings.Compare(tag, "") != 0 {
		url += "/" + strings.TrimSpace(tag)
	}

	return url
}

func GetBodyRequest(account string, args []string) ([]byte, error) {
	//init requestBody struct
	requestBodyStruct := requestBody{
		args:    args,
		account: account,
	}

	//TODO: is account & args is required.
	//get request body in json
	requestBody, err := json.Marshal(requestBodyStruct)
	if err != nil {
		return nil, fmt.Errorf("error while creating request body: %v", err)
	}
	return requestBody, nil
}
