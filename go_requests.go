package main

import (
  "github.com/goncalopereira/go_requests/urllib"
  "github.com/goncalopereira/go_requests/debug"
  "github.com/goncalopereira/go_requests/config"
  "log"
  "net/http"
)

type configValues struct {
  urlFormat string
}

func (v *configValues) ServeHTTP(w http.ResponseWriter, r *http.Request) {

  poolId, trackId, formatId := 66, 29355149, 17

  u, err := urllib.CreateUrl(v.urlFormat, poolId, trackId, formatId)

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

func main() {
  path, port := "/", ":8888"  
  log.Print("listening on " + path  + port)

  v, err := config.Read("config.csv")  
  if err != nil {
    log.Fatal(err)
  }  
  
  setValues := &configValues{urlFormat: v["urlFormat"]}
  
  http.Handle(path, setValues)

  err = http.ListenAndServe(port,nil)
  if err != nil {
    log.Fatal(err)
  }
}


