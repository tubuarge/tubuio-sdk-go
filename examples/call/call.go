package main

/*
call.go demonstrates a sample Contract Call.
to run:
	go run call.go
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	tubu "github.com/tubuarge/tubuio-sdk-go/api"
)

func main() {
	//create new apiStruct
	app := tubu.NewContract("YOUR-API-KEY")

	contract := app.CreateContract("ShortID")

	resp, err := contract.Call("addItem", "v2", "account", nil)
	if err != nil {
		panic(err)
	}

	//close response when there is no need.
	defer resp.Body.Close()

	//check response status code
	if resp.StatusCode != 200 {
		panic("NOT OK")
	}

	respMap := make(map[string]interface{})

	//convert response body to []byte
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//parse JSON response to a map
	err = json.Unmarshal(respBody, &respMap)
	if err != nil {
		panic(err)
	}

	//print "data" from parsed response
	fmt.Println(respMap["data"])

	//print "message" from parsed response
	fmt.Println(respMap["message"])
}
