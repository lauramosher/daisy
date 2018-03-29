package main

import (
  "os/exec"
)

func start(args []string) {
  if Include(args, "-s") || Include(args, "--skip-message") {
    printWarn("\u2757 Skipping Slack message")
  } else {
    printInfo("Posting message to Slack")
    postMessage("Good morning! :city_sunrise:")
    printInfo("Message posted")
  }
  setStatus("Working Remotely", ":house_with_garden:")
  setPresence("auto")
  boxUpgrade()
  boxUpdate()
  boxStart()
}

func boxUpdate() {
  printInfo("box update")

  cmd := exec.Command("box", "update")
  handleStdoutPipe(cmd)
}

func boxStart() {
  printInfo("box start")

  cmd := exec.Command("box", "start")
  handleStdoutPipe(cmd)
}

func boxUpgrade() {
  printInfo("Upgrading pco-box")

  // look into os.Getenv("HOME")
  cmd := exec.Command("sh", "-c", "cd $HOME/pco-box && git pull && bin/box update")
  handleStdoutPipe(cmd)
}
