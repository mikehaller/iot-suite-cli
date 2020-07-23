package iotsuite

import (
	"net/http"
)

func RolloutsTargets(httpClient *http.Client, conf *Configuration) {
	var url = "https://api.eu1.bosch-iot-rollouts.com/rest/v1/targets"
	Get(httpClient, url)
}