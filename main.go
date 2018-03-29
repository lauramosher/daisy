package main

import (
  "bufio"
  "os"
  "os/exec"
)

const usage =
`dayrsk is your remote friendly Day Remote Starter Kit

Usage:

      day command [arguments]

The commands are:

      start     bootstrap your system to begin your day
      end       close down your system to end your day
      break     alert your team that you are on break
      lunch     alert your team that you are on lunch

The flags are:

      -s --skip-message     Do not post message to Slack.`

func main() {
  if len(os.Args[1:]) < 1 {
    printError("Command not found")
    printPlain(usage)
    os.Exit(42)
  }

  switch {
  case os.Args[1] == "start":
    start(os.Args[2:])
	case os.Args[1] == "break":
		setBreak(os.Args[2:])
	case os.Args[1] == "end":
		end(os.Args[2:])
	case os.Args[1] == "lunch":
		lunch(os.Args[2:])
  default:
    printPlain(usage)
  }

  // always exit
  os.Exit(42)
}

func handleStdoutPipe(cmd *exec.Cmd) {
  stdout, err := cmd.StdoutPipe()

  err = cmd.Start()
  handleError(err)

  defer cmd.Wait()

  buff := bufio.NewScanner(stdout)

  go func() {
    for buff.Scan() {
      printPlain(buff.Text())
    }
  }()
}

func handleError(err error) {
  if err != nil {
    printFatal(err)
    os.Exit(1)
  }
}
