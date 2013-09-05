package main

import (
  "github.com/goncalopereira/go_requests/urllib"
  "github.com/goncalopereira/go_requests/debug"
  "github.com/goncalopereira/go_requests/config"
  "log"
  "net/http"
  "fmt"
  "strconv"
)

type configValues struct {
  urlFormat string
}

func (v *configValues) ServeHTTP(w http.ResponseWriter, r *http.Request) {

  log.Print(r.URL.String())
  
  values := r.URL.Query()

 // http://localhost:8888/?trackId=29355149&formatId=17&poolId=66"  
  if values["poolId"] == nil || values["trackId"] == nil || values["formatId"] == nil {
    fmt.Fprintf(w, "missing parameters")
    return
  }

  poolId, err := strconv.Atoi(values["poolId"][0])
  trackId, err := strconv.Atoi(values["trackId"][0])
  formatId, err := strconv.Atoi(values["formatId"][0])

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


