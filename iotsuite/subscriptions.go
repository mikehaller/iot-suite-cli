package iotsuite

import (
	"fmt"
//	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"os"
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"net/http/httptrace"
	"time"
)


type SubscriptionCreateRequest struct {
	ProductId string `json:"productId"`
	InstanceName string `json:"instanceName"`
}

type SubscriptionCancelRequest struct {
	InstanceName string `json:"instanceName"`
}


type SubscriptionListRequest struct {
//	Filter string `json:"filter"`
//	Page int`json:"page"`
}

type Subscription struct {
	ServiceInstanceName string `json:"instanceName"`
	FreePlan bool `json:"freePlan"`
	OrderDate time.Time`json:"orderDate"`
	OrgId string `json:"organizationId"`
	OrgName string `json:"organizationName"`
	PlanDescription string `json:"planDescription"`
	PlanName string `json:"planName"`
	Platform string `json:"platform"`
	ProductId string `json:"productId"`
	ProductName string `json:"productName"`
	ProductRegion string `json:"productRegion"`
	ProvisioningDate time.Time `json:"provisioningDate"`
	Status string `json:"status"`
	SubscriptionId string `json:"subscriptionId"`
	ServiceInstances map[string]interface{} `json:"serviceInstances"`
}

/*

"freePlan": true,
        "instanceName": "DONT DELETE. AC dummy paid. Allows to switch plan of AC subscriptions.",
        "orderDate": "2020-03-04T12:34:27Z",
        "organizationId": "a058d238-0902-4050-9056-7e27b6710711",
        "organizationName": "IOC/PAP-SP-Testorg",
        "planDescription": "Free Plan",
        "planName": "Free Plan",
        "platform": "AWS",
        "productId": "7f430995-e293-49ac-b485-3ae736113fb4",
        "productName": "com.bosch.iot.suite.telemetry - free - eu-1",
        "productRegion": "EU_1",
        "provisioningDate": "2020-03-04T12:34:27Z",
        "serviceInstances": {
            "iot-hub": "314f6c35-0e0e-4f64-9d3c-ac79552d3f97",
            "iot-things": "314f6c35-0e0e-4f64-9d3c-ac79552d3f97"
        },
        "status": "Active",
        "subscriptionId": "adc2add0-6def-436d-9d60-cf3b9b7819d6",
        "terminationDate": ""
*/

/**
"cloudServiceDescription": "A simple example",
        "cloudServiceId": "example-service",
        "cloudServiceName": "example",
        "currency": "EUR",
        "dataCenter": "eu-1",
        "deleted": false,
        "evalUsersCanBook": true,
        "freeplan": true,
        "monthlyPrice": 0,
        "planEnabled": true,
        "planId": "free-plan",
        "planName": "free",
        "productId": "11111111-e936-458b-b9ac-16f8ce19e32c",
        "productListingName": "example-service - free-plan - eu-1",
        "productOfficialName": "Example Service"
*/
type Product struct {
	Id string `json:"productId"`
	ProductName string `json:"productName"`
	ProductVariant string `json:"productVariant"`
	PlanName string `json:"planName"`
	ProductRegion string `json:"productRegion"`
	FreePlan bool `json:"freePlan"`
}

func ProductsList(httpClient *http.Client) {
	var baseurl = viper.GetString("baseurl")
	var url = baseurl + "/api/v3/products";
	log.WithFields(log.Fields{"url": url}).Info("Using Subscription Management endpoint")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query() // Get a copy of the query values.
	q.Add("orgId", viper.GetString("orgId"))
	req.URL.RawQuery = q.Encode() // Encode and assign back to the original query.

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("ProductsList HTTP Request generally failed:",err)
	}
	defer resp.Body.Close()

	fmt.Println()

	DumpJsonResponse(resp)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Reading response body for ProductsList failed:",err)
		os.Exit(2)
	}

	log.Debug("Body:", string(responseData));

	var responseObject []Product
	jsonErr := json.Unmarshal([]byte(responseData), &responseObject)
	if (jsonErr != nil) {
		log.Fatal("Unmarshalling response body for ProductsList failed:",jsonErr);
		os.Exit(2);
	}

	fmt.Printf("%-38q %-50q %-20q %-7q %q\n", "Product ID", "Product Name", "Plan Name","Region","Free?")
	for i := 0; i < len(responseObject); i++ {
		var product = responseObject[i]
		fmt.Printf("%-38q %-50q %-20q %-7q %t\n", product.Id, product.ProductName, product.PlanName, product.ProductRegion, product.FreePlan)
	}

}

