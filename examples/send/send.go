package main

/*
send.go demonstrates a sample Contract Send.
to run:
	go run send.go
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
	contract := app.CreateContract("shortID")

	resp, err :=contract.Send("addItem", "", "", "item", 123, true)
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
