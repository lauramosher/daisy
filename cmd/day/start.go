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
    util.PrintWarn("\u2757 Skipping Slack message\n")
  } else {
    util.PrintPlain("Posting message to Slack...\t\t\t")
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
    util.PrintClear("Done!\n")
  }
  util.PrintPlain("Setting Slack status & presence...\t\t")
  slack.SetStatus("Working Remotely", ":house_with_garden:")
  slack.SetPresence("auto")
  util.PrintClear("Done!\n")

  util.PrintPlain("Running PCO commands...\n")
  boxUpdateApps()
  boxStart()

  util.PrintClear("Done!\n")
  util.PrintCallout("Good morning! Have a fantastic day today!")
}

func boxUpdateApps() {
  cmd := exec.Command("box", "update-apps")
  handleStdoutPipe(cmd)
}

func boxStart() {
  cmd := exec.Command("box", "start")

  stdout, err := cmd.StdoutPipe()

  err = cmd.Start()
  handleError(err)

  defer cmd.Wait()

  buff := bufio.NewScanner(stdout)
  updateBox := false

  go func() {
    for buff.Scan() {
      util.PrintPlain(buff.Text() + "\n")

      matched, err := regexp.MatchString("You should upgrade", buff.Text())
      handleError(err)
      if matched {
        updateBox = true
        util.PrintWarn("Aborting to update PCO Box\n")
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

  cmd := exec.Command("sh", "-c", "cd $HOME/pco-box && git pull && bin/box update")
  handleStdoutPipe(cmd)
}
