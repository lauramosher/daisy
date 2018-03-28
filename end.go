package main

func end(args []string) {
  setPresence("away")
  postMessage("Have a fantastic evening! Until next time... :woman-bowing:")
  setStatus("EOD", ":hedgehog:")
}
