package urllib 

import (
  "strconv"
  "net/url"
  "fmt"
) 
 
func CreateUrl(urlFormat string, poolId, trackId, formatId int) (u *url.URL, err error) {
  p := strconv.Itoa(poolId)
  t := strconv.Itoa(trackId)
  f := strconv.Itoa(formatId)
   
  rawUrl := fmt.Sprintf(urlFormat, p, t, f)
   
  u, err = url.ParseRequestURI(rawUrl)
  return
}
