package iotsuite

import (
	"encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
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

/*
grant_type=client_credentials
&client_id=e726c3f8-...
&client_secret=...
&scope=service:iot-hub-prod:t0262c358aab544399b78b4811bfd862b_hub/full-access%20service:iot-manager:0262c358-aab5-4439-9b78-b4811bfd862b_iot-manager/full-access%20service:iot-rollouts:0262c358-aab5-4439-9b78-b4811bfd862b_rollouts/full-access%20service:iot-things-eu-1:0262c358-aab5-4439-9b78-b4811bfd862b_things/full-access

*/

func err() {
	fmt.Println("You need to specify all three OAuth2 client parameters: clientId, clientSecret and scope")
	fmt.Println("See https://accounts.bosch-iot-suite.com/oauth2-clients/")
	os.Exit(2)
}

func Authorize(conf *Configuration) string {

	var existingToken = LoadToken()
	if existingToken.AccessToken != "" {
		fmt.Printf("%v\n", color.YellowString("Warning: There is an existing token in the disk cache."))
		fmt.Printf("%v\n", color.YellowString("Warning: Access token is being refreshed and updated in disk cache."))
	}

	var clientId = conf.ClientId
	var clientSecret = conf.ClientSecret
	var scope = conf.Scope

	if clientId == "" || clientSecret == "" || scope == "" {
		err()
	}

	response, err := http.PostForm("https://access.bosch-iot-suite.com/token",
		url.Values{
			"grant_type":    {"client_credentials"},
			"client_id":     {clientId},
			"client_secret": {clientSecret},
			"scope":         {scope}})

	if response.StatusCode != 200 {
		fmt.Println("HTTP Response:")
		fmt.Println(response.Status)
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
		fmt.Println(string(s))

		os.Exit(3)
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

	fmt.Printf("%v %v\n", color.BlueString("Token Type:"), color.GreenString(responseObject.TokenType))
	fmt.Printf("\n%v %v\n", color.BlueString("Scope:"), color.GreenString(responseObject.Scope))
	fmt.Printf("\n%v\n%v\n", color.BlueString("Access Token:"), color.GreenString(responseObject.AccessToken))

	fmt.Println()

	fmt.Printf("%v %v %v\n", color.YellowString("Warning: Access token will expire in"), color.RedString(strconv.Itoa(responseObject.ExpiresIn)), color.YellowString("seconds."))

	StoreToken(&responseObject)

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
