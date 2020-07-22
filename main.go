// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json" // json rest api responses
    "fmt"  // printf system console output
    "io/ioutil"
    "log"
    "net/http"
   
    "os"
    "strings"
    "time" // print current date and time
    
    "github.com/mikehaller/iot-suite-cli/iotsuite"
    
    "github.com/TylerBrock/colorjson"
)

// STATUS.BOSCH-IOT-SUITE.COM JSON
type Response struct {
    Name    string    `json:"name"`
    StatusComponents []StatusComponent `json:"data"`
}

/*

{
"id": 1,
"name": "Bosch IoT Things (EU-1)",
"description": "Managed inventory of digital twins for IoT device assets",
"link": "https://www.bosch-iot-suite.com/things/",
"status": 1,
"order": 16,
"group_id": 1,
"enabled": true,
"meta": null,
"created_at": "2018-06-15 13:09:48",
"updated_at": "2020-04-20 08:30:02",
"deleted_at": null,
"status_name": "Operational",
"tags": {
"": ""
}
},

*/
type StatusComponent struct {
	Id int `json:"id"`
	Name string `json:"name"`
	StatusName string `json:"status_name"`
	Description string `json:"description"`
	Link string `json:"link"`
	Status int `json:"status"`
	Order int `json:"order"`
	UpdatedAt string `json:"updated_at"`
}

func showServiceStatusHealth(conf *iotsuite.Configuration) {
	dt := time.Now().UTC()
	fmt.Println(iotsuite.Teal("Service Status Health as of ", dt.String()," in region ",conf.Region," sorted by ",conf.Sort));
	
	response, err := http.Get("https://status.bosch-iot-suite.com/api/v1/components?sort="+conf.Sort+"&per_page=50")
	if err != nil {
        fmt.Printf("Error: %s",err.Error())
        os.Exit(1)
    }
	responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
        os.Exit(2)
    }
    
    //fmt.Println(string(responseData)
    var responseObject Response
	json.Unmarshal(responseData, &responseObject)

    //Print raw JSON
	//fmt.Println(responseObject.Name)
	//fmt.Println(len(responseObject.StatusComponents))
	
	for i := 0; i < len(responseObject.StatusComponents); i++ {
		var statusComponent = responseObject.StatusComponents[i]
		if conf.Region=="all" || strings.Contains(statusComponent.Name, strings.ToUpper(conf.Region)) { 
			if conf.Verbose {
				fmt.Printf("\n%s\n",iotsuite.Purple(statusComponent.Name))
				fmt.Printf("\t%15s %-60s\n","Description:",iotsuite.Teal(statusComponent.Description))
				fmt.Printf("\t%15s %-60s\n","Link:",iotsuite.Teal(statusComponent.Link))
				fmt.Printf("\t%15s %-60s\n","Updated At",iotsuite.Teal(statusComponent.UpdatedAt))
				fmt.Printf("\t%15s %-60s\n","Status:",iotsuite.Green(statusComponent.StatusName))
			} else {
				fmt.Printf("%-60s %10s\n",statusComponent.Name,iotsuite.Green(statusComponent.StatusName))
			}
		} else {
			// fmt.Println("Filtered by region")
		}
	}
}



func things(accessToken string, conf *iotsuite.Configuration) {
	if accessToken == "" {
		fmt.Println(iotsuite.Fatal("Not authenticated, please authorize first with the 'auth' command"))
		os.Exit(3)
	}
	
	client := &http.Client{}

	fmt.Println(iotsuite.Teal("Requested Fields:"), iotsuite.Warn(conf.Fields))
	
	req,err1 := http.NewRequest("GET", "https://things.eu-1.bosch-iot-suite.com/api/2/search/things?fields="+conf.Fields, nil)
	if err1 != nil {
        log.Fatal(err1)
        os.Exit(2)
    }
	req.Header.Add("Authorization", `Bearer `+ accessToken)
	
	resp,err2 := client.Do(req)
	if err2 != nil {
        log.Fatal(err2)
        os.Exit(2)
    }

	fmt.Println("Response:")
	fmt.Println(iotsuite.Magenta(resp.Status))
	responseData,err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
        log.Fatal(err3)
        os.Exit(2)
    }
	var obj map[string]interface{}
		json.Unmarshal([]byte(responseData), &obj)
	    // Make a custom formatter with indent set
	    f := colorjson.NewFormatter()
	    f.Indent = 4
	    // Marshall the Colorized JSON
	    s, _ := f.Marshal(obj)
	    fmt.Println(string(s))
}

var (
	conf *iotsuite.Configuration
)

func main() {
	iotsuite.InitWindowsColors()
	fmt.Println("Bosch IoT Suite CLI v0.1\nCopyright (c) Bosch.IO GmbH, All right reserved.")
	fmt.Println()
	iotsuite.Hello("Mike")
	
	//args := os.Args[1:]
	
	conf = iotsuite.ReadConfig()
	
	if conf.Verbose {
		fmt.Println("Verbose:",iotsuite.Teal(conf.Verbose))
		fmt.Println()
		fmt.Println("clientId:",iotsuite.Teal(conf.ClientId))
		fmt.Println("clientSecret:",iotsuite.Teal(conf.ClientSecret))
		fmt.Println("scope:",iotsuite.Teal(conf.Scope))
		fmt.Println()
		fmt.Println("config",conf)
	}
	
	switch os.Args[1] {
		case "status":
			showServiceStatusHealth(conf)
		case "auth":
			iotsuite.Authorize(conf)
		case "things":
	        token := iotsuite.Authorize(conf)
			things(token, conf)
		default:
			fmt.Println(iotsuite.Warn("Unknown command:", os.Args[1]))
	}
	
	fmt.Println(iotsuite.Magenta("\n#likeabosch"))
}
