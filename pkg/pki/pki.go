package pki

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

type KeyGenerator struct{}

func NewKeyGenerator() *KeyGenerator {
	return &KeyGenerator{}
}

type ICertificate interface {
	GetCertificateData() []byte
}

type IPrivateKey interface {
	GetPrivateKeyData() []byte
}

func (generator *KeyGenerator) CreateCommonAuthority(data Certificate) ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, data.KeySize)
	if err != nil {
		return nil, nil, err
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	now := time.Now()
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Country:            []string{data.CountryCode},
			Province:           []string{data.State},
			Locality:           []string{data.Locality},
			StreetAddress:      []string{data.StreetAddress},
			PostalCode:         []string{data.PostalCode},
			Organization:       []string{data.Organization},
			OrganizationalUnit: []string{data.OrganizationalUnit},
			CommonName:         data.CommonName,
		},
		NotBefore:             now,
		NotAfter:              now.AddDate(0, 0, data.ValidFor),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	certificateBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}

	return privateKeyBytes, certificateBytes, nil
}

func (generator *KeyGenerator) GeneratePrivateKey() error {
	//template := &x509.Certificate{}

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	if err := pem.Encode(os.Stdout, block); err != nil {
		return err
	}

	publicKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	if err := pem.Encode(os.Stdout, block); err != nil {
		return err
	}

	now := time.Now()
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Country:            []string{"SE"},
			Province:           []string{"Västra Götalands län"},
			Locality:           []string{"Borås"},
			StreetAddress:      []string{"Arlagatan 2"},
			PostalCode:         []string{"503 36"},
			Organization:       []string{"Anton AB"},
			OrganizationalUnit: []string{"Operations"},
			CommonName:         "Anton CA",
		},
		NotBefore:             now,
		NotAfter:              now.AddDate(0, 0, 365),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	certificateBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &privateKey.PublicKey, privateKey)
	if err != nil {
		fmt.Println("lul?")
		return err
	}

	block = &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificateBytes,
	}
	if err := pem.Encode(os.Stdout, block); err != nil {
		return err
	}

	return nil
}

func (generator *KeyGenerator) CertificateToPem(certificate ICertificate) (string, error) {
	return toPem("CERTIFICATE", certificate.GetCertificateData())
}

func (generator *KeyGenerator) PrivateKeyToPem(privateKey IPrivateKey) (string, error) {
	return toPem("RSA PRIVATE KEY", privateKey.GetPrivateKeyData())
}

func toPem(dataType string, dataBytes []byte) (string, error) {
	block := &pem.Block{
		Type:  dataType,
		Bytes: dataBytes,
	}

	buffer := new(bytes.Buffer)
	if err := pem.Encode(buffer, block); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func ToCertificate(certificate ICertificate) (*x509.Certificate, error) {
	return x509.ParseCertificate(certificate.GetCertificateData())
}
