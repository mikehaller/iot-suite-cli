// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt" // printf system console output
	"github.com/fatih/color"
	"github.com/mikehaller/iot-suite-cli/cmd"
)

var (
	AppVersion = "v0.0.1" 
)

func main() {
	color.Cyan("Bosch IoT Suite CLI "+AppVersion+"\nCopyright (c) Bosch.IO GmbH, All right reserved.")
	fmt.Println()
	color.Unset() // Don't forget to unset

	cmd.Execute()
	
	fmt.Println()
	color.Magenta("#likeabosch\n")
	color.Unset() // Don't forget to unset
}
