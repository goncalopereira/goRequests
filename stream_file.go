package main

import (
  "fmt"
  "strconv"
  "net/url"
  "net/http"
  "strings"
)

func createUrl(poolId, trackId, formatId int) (u *url.URL, err error) {
  p := strconv.Itoa(poolId)
  t := strconv.Itoa(trackId)
  f := strconv.Itoa(formatId)

  rawUrl :=  "http://mediapool" + p + ".nix.sys.7d/track/" + t + "/format/" + f 
  
  u, err = url.Parse(rawUrl)
  return
}

func printInternalHeaders(res *http.Response) {
  fmt.Println("Content Length:" + strconv.FormatInt(res.ContentLength,10))    
  fmt.Println("Content-Type:" + strings.Join(res.Header["Content-Type"],","))  
  fmt.Println("Content-Disposition:" + strings.Join(res.Header["Content-Disposition"],","))  
  fmt.Println("X-7dig:" + strings.Join(res.Header["X-7dig"],","))    
  fmt.Println("Last-Modified:" + strings.Join(res.Header["Last-Modified"],","))  
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
  
  res, err := http.Get(u.String())
  defer res.Body.Close()

  printInternalHeaders(res)  
}
