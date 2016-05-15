package main

import (
  "os"
  "fmt"
  "log"
  "flag"
  "bufio"
  "strconv"
  "net/http"
  "crypto/rand"
  "encoding/pem"
  "github.com/dreae/colors"

  "github.com/dreae/erebus/lib/tls"
  "github.com/dreae/erebus/lib/config"
  rpc "github.com/dreae/erebus/lib/rpc/server"
)

func main() {
  configPath := flag.String("c", "config.json", "Path to the config file")

  var conf *config.Config
  _, err := os.Stat(*configPath)
  if err != nil {
    conf = firstTimeSetup()
    if err = config.WriteConfig(*configPath, conf); err != nil {
      panic(err)
    }
  } else {
    if conf, err = config.ReadConfig(*configPath); err != nil {
      log.Fatal("Error reading config: ", err)
    }
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

func firstTimeSetup() *config.Config {
  baseConf := config.DefaultConfig()
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Config file doesn't exist, starting first run setup")

  fmt.Println("Please enter a port for agent communication")
  fmt.Printf("Port (%s): ", colors.Green(strconv.Itoa(baseConf.RPCPort)))
  text, _ := reader.ReadString('\n')
  if text != "" {
    port, err := strconv.Atoi(text)
    if err == nil {
      baseConf.RPCPort = port
    }

  }

  fmt.Println("Please enter a port for the web server to listen on")
  fmt.Printf("Port (%s): ", colors.Green(strconv.Itoa(baseConf.WebPort)))
  text, _ = reader.ReadString('\n')
  if text != "" {
    port, err := strconv.Atoi(text)
    if err == nil {
      baseConf.WebPort = port
    }
  }

  fmt.Println("Generating registration token key")
  key := make([]byte, 128)
  if _, err := rand.Read(key); err != nil {
    panic(err)
  }
  baseConf.APIKey = fmt.Sprintf("%x", key)

  fmt.Println("Generating CA Certificate")
  key, cert, err := tls.GenerateCACert("Erebus")
  if err != nil {
    panic(err)
  }
  caPath := "ca.pem"

  fmt.Printf("Writing cert to %s\n", caPath)
  certFile, err := os.OpenFile(caPath, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0666)
  if err != nil {
    panic(err)
  }
  defer certFile.Close()
  pem.Encode(certFile, &pem.Block{
    Type: "CERTIFICATE",
    Bytes: cert,
  })

  keyPath := "caKey.pem"
  fmt.Printf("Writing key to %s\n", keyPath)
  keyFile, err := os.OpenFile(keyPath, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0600)
  if err != nil {
    panic(err)
  }
  defer keyFile.Close()
  pem.Encode(keyFile, &pem.Block{
    Type: "RSA PRIVATE KEY",
    Bytes: key,
  })
  baseConf.CACertFile = caPath
  baseConf.CAKeyFile = keyPath

  return baseConf
}
