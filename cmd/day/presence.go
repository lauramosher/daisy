package main

import (
  "github.com/lauramosher/daisy/cmd/slack"
  "github.com/lauramosher/daisy/cmd/util"
)

func back(args []string) {
  util.PrintPlain("Setting Slack presence...\t\t")
  slack.SetPresence("auto")
  util.PrintClear("Done!\n")

  util.PrintPlain("Setting Slack status...\t\t\t")
  slack.SetStatus("Working remotely", ":house_with_garden:")
  util.PrintClear("Done!\n")

  if util.Include(args, "-m") || util.Include(args, "--message") {
    for i, v := range args {
      if v == "-m" || v =="--message" {
        util.PrintPlain("Posting message to Slack...\t\t")
        value :=  args[i+1]
        slack.PostMessage(value)
        util.PrintClear("Done!\n")
      }
      break
    }
  }
  util.PrintCallout("Welcome back!\n")
}

func away(args []string) {
  util.PrintPlain("Setting Slack presence...\t\t")
  slack.SetPresence("away")
  util.PrintClear("Done!\n")

  util.PrintPlain("Setting Slack status...\t\t\t")
  slack.SetStatus("Heads down", ":headphones:")
  util.PrintClear("Done!\n")

  if util.Include(args, "-m") || util.Include(args, "--message") {
    for i, v := range args {
      if v == "-m" || v =="--message" {
        util.PrintPlain("Posting message to Slack...\t\t")
        value :=  args[i+1]
        slack.PostMessage(value)
        util.PrintClear("Done!\n")
      }
      break
    }
  }
}
