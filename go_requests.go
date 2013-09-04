package main

import (
  "github.com/goncalopereira/go_requests/urllib"
  "github.com/goncalopereira/go_requests/debug"
  "fmt"
  "net/http"
)

func main() {
  poolId := 66
  trackId := 29355149
  formatId := 17

  u, err := urllib.CreateUrl(poolId, trackId, formatId)

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(u)
  
  res, err := http.Get(u.String())
  defer res.Body.Close()

  debug.PrintInternalHeaders(res)  
}
