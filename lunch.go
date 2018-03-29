package main

func lunch(args []string) {
  if Include(args, "-s") || Include(args, "--skip-message") {
    printWarn("\u2757 Skipping Slack message")
  } else {
    printInfo("Posting message to Slack")
    if Include(args, "-m") || Include(args, "--message") {
      for i, v := range args {
        if v == "-m" || v =="--message" {
          value :=  args[i+1]
          postMessage(value)
        }
        break
      }
    } else {
      postMessage(":lunch:")
    }
    printInfo("Message posted")
  }
  setPresence("away")
  setStatus("om nom nom", ":lunch:")
}
