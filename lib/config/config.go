package config

import (
  "os"
  "encoding/json"
)

// Config holds master server configuration
type Config struct {
  CACertFile string `json:"caCert"`
  CAKeyFile string `json:"caKey"`
  CertFile string `json:"cert"`
  KeyFile string `json:"key"`
  WebPort int `json:"webport"`
  RPCPort int `json:"rpcport"`
}

// DefaultConfig generates a config containing default settings
func DefaultConfig() *Config {
  return &Config{
    WebPort: 8080,
    RPCPort: 5455,
  }
}

// ReadConfig reads the master server configuration from a file
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

// WriteConfig saves the configuration to a file, truncating the file
// if it exists
func WriteConfig(configPath string, config *Config) error {
  file, err := os.OpenFile(configPath, os.O_RDWR | os.O_TRUNC | os.O_CREATE, 0600)
  if err != nil {
    return err
  }
  defer file.Close()

  encoder := json.NewEncoder(file)
  if err = encoder.Encode(config); err != nil {
    return err
  }
  return nil
}
