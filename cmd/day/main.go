package main

import (
  "bufio"
  "os"
  "os/exec"

  "github.com/lauramosher/daisy/cmd/util"
)

const usage =
`daisy is your remote friendly Day Remote Starter Kit

Usage:

      day command [arguments]

The commands are:

      start     bootstrap your system to begin your day
      end       close down your system to end your day
      break     alert your team that you are on break
      lunch     alert your team that you are on lunch
      away      set yourself away for headsdown time in Slack
      return    global return to available in Slack

The flags are:

      -s --skip-message                       Skips sending messages to Slack. Overrides -m or --message flags.
      -m "<message>", --message "<message>"   Send a custom message. Replaces default message for the command given. This will also send a message even if the command typically is message-less.`

func main() {
  if len(os.Args[1:]) < 1 {
    util.PrintPlain(usage)
    os.Exit(0)
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
    case os.Args[1] == "away":
      away(os.Args[2:])
    case os.Args[1] == "return":
      back(os.Args[2:])
    default:
      util.PrintError("Command not found")
      util.PrintPlain(usage)
      os.Exit(1)
  }

  // always exit
  os.Exit(0)
}

func handleStdoutPipe(cmd *exec.Cmd) {
  stdout, err := cmd.StdoutPipe()

  err = cmd.Start()
  handleError(err)

  defer cmd.Wait()

  buff := bufio.NewScanner(stdout)

  go func() {
    for buff.Scan() {
      util.PrintPlain(buff.Text())
    }
  }()
}

func handleError(err error) {
  if err != nil {
    util.PrintFatal(err)
    os.Exit(1)
  }
}
