package auth

import (
	"os"
	"fmt"
    "log"
    "net/url"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "bosch-iot-suite/utils"
    "bosch-iot-suite/config"
    "github.com/TylerBrock/colorjson"
)

// OAuth Token Response
/*
{
 "access_token":"eyJhbGciOiJSUzI1NiIsImtpZCI6InB1YmxpYzo5YzI0MTg0OC01MmQyLTRkM2YtYmFmOS1mNzE3OTYxMmE3YWMiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOltdLCJjbGllbnRfaWQiOiJlNzI2YzNmOC03NjJlLTRhNjUtYjMwZC02ZWZlM2FjNTYwNzUiLCJleHAiOjE1OTUwMzE5MzYsImV4dCI6e30sImlhdCI6MTU5NTAyODMzNiwiaXNzIjoiaHR0cHM6Ly9hY2Nlc3MuYm9zY2gtaW90LXN1aXRlLmNvbS92Mi8iLCJqdGkiOiI5YWEyZjgxZS1iZDQ1LTQ5YTItYWE2Ny0zYmIzZTRjYWZkZTUiLCJuYmYiOjE1OTUwMjgzMzYsInNjcCI6WyJzZXJ2aWNlOmlvdC1odWItcHJvZDp0MDI2MmMzNThhYWI1NDQzOTliNzhiNDgxMWJmZDg2MmJfaHViL2Z1bGwtYWNjZXNzIiwic2VydmljZTppb3QtbWFuYWdlcjowMjYyYzM1OC1hYWI1LTQ0MzktOWI3OC1iNDgxMWJmZDg2MmJfaW90LW1hbmFnZXIvZnVsbC1hY2Nlc3MiLCJzZXJ2aWNlOmlvdC1yb2xsb3V0czowMjYyYzM1OC1hYWI1LTQ0MzktOWI3OC1iNDgxMWJmZDg2MmJfcm9sbG91dHMvZnVsbC1hY2Nlc3MiLCJzZXJ2aWNlOmlvdC10aGluZ3MtZXUtMTowMjYyYzM1OC1hYWI1LTQ0MzktOWI3OC1iNDgxMWJmZDg2MmJfdGhpbmdzL2Z1bGwtYWNjZXNzIl0sInN1YiI6ImU3MjZjM2Y4LTc2MmUtNGE2NS1iMzBkLTZlZmUzYWM1NjA3NSJ9.fsuQtxAJ_pzfePb56jHNwfWg_MZdl8tSJw3ou3gGphCRSpnOfaE9yDc8WbiYr3LD1NxbF0MvDKK-FlVcr41b2N4C7J9cBLQuCeXrV3KwNfEYmepy7A0s8iDkQmXs0NnS5ZAtvJYyGAH95VWfSQm3FTrn3NGeNSATXs4AUyNtWxxKr1hzQ2WYs37NtaL6mhqnyqa4_lILqwLg0q2i7SNaDAX5Lu6d7EZylL6dtFjvhNUJC0Ly26Vk9q3R9V72Ddu44lPIvUGs11TTflguzoi1Sz9fPYbAH3ZJY5EPUkttqRTYhfp93MuCLtm5CNoswM1RnseSZVYYGM5bsRiiWmYCut2GlW4RlwgaQts-aVPSUcyLVax3YdPtTf4l2z3R1kh08lVs5TEMWF0IgaatN5NRzv0MmBc4_GoKAbPRfoisaVd2x7D7nhKC2jhBQs3_JwBuHhiKyLLtNTZ2TfdsIBzKV_ZtaT86bsnxLJvWpjAVuSgj9sLE7lPheuLephUvHCGh5ZyO9hCqeV1CDbxt2xEC2ykb4Rm8wUKW8m0YBd0iKciYPoEcQEgLFU1RFqEsvawuhH0Aj1u76J1_n07Ko6kVAOwrl4rJ7P4QFIX96IlzUpQyytXJGPqNFsaT0A1fpHB4h9UGwnsgJGDCIShqUagefHzxhhYI5-Ve30xRa0GkTCE",
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
&client_id=e726c3f8-762e-4a65-b30d-6efe3ac56075
&client_secret=CJwzrQqPG-BC
&scope=service:iot-hub-prod:t0262c358aab544399b78b4811bfd862b_hub/full-access%20service:iot-manager:0262c358-aab5-4439-9b78-b4811bfd862b_iot-manager/full-access%20service:iot-rollouts:0262c358-aab5-4439-9b78-b4811bfd862b_rollouts/full-access%20service:iot-things-eu-1:0262c358-aab5-4439-9b78-b4811bfd862b_things/full-access

*/
func Authorize(conf *config.Configuration) string {
	
	var clientId = conf.ClientId
	var clientSecret = conf.ClientSecret
	var scope = conf.Scope
	
	if clientId == "" {
		fmt.Println(utils.Fatal("You need to specify the OAuth2 Client ID with -clientId"))
		fmt.Println("\nTo read about the command line options, use '" +os.Args[0] + "auth -h'",);
		fmt.Println("See https://accounts.bosch-iot-suite.com/oauth2-clients/");
		os.Exit(2)
	}

	if clientSecret == "" {
		fmt.Println(utils.Fatal("You need to specify the OAuth2 Client Secret with -clientSecret"))
		fmt.Println("\nTo read about the command line options, use '" +os.Args[0] + "auth -h'",);
		fmt.Println("See https://accounts.bosch-iot-suite.com/oauth2-clients/");
		os.Exit(2)
	}
	
	if scope == "" {
		fmt.Println(utils.Fatal("You need to specify the OAuth2 Scope with -scope"))
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
		fmt.Println(utils.Fatal(response.Status))
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
	
	fmt.Println(utils.Teal("Token Type:"), utils.Warn(responseObject.TokenType))
	fmt.Println(utils.Teal("Scope:"), utils.Warn(responseObject.Scope))
	fmt.Println(utils.Teal("Access Token:"), utils.Warn(responseObject.AccessToken))
	fmt.Println()
	fmt.Println(utils.Warn("Warning:"),"Access token will expire in", responseObject.ExpiresIn,"seconds.")
	
	return responseObject.AccessToken
}