package main

import (
  "time"
)

func setBreak(args []string) {
  setPresence("away")

  if Include(args, "-m") || Include(args, "--message") {
    for i, v := range args {
      if v == "-m" || v =="--message" {
        value :=  args[i+1]
        postMessage(value)
      }
      break
    }
  }

  if time.Now().Hour() > 12 {
    setStatus("Afternoon Tea", ":tea:")
  } else {
    setStatus("Elevensies", ":croissant:")
  }
}
