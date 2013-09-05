package main

import (
  "github.com/goncalopereira/go_requests/urllib"
  "github.com/goncalopereira/go_requests/debug"
  "github.com/goncalopereira/go_requests/config"
  "log"
  "net/http"
  "fmt"
  "strconv"
  "strings"
  "io"
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

  Send(w, res)
}

func Send(w http.ResponseWriter, res *http.Response) {
  w.Header().Set("Content-Length", strconv.FormatInt(res.ContentLength,10))
  w.Header().Set("Content-Type", strings.Join(res.Header["Content-Type"],","))
  w.Header().Set("Content-Disposition", strings.Join(res.Header["Content-Disposition"],","))
  w.Header().Set("X-7dig", strings.Join(res.Header["X-7dig"],","))
  w.Header().Set("Last-Modified", strings.Join(res.Header["Last-Modified"],","))
  w.WriteHeader(http.StatusOK)
  
  io.Copy(w, res.Body)
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


