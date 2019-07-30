package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"pkims.io/pkims/pkg/db"
	"pkims.io/pkims/pkg/pki"

	"github.com/labstack/echo/v4"
)

type CommonAuthorityAPI struct {
	generator          *pki.KeyGenerator
	commonAuthorityDAO *db.CommonAuthorityDAO
}

type Subject struct {
	CountryCode        string `json:"countryCode"`
	State              string `json:"state"`
	Locality           string `json:"locality"`
	StreetAddress      string `json:"streetAddress"`
	PostalCode         string `json:"postalCode"`
	Organization       string `json:"organization"`
	OrganizationalUnit string `json:"organizationalUnit"`
}

type CertificateData struct {
	CommonName string  `json:"commonName"`
	Subject    Subject `json:"subject"`
	ValidFor   int     `json:"validFor"`
	KeySize    int     `json:"keySize"`
}

type CommonAuthorityData struct {
	CommonAuthorityId int      `json:"commonAuthorityId"`
	Name              string   `json:"name"`
	CreatedBy         UserInfo `json:"createdBy"`
	CertificateData   string   `json:"privateKeyData"`
}

type UserInfo struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CommonAuthorityPrivateData struct {
	CommonAuthorityId int    `json:"commonAuthorityId"`
	PrivateKeyData    string `json:"privateKeyData"`
}

func NewCommonAuthorityAPI(generator *pki.KeyGenerator, commonAuthorityDAO *db.CommonAuthorityDAO) *CommonAuthorityAPI {
	return &CommonAuthorityAPI{generator, commonAuthorityDAO}
}

func (api *CommonAuthorityAPI) Register(group *echo.Group) {
	group.POST("", api.createCommonAuthority)
	group.GET("/:commonAuthorityId", api.getCommonAuthority)
	group.GET("/:commonAuthorityId/private-key", api.getCommonAuthorityPrivateKey)
}

func (api *CommonAuthorityAPI) createCommonAuthority(context echo.Context) error {
	userId := userId(context)
	var data CertificateData
	if err := context.Bind(&data); err != nil {
		return err
	}

	certificateData := pki.Certificate{
		CountryCode:        data.Subject.CountryCode,
		State:              data.Subject.State,
		Locality:           data.Subject.Locality,
		StreetAddress:      data.Subject.StreetAddress,
		PostalCode:         data.Subject.PostalCode,
		Organization:       data.Subject.Organization,
		OrganizationalUnit: data.Subject.OrganizationalUnit,
		CommonName:         data.CommonName,
		ValidFor:           data.ValidFor,
		KeySize:            data.KeySize,
	}

	privateKeyBytes, certificateBytes, err := api.generator.CreateCommonAuthority(certificateData)
	if err != nil {
		return err
	}

	commonAuthorityId, err := api.commonAuthorityDAO.SaveCommonAuthority(data.CommonName, privateKeyBytes, certificateBytes, userId)
	if err != nil {
		return err
	}

	context.Response().Header().Add("Location", location(context, commonAuthorityId))
	context.Response().WriteHeader(http.StatusCreated)

	return nil
}

func (api *CommonAuthorityAPI) getCommonAuthority(context echo.Context) error {
	commonAuthorityId, err := strconv.Atoi(context.Param("commonAuthorityId"))
	if err != nil {
		fmt.Println("Could not parse commonAuthorityId:", err)
		return err
	}

	commonAuthority, err := api.commonAuthorityDAO.LoadCommonAuthority(commonAuthorityId)
	if err != nil {
		return err
	}

	certificateData, err := api.generator.CertificateToPem(commonAuthority)
	if err != nil {
		return err
	}

	response := CommonAuthorityData{
		CommonAuthorityId: commonAuthority.CommonAuthorityId,
		Name:              commonAuthority.Name,
		CertificateData:   certificateData,
		CreatedBy: UserInfo{
			Username:  commonAuthority.CreatedBy.Username,
			FirstName: commonAuthority.CreatedBy.FirstName,
			LastName:  commonAuthority.CreatedBy.LastName,
		},
	}

	return context.JSON(http.StatusOK, response)
}

func (api *CommonAuthorityAPI) getCommonAuthorityPrivateKey(context echo.Context) error {
	commonAuthorityId, err := strconv.Atoi(context.Param("commonAuthorityId"))
	if err != nil {
		fmt.Println("Could not parse commonAuthorityId:", err)
		return err
	}

	commonAuthority, err := api.commonAuthorityDAO.LoadCommonAuthority(commonAuthorityId)
	if err != nil {
		return err
	}

	privateKeyData, err := api.generator.PrivateKeyToPem(commonAuthority)
	if err != nil {
		return err
	}

	response := CommonAuthorityPrivateData{
		CommonAuthorityId: commonAuthority.CommonAuthorityId,
		PrivateKeyData:    privateKeyData,
	}

	return context.JSON(http.StatusOK, response)
}