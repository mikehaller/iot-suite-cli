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
	conf *iotsuite.Configuration
)

func main() {
	conf = iotsuite.ReadConfig()
	if conf.NoColor {
		color.NoColor = true
	}

	color.Cyan("Bosch IoT Suite CLI v0.1\nCopyright (c) Bosch.IO GmbH, All right reserved.")
	fmt.Println()

	cmd.Execute()
	
	fmt.Println()
	color.Magenta("#likeabosch\n")
}
