<p align="left" style="margin: 10px 0 25px 0">
  <c href="https://github.com/tubuarge/tubuio-sdk-node">
    <img alt="tubu.io logo" src="https://raw.githubusercontent.com/tubuarge/tubuio-sdk-node/master/logo.png" width="200"/>
  </c>
</p>

Golang SDK for [tubu.io](https://www.tubu.io)

## Example Usage
```go
package main

import(
    "io/ioutil"
    "log"
    "fmt"
    
    tubu "github.com/tubuarge/tubuio-sdk-go/api"
) 

func main() {
    //create new api struct with your API Key.
    app := tubu.NewContract("YOUR API KEY")
    contract := app.CreateNewContract("ShortID")

    //make the integration request if there is no error when making the 
    //request callResp variable contains request response as *http.Response. 
    callResp, err := contract.Call("Method", "Tag", "", nil)
    if err != nil {
        log.Fatal(err)
    }
    
    defer callResp.Body.Close()
    
    callRespBody, err := ioutil.ReadAll(callResp.Body)
    if err != nil {
        log.Fatal(err)
    }

    //print the response
    fmt.Println(string(callRespBody))
    
    //make the send request if there is no error when making the request
    //callResp variable contains request response as *http.Response.
    sendResp, err := contract.Call("Method", "Tag", "Account Address", "item", 123, true)
    if err != nil {
        log.Fatal(err)
    }

    defer sendResp.Body.Close()

    sendRespBody, err := ioutil.ReadAll(sendResp.Body)
    if err != nil {
        log.Fatal(err)
    }
    
    //print the response body
    fmt.Println(string(sendRespBody))
}
```
More examples can be found at [examples](examples) folder.

## License

[MIT](LICENSE)