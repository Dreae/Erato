package main

import (
  "os"
  "fmt"
  "flag"
  "math/big"
  "math/rand"
  "io/ioutil"
  "crypto/x509"
  "encoding/pem"
  "github.com/dreae/erebus/lib/tls"
)

func main() {
  organization := flag.String("org", "Erebus", "Name of CA organization")
  outfile := flag.String("out", "cert.pem", "File to output certificate to")
  keyfile := flag.String("key", "key.pem", "File to output private key to")
  cafile := flag.String("ca", "ca.pem", "Path of the CA cert file")
  cakey := flag.String("caKey", "cakey.pem", "Path of CA private key")
  clientId := flag.String("clientId", "Server1", "OrganizationalUnit of client cert")

  flag.Parse()
  command := flag.Arg(0)
  if command == "" {
    command = "ca-cert"
  }

  switch command {
  case "ca-cert":
    fmt.Println("Generating CA Certificate")
    key, cert, err := tls.GenerateCACert(*organization)
    if err != nil {
      panic(err)
    }

    fmt.Printf("Writing cert to %s\n", *outfile)
    certFile, err := os.OpenFile(*outfile, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0666)
    if err != nil {
      panic(err)
    }
    defer certFile.Close()
    pem.Encode(certFile, &pem.Block{
      Type: "CERTIFICATE",
      Bytes: cert,
    })

    fmt.Printf("Writing key to %s\n", *keyfile)
    keyFile, err := os.OpenFile(*keyfile, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0600)
    if err != nil {
      panic(err)
    }
    defer keyFile.Close()
    pem.Encode(keyFile, &pem.Block{
      Type: "RSA PRIVATE KEY",
      Bytes: key,
    })
  case "client-cert":
    keyFile, err := os.Open(*cakey)
    if err != nil {
      panic(err)
    }
    defer keyFile.Close()
    keyB, err := ioutil.ReadAll(keyFile)
    if err != nil {
      panic(err)
    }
    pemBlock, _ := pem.Decode(keyB)
    if pemBlock == nil {
      fmt.Println("Unable to read private key")
    }
    key, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
    if err != nil {
      panic(err)
    }


    caFile, err := os.Open(*cafile)
    if err != nil {
      panic(err)
    }
    defer caFile.Close()
    caB, err := ioutil.ReadAll(caFile)
    if err != nil {
      panic(err)
    }
    pemBlock, _ = pem.Decode(caB)
    if pemBlock == nil {
      fmt.Println("Unable to read CA cert")
    }
    certs, err := x509.ParseCertificates(pemBlock.Bytes)
    if err != nil {
      panic(err)
    }

    ca := tls.NewCertificateAuthority(certs[0], key, *organization, func(_ string) *big.Int {
      return big.NewInt(rand.Int63())
    })

    clientKey, clientCert, err := ca.IssueCertificate(*clientId)
    if err != nil {
      panic(err)
    }

    fmt.Printf("Writing cert to %s\n", *outfile)
    certFile, err := os.OpenFile(*outfile, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0666)
    if err != nil {
      panic(err)
    }
    defer certFile.Close()
    pem.Encode(certFile, &pem.Block{
      Type: "CERTIFICATE",
      Bytes: clientCert,
    })

    fmt.Printf("Writing key to %s\n", *keyfile)
    keyFile, err = os.OpenFile(*keyfile, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0600)
    if err != nil {
      panic(err)
    }
    defer keyFile.Close()
    pem.Encode(keyFile, &pem.Block{
      Type: "RSA PRIVATE KEY",
      Bytes: clientKey,
    })
  default:
    fmt.Println("Must provide a valid command")
  }

}
