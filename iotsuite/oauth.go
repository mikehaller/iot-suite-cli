package iotsuite

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"strings"
	"github.com/spf13/viper"
)

func InitOAuth(conf *Configuration) *http.Client {
	oauthConf := clientcredentials.Config{
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		Scopes:       strings.Fields(conf.Scope),
		TokenURL:     viper.GetString("tokenUrl"),
		AuthStyle:    oauth2.AuthStyleInParams}
	client := oauthConf.Client(context.Background())
	return client
}
