package iotsuite

import (
	"log"
	"net/http"
	"github.com/fatih/color"
)

func Get(httpClient *http.Client, url string) {
	color.Cyan(url)
	req,err := http.NewRequest(http.MethodGet, url, nil)
	if (err != nil) { log.Fatal(err) }
	resp,err := httpClient.Do(req)
	if (err != nil) { log.Fatal(err) }
	defer resp.Body.Close()
	DumpJsonResponse(resp)
}