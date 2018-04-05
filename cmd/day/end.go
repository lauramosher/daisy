package main

import (
  "os/exec"

  "github.com/lauramosher/daisy/cmd/slack"
  "github.com/lauramosher/daisy/cmd/util"
)

func end(args []string) {
  if util.Include(args, "-s") || util.Include(args, "--skip-message") {
    util.PrintWarn("\u2757 Skipping Slack message")
  } else {
    util.PrintPlain("Posting message to Slack\t\t")

    if util.Include(args, "-m") || util.Include(args, "--message") {
      for i, v := range args {
        if v == "-m" || v =="--message" {
          value :=  args[i+1]
          slack.PostMessage(value)
        }
        break
      }
    } else {
      // default message
      slack.PostMessage("Have a fantastic evening! Until next time... :woman-bowing:")
    }
    util.PrintClear("Done!\n")
  }

  util.PrintPlain("Setting Slack status & presence...\t")
  slack.SetPresence("away")
  slack.SetStatus("EOD", ":hedgehog:")
  util.PrintClear("Done!\n")

  util.PrintPlain("Stopping box...\n")
  boxStop()
  util.PrintClear("Done!\n")

  util.PrintCallout("Have a good evening! :)\n")
}

func boxStop() {
  cmd := exec.Command("box", "stop")
  handleStdoutPipe(cmd)
}
