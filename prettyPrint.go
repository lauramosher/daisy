package main

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


func printPlain(text string) {
  fmt.Println(text)
}

func printCallout(text string) {
  fmt.Printf(purple)
  fmt.Println(text)
  fmt.Printf(nc)
}

func printInfo(text string) {
  fmt.Printf(blue)
  fmt.Println(text)
  fmt.Printf(nc)
}

func printWarn(text string) {
  fmt.Printf(yellow)
  fmt.Println(text)
  fmt.Printf(nc)
}

func printError(text string) {
  fmt.Printf(red)
  fmt.Println(text)
  fmt.Printf(nc)
}

func printFatal(err error) {
  fmt.Printf(red)
  fmt.Println(err)
  fmt.Printf(nc)
}
