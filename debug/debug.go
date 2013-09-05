package debug

import (
  "net/http"
  "strconv"
  "strings"
  "log"
  "net/url"
)

func PrintInternalHeaders(res *http.Response) {
  log.Print("Content Length:" + strconv.FormatInt(res.ContentLength,10))
  log.Print("Content-Type:" + strings.Join(res.Header["Content-Type"],","))
  log.Print("Content-Disposition:" + strings.Join(res.Header["Content-Disposition"],","))
  log.Print("X-7dig:" + strings.Join(res.Header["X-7dig"],","))
  log.Print("Last-Modified:" + strings.Join(res.Header["Last-Modified"],","))
}

func ShowUrl(u *url.URL) {
  log.Print(u.String())
}

