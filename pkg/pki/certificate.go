package pki

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"kertificate.io/kertificate/pkg/db"
)

type CreateCertificateData struct {
	CountryCode        string
	State              string
	Locality           string
	StreetAddress      string
	PostalCode         string
	Organization       string
	OrganizationalUnit string
	CommonName         string
	ValidFor           int
	KeySize            int
}

type KeyGenerator struct {
	commonAuthorityDAO *db.CommonAuthorityDAO
}

func NewKeyGenerator(commonAuthorityDAO *db.CommonAuthorityDAO) *KeyGenerator {
	return &KeyGenerator{commonAuthorityDAO}
}

type ICertificate interface {
	GetCertificateData() []byte
}

type IPrivateKey interface {
	GetPrivateKeyData() []byte
}

func (generator *KeyGenerator) CreateCommonAuthority(data CreateCertificateData) (*x509.Certificate, []byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, data.KeySize)
	if err != nil {
		return nil, nil, nil, err
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
		return nil, nil, nil, err
	}

	return ca, privateKeyBytes, certificateBytes, nil
}

func (generator *KeyGenerator) CreateCertificate(commonAuthorityId int, data CreateCertificateData) (*x509.Certificate, []byte, []byte, error) {
	info, err := generator.commonAuthorityDAO.LoadCommonAuthority(commonAuthorityId)
	if err != nil {
		return nil, nil, nil, err
	}

	commonAuthority, err := ToCertificate(info)
	if err != nil {
		return nil, nil, nil, err
	}

	commonAuthorityPrivateKey, err := x509.ParsePKCS1PrivateKey(info.GetPrivateKeyData())
	if err != nil {
		return nil, nil, nil, err
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, data.KeySize)
	if err != nil {
		return nil, nil, nil, err
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	now := time.Now()
	certificate := &x509.Certificate{
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

	certificateBytes, err := x509.CreateCertificate(rand.Reader, certificate, commonAuthority, &privateKey.PublicKey, commonAuthorityPrivateKey)
	if err != nil {
		return nil, nil, nil, err
	}

	return certificate, privateKeyBytes, certificateBytes, nil
}
