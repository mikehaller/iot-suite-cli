package iotsuite

import (
	"os"
	"fmt"
    "log"
    "net/url"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/TylerBrock/colorjson"
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
	ExpiresIn int `json:"expires_in"`
	Scope string `json:"scope"`
	TokenType string `json:"token_type"`
}


/*
grant_type=client_credentials
&client_id=e726c3f8-...
&client_secret=...
&scope=service:iot-hub-prod:t0262c358aab544399b78b4811bfd862b_hub/full-access%20service:iot-manager:0262c358-aab5-4439-9b78-b4811bfd862b_iot-manager/full-access%20service:iot-rollouts:0262c358-aab5-4439-9b78-b4811bfd862b_rollouts/full-access%20service:iot-things-eu-1:0262c358-aab5-4439-9b78-b4811bfd862b_things/full-access

*/
func Authorize(conf *Configuration) string {
	
	var clientId = conf.ClientId
	var clientSecret = conf.ClientSecret
	var scope = conf.Scope
	
	if clientId == "" {
		fmt.Println(Fatal("You need to specify the OAuth2 Client ID with -clientId"))
		fmt.Println("\nTo read about the command line options, use '" +os.Args[0] + "auth -h'",);
		fmt.Println("See https://accounts.bosch-iot-suite.com/oauth2-clients/");
		os.Exit(2)
	}

	if clientSecret == "" {
		fmt.Println(Fatal("You need to specify the OAuth2 Client Secret with -clientSecret"))
		fmt.Println("\nTo read about the command line options, use '" +os.Args[0] + "auth -h'",);
		fmt.Println("See https://accounts.bosch-iot-suite.com/oauth2-clients/");
		os.Exit(2)
	}
	
	if scope == "" {
		fmt.Println(Fatal("You need to specify the OAuth2 Scope with -scope"))
		fmt.Println("\nTo read about the command line options, use '" +os.Args[0] + "auth -h'",);
		fmt.Println("See https://accounts.bosch-iot-suite.com/oauth2-clients/");
		os.Exit(2)
	}
	
	response, err := http.PostForm("https://access.bosch-iot-suite.com/token",
		url.Values{
				"grant_type" : { "client_credentials" },
				"client_id" : { clientId },
				"client_secret" : { clientSecret },
				"scope" : { scope }})
	
	if response.StatusCode != 200 {
		fmt.Println("HTTP Response:")
		fmt.Println(Fatal(response.Status))
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
        fmt.Print("Error: %s",err.Error())
        os.Exit(1)
    }
	responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
        os.Exit(2)
    }
    
    var responseObject OAuthToken
	json.Unmarshal(responseData, &responseObject)
	
	fmt.Println(Teal("Token Type:"), Warn(responseObject.TokenType))
	fmt.Println(Teal("Scope:"), Warn(responseObject.Scope))
	fmt.Println(Teal("Access Token:"), Warn(responseObject.AccessToken))
	fmt.Println()
	fmt.Println(Warn("Warning:"),"Access token will expire in", responseObject.ExpiresIn,"seconds.")
	
	return responseObject.AccessToken
}