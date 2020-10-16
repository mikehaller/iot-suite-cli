package iotsuite

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

// STATUS.BOSCH-IOT-SUITE.COM JSON
type Response struct {
	Name             string            `json:"name"`
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
	Id          int    `json:"id"`
	Name        string `json:"name"`
	StatusName  string `json:"status_name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Status      int    `json:"status"`
	Order       int    `json:"order"`
	UpdatedAt   string `json:"updated_at"`
}

func ShowServiceStatusHealth(region string, sort string, verbose bool, watch bool, waittime int) {

	if waittime < 5 {
		waittime = 5
	}
	if waittime > 600 {
		waittime = 600
	}

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	for {

		dt := time.Now().UTC()
		color.Cyan("Bosch IoT Suite Service Status Health as of %s in region %s sorted by %s", dt.String(), region, sort)
		fmt.Println()

		//color.Info.Tips("tips style message")
		//color.Red.Println("Simple to use color")

		response, err := http.Get("https://status.bosch-iot-suite.com/api/v1/components?sort=" + sort + "&per_page=50")
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
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
			if region == "all" || strings.Contains(statusComponent.Name, strings.ToUpper(region)) {
				if verbose {
					fmt.Printf("\n%s\n", (statusComponent.Name))
					fmt.Printf("\t%15s %-60s\n", "Description:", statusComponent.Description)
					fmt.Printf("\t%15s %-60s\n", "Link:", statusComponent.Link)
					fmt.Printf("\t%15s %-60s\n", "Updated At", statusComponent.UpdatedAt)
					statusComponent.Status = rand.Intn(5)
					switch statusComponent.Status {
					case 1: // Operational
						fmt.Printf("\t%15s %s (%d)\n", "Status:", green(statusComponent.StatusName), statusComponent.Status)
					case 2: // Performance Issues
						fmt.Printf("\t%15s %s (%d)\n", "Status:", yellow(statusComponent.StatusName), statusComponent.Status)
					case 3: // Partial Outage
						fmt.Printf("\t%15s %s (%d)\n", "Status:", yellow(statusComponent.StatusName), statusComponent.Status)
					case 4: // Major Outage
						fmt.Printf("\t%15s %s (%d)\n", "Status:", red(statusComponent.StatusName), statusComponent.Status)
					default:
						fmt.Printf("\t%15s %s (%d)\n", "Status:", statusComponent.StatusName, statusComponent.Status)
					}
				} else {
					switch statusComponent.Status {
					case 1: // Operational
						color.Green("%-60s %10s\n", statusComponent.Name, statusComponent.StatusName)
					case 2: // Performance Issues
						color.Yellow("%-60s %10s\n", statusComponent.Name, statusComponent.StatusName)
					case 3: // Partial Outage
						color.Yellow("%-60s %10s\n", statusComponent.Name, statusComponent.StatusName)
					case 4: // Major Outage
						color.Red("%-60s %10s\n", statusComponent.Name, statusComponent.StatusName)
					default:
						fmt.Printf("%-60s %10s\n", statusComponent.Name, statusComponent.StatusName)
					}
				}
			} else {
				// fmt.Println("Filtered by region")
			}

		}

		if !watch {
			break
		} else {
			color.Yellow("\nRefreshing every %d seconds.\n",waittime)
			time.Sleep(time.Duration(waittime) * time.Second)
			
			screen.Clear()
			screen.MoveTopLeft()
		}

	}

}
