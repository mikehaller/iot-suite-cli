package iotsuite

import (
	"encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"github.com/spf13/viper"
)

// OAuth Token Response
/*
{
 "access_token":"ey...",
 "expires_in":3599,
 "scope":"service:iot-hub-prod:t0262c358aab544399b78b4811bfd862b_hub/full-access service:iot-manager:0262c358-aab5-4439-9b78-b4811bfd862b_iot-manager/full-access service:iot-rollouts:0262c358-aab5-4439-9b78-b4811bfd862b_rollouts/full-access service:iot-things-eu-1:0262c358-aab5-4439-9b78-b4811bfd862b_things/full-access",
 "token_type":"bearer"
}*/
type OAuthToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func Authorize(conf *Configuration) string {

	var existingToken = LoadToken()
	if existingToken.AccessToken != "" {
		log.Info("Existing token in disk cache will be refreshed (accesstoken.json)")
	}

	var clientId = conf.ClientId
	var clientSecret = conf.ClientSecret
	var scope = conf.Scope

	var tokenurl = viper.GetString("tokenurl")
	log.WithFields(log.Fields{"tokenurl":tokenurl}).Debug("Using authorization server token endpoint");
	
	var data = url.Values{
			"grant_type":    {"client_credentials"},
			"client_id":     {clientId},
			"client_secret": {clientSecret},
			"scope":         {scope}}
	
	log.WithFields(log.Fields{"data":data}).Trace("HTTP Request Parameters")
	
	response, err := http.PostForm(tokenurl,data)

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	log.Println(response.StatusCode)

	if response.StatusCode != 200 {
		
		log.WithFields(log.Fields{"status":response.Status}).Debug("HTTP Response Status")
		log.WithFields(log.Fields{"header":response.Header}).Trace("HTTP Response Headers")
		
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}
		
		// https://stackoverflow.com/questions/19038598/how-can-i-pretty-print-json-using-go
		var obj map[string]interface{}
		json.Unmarshal([]byte(responseData), &obj)
		// Make a custom formatter with indent set
		f := colorjson.NewFormatter()
		f.Indent = 4
		// Marshall the Colorized JSON
		s, _ := f.Marshal(obj)
		log.Fatal("Unable to authorize at ", tokenurl, "\n", string(s))

		os.Exit(3)
	} else {
		fmt.Println("#authorized ok")
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
	
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	var responseObject OAuthToken
	json.Unmarshal(responseData, &responseObject)

	fmt.Printf("%v %v\n", "Token Type:", responseObject.TokenType)
	fmt.Printf("%v %v\n", "Scope:", responseObject.Scope)

	StoreToken(&responseObject)

	var output = viper.GetString("output");
	if (output != "") {
		fmt.Println("Access token written to file in JWT format:", output)
		ioutil.WriteFile(output, responseData , 0644)
	} else {
	}

	log.WithFields(log.Fields{	"accessToken":responseObject.AccessToken,
								"tokenType":responseObject.TokenType,
								"scope":responseObject.Scope,
								"expire":responseObject.ExpiresIn}).Trace("Access Token Response")
	
	var expirySeconds = strconv.Itoa(responseObject.ExpiresIn);
	log.WithFields(log.Fields{"expiration":expirySeconds}).Info("Access token will expire")

	return responseObject.AccessToken
}

func configPath() string {
	return "accesstoken.json"
}

func StoreToken(tokenObject *OAuthToken) {
	jsonC, _ := json.Marshal(tokenObject)
	ioutil.WriteFile(configPath(), jsonC, os.ModeAppend)
}

func LoadToken() OAuthToken {
    data, _ := ioutil.ReadFile(configPath())
    var token OAuthToken
    json.Unmarshal(data, &token)
    return token
}
