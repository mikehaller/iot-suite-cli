package iotsuite

import (
	"net/http"
)

func ThingsSolutions(httpClient *http.Client, conf *Configuration) {
	var url = "https://things.eu-1.bosch-iot-suite.com/api/2/solutions/"
	Get(httpClient, url)
}

func Things(httpClient *http.Client, conf *Configuration) {
	var url = "https://things.eu-1.bosch-iot-suite.com/api/2/search/things?fields="+conf.Fields
	Get(httpClient, url)
}
