package main

func end(args []string) {
  if Include(args, "-s") || Include(args, "--skip-message") {
    printWarn("\u2757 Skipping Slack message")
  } else {
    printInfo("Posting message to Slack")
    postMessage("Have a fantastic evening! Until next time... :woman-bowing:")
    printInfo("Message posted")
  }
  setPresence("away")
  setStatus("EOD", ":hedgehog:")
}
