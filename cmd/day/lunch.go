package main

import (
  "github.com/lauramosher/daisy/cmd/slack"
  "github.com/lauramosher/daisy/cmd/util"
)

func lunch(args []string) {
  if util.Include(args, "-s") || util.Include(args, "--skip-message") {
    util.PrintWarn("\u2757 Skipping Slack message")
  } else {
    util.PrintPlain("Posting message to Slack\t\t\t")
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
    util.PrintClear("Done!\n")
  }

  util.PrintPlain("Setting Slack status & presence...\t\t")
  slack.SetPresence("away")
  slack.SetStatus("om nom nom", ":lunch:")
  util.PrintClear("Done!\n")

  util.PrintCallout("Enjoy your lunch! Om nom nom...\n")
}
