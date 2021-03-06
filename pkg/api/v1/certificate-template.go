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

	"kertificate.io/kertificate/pkg/db"

	"github.com/labstack/echo/v4"
)

type CertificateTemplateAPI struct {
	certificateTemplateDAO *db.CertificateTemplateDAO
}

func NewCertificateTemplateAPI(certificateTemplateDAO *db.CertificateTemplateDAO) *CertificateTemplateAPI {
	return &CertificateTemplateAPI{certificateTemplateDAO}
}

func (api *CertificateTemplateAPI) Register(group *echo.Group) {
	group.GET("", api.listCertificateTypes)
	group.POST("", api.createCertificateType)
	group.GET("/:templateId", api.getCertificateType)
	group.PUT("/:templateId", api.updateCertificateType)
	group.DELETE("/:templateId", api.deleteCertificateType)
}

func (api *CertificateTemplateAPI) listCertificateTypes(context echo.Context) error {
	certificateTemplates, err := api.certificateTemplateDAO.List()
	if err != nil {
		return err
	}
	context.JSON(http.StatusOK, certificateTemplates)
	return nil
}

func (api *CertificateTemplateAPI) createCertificateType(context echo.Context) error {
	var data db.CertificateTemplateData
	if err := context.Bind(&data); err != nil {
		return err
	}

	templateId, err := api.certificateTemplateDAO.Create(data)
	if err != nil {
		return err
	}

	context.Response().Header().Add("Location", location(context, templateId))
	context.Response().WriteHeader(http.StatusCreated)
	return nil
}

func (api *CertificateTemplateAPI) getCertificateType(context echo.Context) error {
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

func (api *CertificateTemplateAPI) updateCertificateType(context echo.Context) error {
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

func (api *CertificateTemplateAPI) deleteCertificateType(context echo.Context) error {
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
