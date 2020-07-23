package iotsuite



import (
	"net/http"
	"log"
	"fmt"
)

func IotmgrRules(httpClient *http.Client, conf *Configuration) {
	var url = "https://manager.eu-1.bosch-iot-suite.com/api/1/mme/rules"
	Get(httpClient, url)
}

func IotmgrTasks(httpClient *http.Client, conf *Configuration) {
	var url = "https://manager.eu-1.bosch-iot-suite.com/api/1/mme/tasks"
	Get(httpClient, url)
}

func IotmgrDevices(httpClient *http.Client, filter string, option string, namespaces string, fields string) {
	fmt.Println("Filter:",filter)
	fmt.Println("Fields:",fields)
	fmt.Println("Option:",option)
	fmt.Println("Namespaces:",namespaces)
	
	var url = "https://manager.eu-1.bosch-iot-suite.com/api/1/di/devices"
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query() // Get a copy of the query values.

	if filter != "" {
		q.Add("filter", filter)
	}
	if fields != "" {
		q.Add("fields", fields)
	}
	if option != "" {
		q.Add("option", option)
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
	DumpJsonResponse(resp)
}

func IotmgrGroups(httpClient *http.Client, conf *Configuration) {
	var url = "https://manager.eu-1.bosch-iot-suite.com/api/1/di/groups/directories"
	Get(httpClient, url)
}