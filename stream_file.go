package main

import (
  "fmt"
  "strconv"
  "net/url"
)

func createUrl(poolId, trackId, formatId int) (u *url.URL, err error) {
  p := strconv.Itoa(poolId)
  t := strconv.Itoa(trackId)
  f := strconv.Itoa(formatId)

  rawUrl :=  "http://mediapool" + p + ".nix.sys.7d/track/" + t + "/format/" + f 
  
  u, err = url.Parse(rawUrl)
  return
}

func main() {
  poolId := 66
  trackId := 29355149
  formatId := 17

  u, err := createUrl(poolId, trackId, formatId)

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(u)
  fmt.Println(err)
}
