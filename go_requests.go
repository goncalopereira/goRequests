package main

import (
  "github.com/goncalopereira/go_requests/urllib"
  "github.com/goncalopereira/go_requests/debug"
  "github.com/goncalopereira/go_requests/config"
  "github.com/goncalopereira/go_requests/httpdata"
  "log"
  "net/http"
  "fmt"
  "strconv"
)

type configValues struct {
  urlFormat string
}

func (v *configValues) ServeHTTP(w http.ResponseWriter, r *http.Request) {

  poolId, trackId, formatId, err := urllib.RequestValidation(r.URL)
  
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w,err.Error())
  }

  u, err := urllib.CreateUrl(v.urlFormat, poolId, trackId, formatId)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, err.Error())
    return    
  }

  debug.ShowUrl(u)
  
  res, err := http.Get(u.String())  

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, err.Error())
    return
  }

  defer res.Body.Close()
  
  if res.StatusCode != http.StatusOK {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Internal Error " + strconv.Itoa(res.StatusCode))
    return
  }

  debug.PrintInternalHeaders(res) 

  httpdata.CopyResponse(w, res)
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


