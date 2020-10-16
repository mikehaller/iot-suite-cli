package iotsuite

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"strings"
	"context"
	"github.com/spf13/viper"
	"net/http/httptrace"
	"crypto/tls"
)

func InitOAuth(conf *Configuration) *http.Client {
	oauthConf := clientcredentials.Config{
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		Scopes:       strings.Fields(conf.Scope),
		TokenURL:     viper.GetString("tokenUrl"),
		AuthStyle:    oauth2.AuthStyleInParams}
	
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	
	client := oauthConf.Client(httptrace.WithClientTrace(context.Background(), newHttpTrace()))
	// client := oauthConf.Client(context.Background())
	
	//	{"error":"invalid_request","error_description":"The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed","error_hint":"Client credentials missing or malformed in both HTTP Authorization header and HTTP POST body.","status_code":400}
	
	return client
}
