package main

import (
  "os"
  "time"
)

func setBreak(args []string) {
  if len(args) < 1 {
    setAway()
    os.Exit(1)
  }
  switch {
  case args[0] == "-b":
    setBack()
  case args[0] == "--back":
    setBack()
  default:
    setAway()
  }
}

func setAway() {
  setPresence("away")

  if time.Now().Hour() > 12 {
    setStatus("Afternoon Tea", ":tea:")
  } else {
    setStatus("Elevensies", ":croissant:")
  }
}

func setBack() {
  setPresence("auto")
  setStatus("Working Remotely", ":house_with_garden:")
}
