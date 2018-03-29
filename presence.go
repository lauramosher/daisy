package main

func back(args []string) {
  setPresence("auto")
  setStatus("Working remotely", ":house_with_garden:")

  if Include(args, "-m") || Include(args, "--message") {
    for i, v := range args {
      if v == "-m" || v =="--message" {
        value :=  args[i+1]
        postMessage(value)
      }
      break
    }
  }
}

func away(args []string) {
  setPresence("away")
  setStatus("Heads down", ":headphones:")

  if Include(args, "-m") || Include(args, "--message") {
    for i, v := range args {
      if v == "-m" || v =="--message" {
        value :=  args[i+1]
        postMessage(value)
      }
      break
    }
  }
}
