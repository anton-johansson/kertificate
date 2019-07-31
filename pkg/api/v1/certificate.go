package v1

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"pkims.io/pkims/pkg/db"
	"pkims.io/pkims/pkg/pki"
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
