package httpdata
import (
  "net/http"
  "strings"
  "io"
  "strconv"
)
 
func CopyResponse(w http.ResponseWriter, res *http.Response) {
   w.Header().Set("Content-Length", strconv.FormatInt(res.ContentLength,10))
   w.Header().Set("Content-Type", strings.Join(res.Header["Content-Type"],","))
   w.Header().Set("Content-Disposition", strings.Join(res.Header["Content-Disposition"],","))
   w.Header().Set("X-7dig", strings.Join(res.Header["X-7dig"],","))
   w.Header().Set("Last-Modified", strings.Join(res.Header["Last-Modified"],","))
   w.WriteHeader(http.StatusOK)
   
   io.Copy(w, res.Body)
} 

