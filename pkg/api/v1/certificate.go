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

package v1

import (
	"fmt"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
	"kertificate.io/kertificate/pkg/db"
	"kertificate.io/kertificate/pkg/pki"
)

type CreateCertificateRequest struct {
	CommonName        string  `json:"commonName"`
	CommonAuthorityId int     `json:"commonAuthorityId"`
	Subject           Subject `json:"subject"`
	ValidFor          int     `json:"validFor"`
	KeySize           int     `json:"keySize"`
}

type CertificateAPI struct {
	generator      *pki.KeyGenerator
	certificateDAO *db.CertificateDAO
}

func NewCertificateAPI(generator *pki.KeyGenerator, certificateDAO *db.CertificateDAO) *CertificateAPI {
	return &CertificateAPI{generator, certificateDAO}
}

func (api *CertificateAPI) Register(group *echo.Group) {
	group.POST("", api.createCertificate)
	group.DELETE("/:certificateId", api.deleteCertificate)
}

func (api *CertificateAPI) createCertificate(context echo.Context) error {
	userId := userId(context)
	var body CreateCertificateRequest
	if err := context.Bind(&body); err != nil {
		return err
	}

	data := pki.CreateCertificateData{
		CountryCode:        body.Subject.CountryCode,
		State:              body.Subject.State,
		Locality:           body.Subject.Locality,
		StreetAddress:      body.Subject.StreetAddress,
		PostalCode:         body.Subject.PostalCode,
		Organization:       body.Subject.Organization,
		OrganizationalUnit: body.Subject.OrganizationalUnit,
		CommonName:         body.CommonName,
		ValidFor:           body.ValidFor,
		KeySize:            body.KeySize,
	}

	certificate, privateKeyBytes, certificateBytes, err := api.generator.CreateCertificate(body.CommonAuthorityId, data)
	if err != nil {
		return err
	}

	certificateId, err := api.certificateDAO.SaveCertificate(body.CommonAuthorityId, body.CommonName, privateKeyBytes, certificateBytes, certificate.NotAfter, userId)
	if err != nil {
		return err
	}

	context.Response().Header().Add("Location", location(context, certificateId))
	context.Response().WriteHeader(http.StatusCreated)
	return nil
}

func (api *CertificateAPI) deleteCertificate(context echo.Context) error {
	certificateId, err := strconv.Atoi(context.Param("certificateId"))
	if err != nil {
		fmt.Println("Could not parse certificateId:", err)
		return err
	}

	if err := api.certificateDAO.DeleteCertificate(certificateId); err != nil {
		return err
	}
	context.Response().WriteHeader(http.StatusOK)
	return nil
}
