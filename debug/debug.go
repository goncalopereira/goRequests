package debug

import (
  "net/http"
  "strconv"
  "strings"
  "fmt"
)

func PrintInternalHeaders(res *http.Response) {
  fmt.Println("Content Length:" + strconv.FormatInt(res.ContentLength,10))
  fmt.Println("Content-Type:" + strings.Join(res.Header["Content-Type"],","))
  fmt.Println("Content-Disposition:" + strings.Join(res.Header["Content-Disposition"],","))
  fmt.Println("X-7dig:" + strings.Join(res.Header["X-7dig"],","))
  fmt.Println("Last-Modified:" + strings.Join(res.Header["Last-Modified"],","))
}

