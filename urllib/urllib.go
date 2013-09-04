package urllib 

import (
  "strconv"
  "net/url"
) 
 
func CreateUrl(poolId, trackId, formatId int) (u *url.URL, err error) {
  p := strconv.Itoa(poolId)
  t := strconv.Itoa(trackId)
  f := strconv.Itoa(formatId)
   
  rawUrl :=  "http://mediapool" + p + ".nix.sys.7d/track/" + t + "/format/" + f
   
  u, err = url.Parse(rawUrl)
  return
}
