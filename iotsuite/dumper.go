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
	var err = json.Unmarshal([]byte(responseData), &obj)
	if err != nil {
		var decoded []interface{}
		err = json.Unmarshal([]byte(responseData), &decoded)
		if err != nil {
			// Root is neither Object nor Array, assume primitive value.
			fmt.Println(string(responseData))
		} else {
			// Make a custom formatter with indent set
			f := colorjson.NewFormatter()
			f.Indent = 4
			// Marshall the Colorized JSON
			s, err := f.Marshal(decoded)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(s))
		}
	} else {
		// Make a custom formatter with indent set
		f := colorjson.NewFormatter()
		f.Indent = 4
		// Marshall the Colorized JSON
		s, err := f.Marshal(obj)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(s))
	}
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
