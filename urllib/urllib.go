package urllib 

import (
  "strconv"
  "net/url"
  "fmt"
  "errors"
  "github.com/goncalopereira/go_requests/debug"
) 
 
func CreateUrl(urlFormat string, poolId, trackId, formatId int) (u *url.URL, err error) {
  p := strconv.Itoa(poolId)
  t := strconv.Itoa(trackId)
  f := strconv.Itoa(formatId)
   
  rawUrl := fmt.Sprintf(urlFormat, p, t, f)
   
  u, err = url.ParseRequestURI(rawUrl)
  return
}

 
func RequestValidation(u *url.URL) (poolId, trackId, formatId int, err error) {

  debug.ShowUrl(u)
   
  values := u.Query()
  
  // http://localhost:8888/?trackId=29355149&formatId=17&poolId=66"
  if values["poolId"] == nil || values["trackId"] == nil || values["formatId"] == nil {
    err = errors.New("missing parameters")
    return
  }
   
  poolId, err = strconv.Atoi(values["poolId"][0])
  trackId, err = strconv.Atoi(values["trackId"][0])
  formatId, err = strconv.Atoi(values["formatId"][0])
  return
}

