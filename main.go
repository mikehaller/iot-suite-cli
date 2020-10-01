// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt" // printf system console output
	"github.com/fatih/color"
	"github.com/mikehaller/iot-suite-cli/cmd"
	"github.com/mikehaller/iot-suite-cli/iotsuite"
)

var (
	AppVersion = "v0.0.1" 
)

var (
	conf *iotsuite.Configuration
)

func main() {
	conf = iotsuite.ReadConfig()
	if conf.NoColor {
		color.NoColor = true
	}

	color.Cyan("Bosch IoT Suite CLI "+AppVersion+"\nCopyright (c) Bosch.IO GmbH, All right reserved.")
	fmt.Println()

	cmd.Execute()
	
	fmt.Println()
	color.Magenta("#likeabosch\n")
}
