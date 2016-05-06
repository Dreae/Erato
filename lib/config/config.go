package config

import (
  "os"
  "encoding/json"
)

type Config struct {
  CACertFile string `json:"caCert"`
  CAKeyFile string `json:"caKey"`
  CertFile string `json:"cert"`
  KeyFile string `json:"key"`
  WebPort int `json:"webport"`
  RpcPort int `json:"rpcport"`
}

func ReadConfig(configPath string) (*Config, error) {
  configFile, err := os.Open(configPath)
  if err != nil {
    return nil, err
  }

  var config Config
  jsonParser := json.NewDecoder(configFile)
  if err = jsonParser.Decode(&config); err != nil {
    return nil, err
  }

  return &config, nil
}
