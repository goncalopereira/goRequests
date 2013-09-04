package main

import (
  "github.com/goncalopereira/go_requests/urllib"
  "github.com/goncalopereira/go_requests/debug"
  "github.com/goncalopereira/go_requests/config"
  "log"
  "net/http"
)

func main() {
  poolId := 66
  trackId := 29355149
  formatId := 17

  values, err := config.Read("config.csv")
  
  if err != nil {
    log.Fatal(err)
  }  

  u, err := urllib.CreateUrl(values["urlFormat"], poolId, trackId, formatId)

  if err != nil {
    log.Fatal(err)
  }

  debug.ShowUrl(u)
  
  res, err := http.Get(u.String())  

  if err != nil {
    log.Fatal(err)
  }

  defer res.Body.Close()
  
  if res.StatusCode != http.StatusOK {
    log.Fatal(res.Status)
  }

  debug.PrintInternalHeaders(res)  
}
