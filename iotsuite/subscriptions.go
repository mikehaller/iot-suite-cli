package iotsuite

import (
	"fmt"
//	"github.com/fatih/color"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
	"github.com/spf13/viper"

)


// STATUS.BOSCH-IOT-SUITE.COM JSON
type SubscriptionListResponse struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type Subscription struct {
	SubscriptionId string `json:"subscriptionId"`
	ServiceInstanceId string `json:"serviceInstanceId"`
	ServiceInstanceName string `json:"serviceInstanceName"`
	Status string `json:"status"`
	PlanName string `json:"planName"`
}


func SubscriptionsList(httpClient *http.Client) {
	var baseurl = viper.GetString("baseurl")
	var url = baseurl + "/api/v3/subscriptions";

	fmt.Printf("Using endpoint: %s\n", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query() // Get a copy of the query values.

//	if fields != "" {
//		q.Add("fields", fields) // Add a new value to the set.
//	}
//	if filter != "" {
//		q.Add("filter", filter)
//	}
//	if namespaces != "" {
//		q.Add("namespaces", namespaces)
//	}

	req.URL.RawQuery = q.Encode() // Encode and assign back to the original query.

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println()

	DumpJsonResponse(resp)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	var responseObject SubscriptionListResponse
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.Subscriptions); i++ {
		var sub = responseObject.Subscriptions[i]
		fmt.Printf("%-10s %32s %32s %32s\n", sub.Status, sub.ServiceInstanceName, sub.ServiceInstanceId, sub.SubscriptionId)
	}
}
