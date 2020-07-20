package utils

import (
	"os"
	"fmt"
	"golang.org/x/sys/windows" // ansi colors
    //"github.com/TylerBrock/colorjson" // json colors
)

func Hello(s string) string {
	return "Hello"
}

// ANSI COLORS FOR CLI OUTPUT
var (
  Info = Teal
  Warn = Yellow
  Fatal = Red
)

var (
  Black   = Color("\033[1;30m%s\033[0m")
  Red     = Color("\033[1;31m%s\033[0m")
  Green   = Color("\033[1;32m%s\033[0m")
  Yellow  = Color("\033[1;33m%s\033[0m")
  Purple  = Color("\033[1;34m%s\033[0m")
  Magenta = Color("\033[1;35m%s\033[0m")
  Teal    = Color("\033[1;36m%s\033[0m")
  White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
  sprint := func(args ...interface{}) string {
    return fmt.Sprintf(colorString,
      fmt.Sprint(args...))
  }
  return sprint
}

func InitWindowsColors() {
    stdout := windows.Handle(os.Stdout.Fd())
    var originalMode uint32

    windows.GetConsoleMode(stdout, &originalMode)
    windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}