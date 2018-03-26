package main

import (
  "fmt"
  "os"
)

const usage =
`dayrsk is your remote friendly Day Remote Starter Kit

Usage:

      day command [arguments]

The commands are:

      start     bootstrap your system to begin your day`

func main() {
  if len(os.Args[1:]) < 1 {
    fmt.Println("Command not found")
    fmt.Println(usage)
    os.Exit(42)
  }

  fmt.Println(usage)

  // always exit
  os.Exit(42)
}
