package main

import (
  "os"
  "log"
  "flag"
  "encoding/pem"
  "github.com/dreae/erebus/lib/tls"
)

func main() {
  organization := flag.String("org", "Erebus", "Name of CA organization")
  outfile := flag.String("out", "cert.pem", "File to output certificate to")
  keyfile := flag.String("key", "key.pem", "File to output private key to")

  flag.Parse()
  log.Print("Generating CA Certificate")
  key, cert, err := tls.GenerateCACert(*organization)
  if err != nil {
    panic(err)
  }

  log.Printf("Writing cert to %s\n", *outfile)
  certFile, err:= os.OpenFile(*outfile, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0666)
  if err != nil {
    panic(err)
  }
  defer certFile.Close()
  pem.Encode(certFile, &pem.Block{
    Type: "CERTIFICATE",
    Bytes: cert,
  })

  log.Printf("Writing key to %s\n", *keyfile)
  keyFile, err:= os.OpenFile(*keyfile, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0600)
  if err != nil {
    panic(err)
  }
  defer keyFile.Close()
  pem.Encode(keyFile, &pem.Block{
    Type: "RSA PRIVATE KEY",
    Bytes: key,
  })

}
