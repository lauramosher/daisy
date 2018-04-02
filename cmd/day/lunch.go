package main

import (
  "github.com/lauramosher/daisy/cmd/slack"
  "github.com/lauramosher/daisy/cmd/util"
)

func lunch(args []string) {
  if util.Include(args, "-s") || util.Include(args, "--skip-message") {
    util.PrintWarn("\u2757 Skipping Slack message")
  } else {
    util.PrintInfo("Posting message to Slack")
    if util.Include(args, "-m") || util.Include(args, "--message") {
      for i, v := range args {
        if v == "-m" || v =="--message" {
          value :=  args[i+1]
          slack.PostMessage(value)
        }
        break
      }
    } else {
      slack.PostMessage(":lunch:")
    }
    util.PrintInfo("Message posted")
  }
  slack.SetPresence("away")
  slack.SetStatus("om nom nom", ":lunch:")
}
