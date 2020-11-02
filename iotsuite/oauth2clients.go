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
)


type OAuthClientCreateRequest struct {
	ClientName string `json:"clientName"`
	TargetInstance string `json:"targetInstance"`
}

type OAuthClientDeleteRequest struct {
	ClientName string `json:"clientName"`
}

type OAuthClientListRequest struct {
//	Filter string `json:"filter"`
//	Page int`json:"page"`
}

type OAuthClientListResponse struct {
	OAuthClients []OAuthClient `json:"clients"`
}

type OAuthClient struct {
	ClientName string `json:"clientName"`
	ClientId string `json:"clientId"`
	Scopes string `json:"scopes"`
}

func ListOAuthClients(httpClient *http.Client) {
	var baseurl = viper.GetString("baseurl")
	var url = baseurl + "/api/v3/oauthclients";
	log.WithFields(log.Fields{"url": url}).Info("Using OAuth Client Management endpoint")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query() // Get a copy of the query values.
		req.URL.RawQuery = q.Encode() // Encode and assign back to the original query.

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println()

	// DumpJsonResponse(resp)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	var responseObject []Product
	jsonErr := json.Unmarshal([]byte(responseData), &responseObject)
	// err := json.Unmarshal([]byte(dataJson), &arr)
	if (jsonErr != nil) {
		log.Fatal(jsonErr);
		os.Exit(2);
	}

fmt.Printf("%-36s %-50s %s\n", "Product ID", "Product Name", "Service Description")
	for i := 0; i < len(responseObject); i++ {
		var product = responseObject[i]
		fmt.Printf("%-36s %-50s %s\n", product.Id, product.Name, product.ServiceDescription)
	}

}

func NewOAuthClient(httpClient *http.Client, product string, instanceName string) {
	var baseurl = viper.GetString("baseurl")
	var url = baseurl + "/api/v3/oauthclients";

	log.WithFields(log.Fields{"url": url}).Info("Using OAuth Management endpoint")

	body := &SubscriptionCreateRequest{
	    ProductId:    product,
	    InstanceName: instanceName,
	}
	bodyBuffer := new(bytes.Buffer)
	json.NewEncoder(bodyBuffer).Encode(body)

	req, err := http.NewRequest(http.MethodPost, url, bodyBuffer)
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
				"product":product,
				"instanceName":instanceName,
				"fullRequestObject":req,
				"err": err }).Fatal("Fatal error on HTTP request")
	}
	defer resp.Body.Close()

	fmt.Println()

	DumpJsonResponse(resp)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"req":req, "resp":resp, "err": err}).Fatal("Fatal error on reading HTTP response body")
		os.Exit(2)
	}

	var responseObject SubscriptionListResponse
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.Subscriptions); i++ {
		var sub = responseObject.Subscriptions[i]
		fmt.Printf("%-10s %32s %32s %32s\n", sub.Status, sub.ServiceInstanceName, sub.ServiceInstanceId, sub.SubscriptionId)
	}
}

func DeleteOAuthClient(httpClient *http.Client, clientName string) {
	var baseurl = viper.GetString("baseurl")
	var url = baseurl + "/api/v3/oauthclient";

	log.WithFields(log.Fields{"url": url}).Info("Using OAuth Management endpoint")

	body := &OAuthClientDeleteRequest{
	    ClientName: clientName,
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
				"clientName":clientName,
				"fullRequestObject":req,
				"err": err }).Fatal("Fatal error on HTTP request")
	}
	defer resp.Body.Close()

	fmt.Println()

	DumpJsonResponse(resp)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"req":req, "resp":resp, "err": err}).Fatal("Fatal error on reading HTTP response body")
		os.Exit(2)
	}

	var responseObject OAuthClientListResponse
	json.Unmarshal(responseData, &responseObject)

	// ...
}