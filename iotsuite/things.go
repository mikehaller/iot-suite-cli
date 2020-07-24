package iotsuite

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"net/http"
)

func ThingsConnections(conf *Configuration, httpClient *http.Client, solutionId string) {
	var url = "https://things.eu-1.bosch-iot-suite.com/api/2/solutions/" + solutionId + "/connections"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	defer resp.Body.Close()
}

func simpleget(httpClient *http.Client, url string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	DumpJsonResponse(resp)
}

func ThingsSolution(conf *Configuration, httpClient *http.Client, solutionId string) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Fprintf(color.Output, "%v %v\n", blue("Solution ID:"), color.GreenString(solutionId))
	color.Unset() // Don't forget to unset

	var url = "https://things.eu-1.bosch-iot-suite.com/api/2/solutions/" + solutionId
	req, err := http.NewRequest(http.MethodGet, url, nil)
	fmt.Println(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	DumpJsonResponse(resp)
}

func Things(httpClient *http.Client, fields string, filter string, namespaces string) {
	var url = "https://things.eu-1.bosch-iot-suite.com/api/2/search/things"
	thingssearch(httpClient, url, fields, filter, namespaces)
}

func ThingsCount(httpClient *http.Client, filter string, namespaces string) {
	var url = "https://things.eu-1.bosch-iot-suite.com/api/2/search/things/count"
	thingssearch(httpClient, url, "", filter, namespaces)
}

func thingssearch(httpClient *http.Client, url string, fields string, filter string, namespaces string) {

	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	if filter != "" {
		fmt.Fprintf(color.Output, "%s %s\n", blue("Filter:"), green(filter))
	}
	if fields != "" {
		fmt.Fprintf(color.Output, "%s %s\n", blue("Fields:"), green(fields))
	}
	if namespaces != "" {
		fmt.Fprintf(color.Output, "%s %s\n", blue("Namespaces:"), green(namespaces))
	}

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

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println()

	DumpJsonResponse(resp)
}
