package main

import (
  "time"

  "github.com/lauramosher/daisy/cmd/slack"
  "github.com/lauramosher/daisy/cmd/util"
)

func setBreak(args []string) {
  util.PrintPlain("Setting Slack presence...\t\t")
  slack.SetPresence("away")
  util.PrintClear("Done!\n")

  util.PrintPlain("Setting Slack status...\t\t\t")
  if time.Now().Hour() > 12 {
    slack.SetStatus("Afternoon Tea", ":tea:")
  } else {
    slack.SetStatus("Elevensies", ":croissant:")
  }
  util.PrintClear("Done!\n")

  if util.Include(args, "-m") || util.Include(args, "--message") {
    util.PrintPlain("Posting message to Slack...\t\t")
    for i, v := range args {
      if v == "-m" || v =="--message" {
        value :=  args[i+1]
        slack.PostMessage(value)
      }
      break
    }
    util.PrintClear("Done!\n")
  }

  util.PrintCallout("Enjoy your break!\n")
}
