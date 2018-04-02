package main

import (
  "github.com/lauramosher/daisy/cmd/slack"
  "github.com/lauramosher/daisy/cmd/util"
)

func back(args []string) {
  slack.SetPresence("auto")
  slack.SetStatus("Working remotely", ":house_with_garden:")

  if util.Include(args, "-m") || util.Include(args, "--message") {
    for i, v := range args {
      if v == "-m" || v =="--message" {
        value :=  args[i+1]
        slack.PostMessage(value)
      }
      break
    }
  }
}

func away(args []string) {
  slack.SetPresence("away")
  slack.SetStatus("Heads down", ":headphones:")

  if util.Include(args, "-m") || util.Include(args, "--message") {
    for i, v := range args {
      if v == "-m" || v =="--message" {
        value :=  args[i+1]
        slack.PostMessage(value)
      }
      break
    }
  }
}
