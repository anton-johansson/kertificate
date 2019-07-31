package pki

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
)

func ToCertificate(certificate ICertificate) (*x509.Certificate, error) {
	return x509.ParseCertificate(certificate.GetCertificateData())
}

func CertificateToPem(certificate ICertificate) (string, error) {
	return toPem("CERTIFICATE", certificate.GetCertificateData())
}

func PrivateKeyToPem(privateKey IPrivateKey) (string, error) {
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
