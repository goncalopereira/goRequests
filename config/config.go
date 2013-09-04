package config

import (
  "encoding/csv"
  "os"
)

func Read(configFile string) (config map[string]string, err error) {
  file, err := os.Open(configFile)
  
  if err != nil {
    return
  }
  
  defer file.Close()

  reader := csv.NewReader(file)
  lines, err := reader.ReadAll()
  
  config = make(map[string]string)
  
  for _, i := range lines {
    config[i[0]] = i[1]
  }
   
  return
}


