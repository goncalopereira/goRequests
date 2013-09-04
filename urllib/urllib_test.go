package urllib

import "testing"

func TestCreateUrl(t *testing.T) {
  poolId, trackId, formatId := 1, 2, 3
  u := "http://mediapool1.nix.sys.7d/track/2/format/3"

  result, err := CreateUrl(poolId, trackId, formatId)
  
  if result.String() != u {
    t.Errorf("expected %v and got %v", u, result)
  }
  
  if err != nil {
    t.Error("Unexpected error returned")
  }
}
