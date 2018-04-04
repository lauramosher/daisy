package main

import (
  "os/exec"

  "github.com/lauramosher/daisy/cmd/slack"
  "github.com/lauramosher/daisy/cmd/util"
)

func start(args []string) {
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
      slack.PostMessage("Good morning! :city_sunrise:")
    }
    util.PrintInfo("Message posted")
  }
  slack.SetStatus("Working Remotely", ":house_with_garden:")
  slack.SetPresence("auto")
  boxUpdateApps()
  boxStart()
}

func boxUpdateApps() {
  util.PrintInfo("box update-apps")

  cmd := exec.Command("box", "update-apps")
  handleStdoutPipe(cmd)
}

func boxStart() {
  util.PrintInfo("box start")

  cmd := exec.Command("box", "start")
  handleStdoutPipe(cmd)
}

func boxUpgrade() {
  util.PrintInfo("Upgrading pco-box")

  // look into os.Getenv("HOME")
  cmd := exec.Command("sh", "-c", "cd $HOME/pco-box && git pull && bin/box update")
  handleStdoutPipe(cmd)
}
