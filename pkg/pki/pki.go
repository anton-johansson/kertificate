// Copyright 2019 Anton Johansson
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
