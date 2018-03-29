package main

func back(args []string) {
  setPresence("auto")
  setStatus("Working remotely", ":house_with_garden:")
}

func away(args []string) {
  setPresence("away")
  setStatus("Heads down", ":headphones:")
}
