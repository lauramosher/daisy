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

func setStatus() {
  url := "https://slack.com/api/users.profile.set"
  var jsonStr = []byte(`{"profile": {"status_text":"Working Remotely","status_emoji":":house_with_garden:"}}`)

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

func postMessage() {
  url := "https://slack.com/api/chat.postMessage"
  var jsonStr = []byte(`{"text":"Good morning! :city_sunrise:", "channel": "checkistrations", "as_user":true}`)

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
