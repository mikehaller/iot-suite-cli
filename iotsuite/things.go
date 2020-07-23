package iotsuite

import (
	"log"
	"fmt"
	"net/http"
	"github.com/fatih/color"
)

func ThingsSolutions(httpClient *http.Client, solutionId string) {
	
	fmt.Printf("%v %v\n", color.BlueString("Solution ID:"), color.GreenString(solutionId))
	
	var url = "https://things.eu-1.bosch-iot-suite.com/api/2/solutions/" + solutionId
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query() // Get a copy of the query values.

	req.URL.RawQuery = q.Encode() // Encode and assign back to the original query.

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	DumpJsonResponse(resp)
}

func Things(httpClient *http.Client, fields string, filter string, namespaces string) {
	var url = "https://things.eu-1.bosch-iot-suite.com/api/2/search/things"
	thingssearch(httpClient,url,fields,filter,namespaces)
}

func ThingsCount(httpClient *http.Client, filter string, namespaces string) {
	var url = "https://things.eu-1.bosch-iot-suite.com/api/2/search/things/count"
	thingssearch(httpClient,url,"",filter,namespaces)
}

func thingssearch(httpClient *http.Client, url string, fields string, filter string, namespaces string) {
	fmt.Println("Filter:",filter)
	fmt.Println("Fields:",fields)
	fmt.Println("Namespaces:",namespaces)
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query() // Get a copy of the query values.

	if fields != "" {
		q.Add("fields", fields) // Add a new value to the set.
	}
	if filter != "" {
		q.Add("filter", filter)
	}
	if namespaces != "" {
		q.Add("namespaces", namespaces)
	}

	req.URL.RawQuery = q.Encode() // Encode and assign back to the original query.

	fmt.Println("Request:",req)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	DumpJsonResponse(resp)
}
