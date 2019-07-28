package v1

import (
	"fmt"
	"net/http"
	"pkims/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CertificateTemplateAPI struct {
	certificateTemplateDAO *db.CertificateTemplateDAO
}

func NewCertificateTemplateAPI(certificateTemplateDAO *db.CertificateTemplateDAO) *CertificateTemplateAPI {
	return &CertificateTemplateAPI{certificateTemplateDAO}
}

func (api *CertificateTemplateAPI) Get(context echo.Context) error {
	templateId, err := strconv.Atoi(context.Param("templateId"))
	if err != nil {
		fmt.Println("Could not parse templateId:", err)
		return err
	}

	certificateTemplate, err := api.certificateTemplateDAO.Get(templateId)
	if err != nil {
		return err
	}
	context.JSON(http.StatusOK, certificateTemplate)
	return nil
}

func (api *CertificateTemplateAPI) List(context echo.Context) error {
	certificateTemplates, err := api.certificateTemplateDAO.List()
	if err != nil {
		return err
	}
	context.JSON(http.StatusOK, certificateTemplates)
	return nil
}

func (api *CertificateTemplateAPI) Update(context echo.Context) error {
	templateId, err := strconv.Atoi(context.Param("templateId"))
	if err != nil {
		fmt.Println("Could not parse templateId:", err)
		return err
	}

	var data db.CertificateTemplateData
	if err := context.Bind(&data); err != nil {
		return err
	}

	if err := api.certificateTemplateDAO.Update(templateId, data); err != nil {
		return err
	}
	context.Response().WriteHeader(http.StatusOK)
	return nil
}

func (api *CertificateTemplateAPI) Delete(context echo.Context) error {
	templateId, err := strconv.Atoi(context.Param("templateId"))
	if err != nil {
		fmt.Println("Could not parse templateId:", err)
		return err
	}

	if err := api.certificateTemplateDAO.Delete(templateId); err != nil {
		return err
	}
	context.Response().WriteHeader(http.StatusOK)
	return nil
}

func (api *CertificateTemplateAPI) Create(context echo.Context) error {
	var data db.CertificateTemplateData
	if err := context.Bind(&data); err != nil {
		return err
	}

	templateId, err := api.certificateTemplateDAO.Create(data)
	if err != nil {
		return err
	}
	context.Response().Header().Add("Location", "/v1/certificate-template/"+strconv.Itoa(templateId))
	context.Response().WriteHeader(http.StatusOK)
	return nil
}
