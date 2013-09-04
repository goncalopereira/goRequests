package main

import (
  "github.com/goncalopereira/go_requests/urllib"
  "github.com/goncalopereira/go_requests/debug"
  "log"
  "net/http"
)

func main() {
  poolId := 66
  trackId := 29355149
  formatId := 14

  u, err := urllib.CreateUrl(poolId, trackId, formatId)

  if err != nil {
    log.Fatal(err)
  }

  debug.ShowUrl(u)
  
  res, err := http.Get(u.String())  
  defer res.Body.Close()

  if err != nil {
    log.Print("got an error from get")
    log.Fatal(err)
  }

  if res.StatusCode != http.StatusOK {
    log.Fatal(res.Status)
  }

  debug.PrintInternalHeaders(res)  
}
