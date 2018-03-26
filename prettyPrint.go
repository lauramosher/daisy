package main

import (
  "fmt"
)

const(
  red = "\033[0;31m"
  purple = "\033[0;35m"
  nc = "\033[0m"
)


func printPlain(text string) {
  fmt.Println(text)
}

func printInfo(text string) {
  fmt.Printf(purple)
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
