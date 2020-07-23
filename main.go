// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt" // printf system console output
	"github.com/fatih/color"
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"os"
)

var (
	conf *iotsuite.Configuration
)

func main() {
	fmt.Println("Bosch IoT Suite CLI v0.1\nCopyright (c) Bosch.IO GmbH, All right reserved.")
	fmt.Println()
	iotsuite.Hello("Mike")

	//args := os.Args[1:]

	conf = iotsuite.ReadConfig()

	if conf.Verbose {
		fmt.Println("Configuration:")
		fmt.Printf("%#v\n\n", conf)
	}

	switch os.Args[1] {
	case "status":
		iotsuite.ShowServiceStatusHealth(conf)
	case "auth":
		iotsuite.Authorize(conf)
	case "oauth2":
		var httpClient = iotsuite.InitOAuth(conf)
		resp, err := httpClient.Get("https://things.eu-1.bosch-iot-suite.com/api/2/search/things")
		if err != nil {
			panic(err)
		}
		if conf.Verbose {
			fmt.Println("Things API Response:")
			fmt.Printf("%#v\n\n", resp)
		} else {
			fmt.Println("Successfull authorization and tested with Things API.")
			fmt.Println(resp.Status)
		}
	case "things":
		iotsuite.Things(iotsuite.InitOAuth(conf), conf)
	case "solution":
		iotsuite.ThingsSolutions(iotsuite.InitOAuth(conf), conf)
	case "rules":
		iotsuite.IotmgrRules(iotsuite.InitOAuth(conf), conf)
	case "devices":
		iotsuite.IotmgrDevices(iotsuite.InitOAuth(conf), conf)
	case "tasks":
		iotsuite.IotmgrTasks(iotsuite.InitOAuth(conf), conf)
	case "groups":
		iotsuite.IotmgrGroups(iotsuite.InitOAuth(conf), conf)
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}

	fmt.Println()
	color.Magenta("#likeabosch\n")
}
