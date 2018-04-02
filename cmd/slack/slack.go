package slack

import (
  "bytes"
  "fmt"
  "net/http"
  "os"
)

var bearerToken = "Bearer " + os.Getenv("DAISY_SLACK_USER_TOKEN")

func Slack() bool {
  if len(os.Getenv("DAISY_SLACK_USER_TOKEN")) == 0 {
    fmt.Printf("\033[0;31m")
    fmt.Println("\u2527  Missing user token. Please see Slack Integration for installation instructions.")
    fmt.Printf("\033[0m")
    return false
  }
  // TODO: flip this to true
  return true
}

func SetPresence(state string) {
  if Slack() == false {
    return
  }
  url := "https://slack.com/api/users.setPresence"
  var jsonStr = []byte(`{"presence":"`+ state +`"}`)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
  req.Header.Set("Authorization", bearerToken)
  req.Header.Set("Content-Type", "application/json;charset=utf-8")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()
}

func SetStatus(status string, emoji string) {
  if Slack() == false {
    return
  }
  url := "https://slack.com/api/users.profile.set"
  var jsonStr = []byte(`{"profile": {"status_text":"`+ status + `","status_emoji":"`+ emoji +`"}}`)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
  req.Header.Set("Authorization", bearerToken)
  req.Header.Set("Content-Type", "application/json;charset=utf-8")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()
}

func PostMessage(message string) {
  if Slack() == false {
    return
  }
  url := "https://slack.com/api/chat.postMessage"
  var jsonStr = []byte(`{"text":"`+ message +`", "channel": "checkistrations", "as_user":true}`)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
  req.Header.Set("Authorization", bearerToken)
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()
}
