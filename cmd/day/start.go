package main

import (
  "bufio"
  "os/exec"
  "regexp"

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

  stdout, err := cmd.StdoutPipe()

  err = cmd.Start()
  handleError(err)

  defer cmd.Wait()

  buff := bufio.NewScanner(stdout)
  updateBox := false

  go func() {
    for buff.Scan() {
      util.PrintPlain(buff.Text())

      matched, err := regexp.MatchString("You should upgrade", buff.Text())
      handleError(err)
      if matched {
        updateBox = true
        util.PrintWarn("Aborting to update PCO Box")
        break
      }
    }

    if updateBox {
      boxUpgrade()
    }
  }()
}

func boxUpgrade() {
  util.PrintInfo("Upgrading pco-box")

  // look into os.Getenv("HOME")
  cmd := exec.Command("sh", "-c", "cd $HOME/pco-box && git pull && bin/box update")
  handleStdoutPipe(cmd)
}
