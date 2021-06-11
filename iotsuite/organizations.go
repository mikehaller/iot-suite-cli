package iotsuite

import (
	"fmt"
//	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"os"
	// "bytes"
	"encoding/json"
	"github.com/spf13/viper"
	// "net/http/httptrace"
	// "time"
)


type Organization struct {
	OrgName string `json:"orgName"`
	OrgId string `json:"orgId"`
}

func OrgList(httpClient *http.Client) {
	var baseurl = viper.GetString("baseurl")
	var url = baseurl + "/api/v3/organisations";
	log.WithFields(log.Fields{"url": url}).Info("Using Organization Management endpoint")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query() // Get a copy of the query values.
		req.URL.RawQuery = q.Encode() // Encode and assign back to the original query.

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("OrgList HTTP Request generally failed:",err)
	}
	defer resp.Body.Close()

	fmt.Println()

	DumpJsonResponse(resp)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Reading response body for OrgList failed:",err)
		os.Exit(2)
	}

	log.Debug("Body:", string(responseData));

	var responseObject []Organization
	jsonErr := json.Unmarshal([]byte(responseData), &responseObject)
	if (jsonErr != nil) {
		log.Fatal("Unmarshalling response body for OrgList failed:",jsonErr);
		os.Exit(2);
	}

	fmt.Printf("%-38q %-50q\n", "Org ID", "Org Name")
	for i := 0; i < len(responseObject); i++ {
		var org = responseObject[i]
		fmt.Printf("%-38q %-50q\n", org.OrgId, org.OrgName)
	}

}
