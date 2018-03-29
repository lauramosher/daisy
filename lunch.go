package main

func lunch(args []string) {
  if Include(args, "-s") || Include(args, "--skip-message") {
    printWarn("\u2757 Skipping Slack message")
  } else {
    printInfo("Posting message to Slack")
    postMessage(":lunch:")
    printInfo("Message posted")
  }
  setPresence("away")
  setStatus("om nom nom", ":lunch:")
}
