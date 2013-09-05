package config

import (
  "testing"
)

func TestRead(t *testing.T) {

  csvFile := "test_config.csv"
 
  values, err := Read(csvFile) 
  
  if err != nil {
    t.Error("Should not return an error with the valid file")
  }
  
  if len(values) != 2 {
    t.Errorf("Incorrect length for keys, it is %v and should be 2", len(values))
  }
  
  if values["FirstKey"] != "FirstValue" {
    t.Errorf("first key's value was %v", values["FirstKey"])
  }

  if values["SecondKey"] != "SecondValue" {
    t.Errorf("second key#s value was %v", values["SecondKey"])
  } 
  
}

func TestFindFile(t *testing.T) {

  csvFile := "random"

   values, err := Read(csvFile)

  if err == nil {
    t.Error("Should error trying to read non existing file")
  }

  if values != nil {
    t.Error("Should not have values")
  }

}

  
