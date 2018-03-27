package main

import (
  "bytes"
  "net/http"
  "os"
)

var bearerToken = "Bearer " + os.Getenv("DAYSY_SLACK_USER_TOKEN")

func setPresence(state string) {
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

func setStatus(status string, emoji string) {
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

func postMessage(message string) {
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
