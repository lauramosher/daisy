package main

import (
  "time"

  "github.com/lauramosher/daisy/cmd/slack"
  "github.com/lauramosher/daisy/cmd/util"
)

func setBreak(args []string) {
  slack.SetPresence("away")

  if util.Include(args, "-m") || util.Include(args, "--message") {
    for i, v := range args {
      if v == "-m" || v =="--message" {
        value :=  args[i+1]
        slack.PostMessage(value)
      }
      break
    }
  }

  if time.Now().Hour() > 12 {
    slack.SetStatus("Afternoon Tea", ":tea:")
  } else {
    slack.SetStatus("Elevensies", ":croissant:")
  }
}