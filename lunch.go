package main

func lunch(args []string) {
  setPresence("away")
  setStatus("om nom nom", ":lunch:")
  postMessage(":lunch:")
}
