package util

import (
  "fmt"
)

const(
  blue = "\033[0;34m"
  green = "\033[0;32m"
  red = "\033[0;31m"
  purple = "\033[0;35m"
  nc = "\033[0m"
  yellow = "\033[0;33m"
)


func PrintPlain(text string) {
  fmt.Println(text)
}

func PrintCallout(text string) {
  fmt.Printf(purple)
  fmt.Println(text)
  fmt.Printf(nc)
}

func PrintInfo(text string) {
  fmt.Printf(blue)
  fmt.Println(text)
  fmt.Printf(nc)
}

func PrintClear(text string) {
  fmt.Printf(green)
  fmt.Println(text)
  fmt.Printf(nc)
}

func PrintWarn(text string) {
  fmt.Printf(yellow)
  fmt.Println(text)
  fmt.Printf(nc)
}

func PrintError(text string) {
  fmt.Printf(red)
  fmt.Println(text)
  fmt.Printf(nc)
}

func PrintFatal(err error) {
  fmt.Printf(red)
  fmt.Println(err)
  fmt.Printf(nc)
}
