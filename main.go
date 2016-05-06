package main

import (
  "os"
  "fmt"
  "log"
  "flag"
  "net/http"

  "github.com/dreae/erebus/lib/config"
  rpc "github.com/dreae/erebus/lib/rpc/server"
)

func main() {
  configPath := flag.String("c", "config.json", "Path to the config file")
  _, err := os.Stat(*configPath)
  if err != nil {
    log.Fatal("Config file doesn't exist: ", *configPath)
  }
  var conf *config.Config
  if conf, err = config.ReadConfig(*configPath); err != nil {
    log.Fatal("Error reading config: ", err)
  }

  rpc.Init(conf)

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world")
  })
  if conf.CertFile == "" {
    http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", conf.WebPort), nil)
  } else {
    http.ListenAndServeTLS(fmt.Sprintf("0.0.0.0:%d", conf.WebPort), conf.CertFile, conf.KeyFile, nil)
  }
}