func SubscriptionsList(httpClient *http.Client) {
	var baseurl = viper.GetString("baseurl")
	var url = baseurl + "/api/v3/subscriptions";

	log.WithFields(log.Fields{"url": url}).Info("Using Subscription Management endpoint")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query() // Get a copy of the query values.
	q.Add("orgId", viper.GetString("orgId"))
	req.URL.RawQuery = q.Encode() // Encode and assign back to the original query.

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// fmt.Println()
	// DumpJsonResponse(resp)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	var responseObject []Subscription
	jsonErr := json.Unmarshal([]byte(responseData), &responseObject)
	if (jsonErr != nil) {
		log.Fatal(jsonErr);
		os.Exit(2);
	}

	for i := 0; i < len(responseObject); i++ {
		var sub = responseObject[i]
		fmt.Printf("%-38s %-12s %-20s %s %s %-20s\n", sub.SubscriptionId, sub.Status, sub.ServiceInstanceName, sub.ProductName, sub.OrderDate.Format(time.RFC3339), sub.ServiceInstances)
	}
}


func NewSubscription(httpClient *http.Client, product string, instanceName string) {
	var baseurl = viper.GetString("baseurl")
	var url = baseurl + "/api/v3/subscriptions";

	log.WithFields(log.Fields{"url": url}).Info("Using Subscription Management endpoint")

	body := &SubscriptionCreateRequest{
	    ProductId:    product,
	    InstanceName: instanceName,
	}
	
	bodyBuffer := new(bytes.Buffer);
	json.NewEncoder(bodyBuffer).Encode(body)

	req, err := http.NewRequest(http.MethodPost, url, bodyBuffer)
	req.Header.Set("Content-Type", "application/json")
	
	q := req.URL.Query() // Get a copy of the query values.
	q.Add("orgId", viper.GetString("orgId"))
	req.URL.RawQuery = q.Encode() // Encode and assign back to the original query.
	
    req = req.WithContext(httptrace.WithClientTrace(req.Context(), newHttpTrace()))
    if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
    		log.WithFields(log.Fields{"err": err}).Fatal("Fatal error on HTTP roundtrip")
    }

		log.WithFields(log.Fields{
				"body":bodyBuffer.String(),
				"url":url,
				"product":product,
				"instanceName":instanceName,
				"fullRequestObject":req,
				"err": err }).Info("Info")


	// Execute HTTP Request
	resp, err := httpClient.Do(req)
	
	if err != nil {
		log.WithFields(log.Fields{
				"baseurl":baseurl,
				"requestMethod":req.Method,
				"url":req.URL,
				"httpProtocol":req.Proto,
				"header":req.Header,
				"contentLength":req.ContentLength,
				"transferEncoding":req.TransferEncoding,
				"closeFlag":req.Close,
				"host":req.Host,
				"form":req.Form,
				"postForm":req.PostForm,
				"multipartForm":req.MultipartForm,
				"trailer":req.Trailer,
				"remoteAddr":req.RemoteAddr,
				"requestURI":req.RequestURI,
				"product":product,
				"instanceName":instanceName,
				"fullRequestObject":req,
				"err": err }).Fatal("Fatal error on HTTP request")
	}
	defer resp.Body.Close()

	fmt.Println()

	DumpJsonResponse(resp)

}

func CancelSubscription(httpClient *http.Client, instanceName string) {
	var baseurl = viper.GetString("baseurl")
	var url = baseurl + "/api/v3/subscriptions";

	log.WithFields(log.Fields{"url": url}).Info("Using Subscription Management endpoint")

	body := &SubscriptionCancelRequest{
	    InstanceName: instanceName,
	}
	bodyBuffer := new(bytes.Buffer)
	json.NewEncoder(bodyBuffer).Encode(body)

	req, err := http.NewRequest(http.MethodDelete, url, bodyBuffer)
	req.Header.Set("Content-Type", "application/json")
    req = req.WithContext(httptrace.WithClientTrace(req.Context(), newHttpTrace()))
    if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
    		log.WithFields(log.Fields{"err": err}).Fatal("Fatal error on HTTP roundtrip")
    }

	// Execute HTTP Request
	resp, err := httpClient.Do(req)
	
	if err != nil {
		log.WithFields(log.Fields{
				"baseurl":baseurl,
				"requestMethod":req.Method,
				"url":req.URL,
				"httpProtocol":req.Proto,
				"header":req.Header,
				"contentLength":req.ContentLength,
				"transferEncoding":req.TransferEncoding,
				"closeFlag":req.Close,
				"host":req.Host,
				"form":req.Form,
				"postForm":req.PostForm,
				"multipartForm":req.MultipartForm,
				"trailer":req.Trailer,
				"remoteAddr":req.RemoteAddr,
				"requestURI":req.RequestURI,
				"instanceName":instanceName,
				"fullRequestObject":req,
				"err": err }).Fatal("Fatal error on HTTP request")
	}
	defer resp.Body.Close()

	fmt.Println()

	DumpJsonResponse(resp)

}