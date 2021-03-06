package iotsuite

import (
	"encoding/json" // json rest api responses
	"github.com/TylerBrock/colorjson"
	"fmt"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"github.com/spf13/viper"
)

func dump(responseData []byte) {
	var output = viper.GetString("output");
	
	if (output != "") {
		fmt.Println("Response body written to file:", output)
		ioutil.WriteFile(output, responseData , 0644)
		return;
	}
	
	fmt.Println("Response:")
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
		log.Debug("HTTP Request Header:",req.Header)
	}
	if req.Body != nil {
		responseData, err3 := ioutil.ReadAll(req.Body)
		if err3 != nil {
			log.Fatal("Unable to read body octets:",err3)
			os.Exit(2)
		}
		dump(responseData)
	}
}

func DumpJsonResponse(resp *http.Response) {
	log.Debug("HTTP Response Status:", resp.Status)
	log.WithFields(log.Fields{"resp":resp}).Trace("HTTP Response Dump")
	if resp.Body != nil {
		responseData, err3 := ioutil.ReadAll(resp.Body)
		if err3 != nil {
			log.Fatal("DumpJsonResponse:",err3)
			os.Exit(2)
		}
		dump(responseData)
	} else {
		log.Debug("HTTP Response Body is nil")	
	}
}
