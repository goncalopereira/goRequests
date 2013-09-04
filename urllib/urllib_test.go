package urllib

import "testing"

func TestCreateUrl(t *testing.T) {
  urlFormat := "http://m%v/track/%v/format/%v"
  poolId, trackId, formatId := 1, 2, 3
  u := "http://m1/track/2/format/3"

  result, err := CreateUrl(urlFormat, poolId, trackId, formatId)
  
  if result.String() != u {
    t.Errorf("expected %v and got %v", u, result)
  }
  
  if err != nil {
    t.Error("Unexpected error returned")
  }
}
