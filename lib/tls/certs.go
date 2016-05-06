package tls

import (
  "time"
  "math/big"
  "crypto/rsa"
  "crypto/x509"
  "crypto/rand"
  "crypto/x509/pkix"
)

type SerialNumberFunc func(string) *big.Int

type CertificateAuthority struct {
  caCert *x509.Certificate
  caKey *rsa.PrivateKey
  Organization string
  GetSerialNumber SerialNumberFunc
}

func NewCertificateAuthority(cert *x509.Certificate, key *rsa.PrivateKey, organization string, getSerialNumber SerialNumberFunc) *CertificateAuthority {
  return &CertificateAuthority{
    caCert: cert,
    caKey: key,
    Organization: organization,
    GetSerialNumber: getSerialNumber,
  }
}

func (ca *CertificateAuthority) IssueCertificate(instance_id string) (privateKey, certificate []byte, err error) {
  key, err := rsa.GenerateKey(rand.Reader, 2048)
  if err != nil {
    return nil, nil, err
  }

  template := &x509.Certificate{
    IsCA: false,
    SerialNumber: ca.GetSerialNumber(instance_id),
    Subject: pkix.Name{
      Organization: []string{ca.Organization},
      OrganizationalUnit: []string{instance_id},
    },
    ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
    NotBefore: time.Now(),
    NotAfter: time.Now().AddDate(100, 0, 0),
  }

  cert, err := x509.CreateCertificate(rand.Reader, template, ca.caCert, &key.PublicKey, ca.caKey)
  if err != nil {
    return nil, nil, err
  }

  return x509.MarshalPKCS1PrivateKey(key), cert, nil
}

func GenerateCACert(organization string) (privateKey, certificate []byte, err error) {
  key, err := rsa.GenerateKey(rand.Reader, 2048)
  if err != nil {
    return nil, nil, err
  }

  template := &x509.Certificate{
    IsCA: true,
    SerialNumber: big.NewInt(0),
    Subject: pkix.Name{
      Organization: []string{organization},
      OrganizationalUnit: []string{"ErebusCertificateAuthority"},
    },
    KeyUsage : x509.KeyUsageDigitalSignature|x509.KeyUsageCertSign,
    NotBefore: time.Now(),
    NotAfter: time.Now().AddDate(100, 0, 0),
  }

  cert, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
  if err != nil {
    return nil, nil, err
  }

  return x509.MarshalPKCS1PrivateKey(key), cert, nil
}
