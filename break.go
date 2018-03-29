package main

import (
  "time"
)

func setBreak(args []string) {
  setPresence("away")

  if time.Now().Hour() > 12 {
    setStatus("Afternoon Tea", ":tea:")
  } else {
    setStatus("Elevensies", ":croissant:")
  }
}
