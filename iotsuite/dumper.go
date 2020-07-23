package iotsuite

import (
	"encoding/json" // json rest api responses
	"fmt"
	"github.com/TylerBrock/colorjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func dump(responseData []byte) {
	var obj map[string]interface{}
	json.Unmarshal([]byte(responseData), &obj)
	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4
	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}

func DumpJsonRequest(req *http.Request) {
	if req.Header != nil {
		fmt.Println(req.Header)
	}
	if req.Body != nil {
		responseData, err3 := ioutil.ReadAll(req.Body)
		if err3 != nil {
			log.Fatal(err3)
			os.Exit(2)
		}
		dump(responseData)
	}
}

func DumpJsonResponse(resp *http.Response) {
	fmt.Println("HTTP Response:", resp.Status)
	if resp.Body != nil {
		responseData, err3 := ioutil.ReadAll(resp.Body)
		if err3 != nil {
			log.Fatal(err3)
			os.Exit(2)
		}
		dump(responseData)
	}
}
