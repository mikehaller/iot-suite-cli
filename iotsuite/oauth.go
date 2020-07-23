package iotsuite

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"strings"
)

func InitOAuth(conf *Configuration) *http.Client {
	oauthConf := clientcredentials.Config{
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		Scopes:       strings.Fields(conf.Scope),
		TokenURL:     "https://access.bosch-iot-suite.com/token",
		AuthStyle:    oauth2.AuthStyleInParams}
	client := oauthConf.Client(context.Background())
	return client
}
