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

type ConsumerTypeAPI struct {
	consumerTypeDAO *db.ConsumerTypeDAO
}

func NewConsumerTypeAPI(consumerTypeDAO *db.ConsumerTypeDAO) *ConsumerTypeAPI {
	return &ConsumerTypeAPI{consumerTypeDAO}
}

func (api *ConsumerTypeAPI) Register(group *echo.Group) {
	group.GET("", api.listConsumerTypes)
	group.POST("", api.createConsumerType)
	group.GET("/:typeId", api.getConsumerType)
	group.PUT("/:typeId", api.updateConsumerType)
	group.DELETE("/:typeId", api.deleteConsumerType)
}

func (api *ConsumerTypeAPI) listConsumerTypes(context echo.Context) error {
	consumerTypes, err := api.consumerTypeDAO.List()
	if err != nil {
		return err
	}
	context.JSON(http.StatusOK, consumerTypes)
	return nil
}

func (api *ConsumerTypeAPI) createConsumerType(context echo.Context) error {
	var data db.ConsumerTypeData
	if err := context.Bind(&data); err != nil {
		return err
	}

	typeId, err := api.consumerTypeDAO.Create(data)
	if err != nil {
		return err
	}

	context.Response().Header().Add("Location", location(context, typeId))
	context.Response().WriteHeader(http.StatusCreated)
	return nil
}

func (api *ConsumerTypeAPI) getConsumerType(context echo.Context) error {
	typeId, err := strconv.Atoi(context.Param("typeId"))
	if err != nil {
		fmt.Println("Could not parse typeId:", err)
		return err
	}

	consumerType, err := api.consumerTypeDAO.Get(typeId)
	if err != nil {
		return err
	}
	context.JSON(http.StatusOK, consumerType)
	return nil
}

func (api *ConsumerTypeAPI) updateConsumerType(context echo.Context) error {
	typeId, err := strconv.Atoi(context.Param("typeId"))
	if err != nil {
		fmt.Println("Could not parse typeId:", err)
		return err
	}

	var data db.ConsumerTypeData
	if err := context.Bind(&data); err != nil {
		return err
	}

	if err := api.consumerTypeDAO.Update(typeId, data); err != nil {
		return err
	}
	context.Response().WriteHeader(http.StatusOK)
	return nil
}

func (api *ConsumerTypeAPI) deleteConsumerType(context echo.Context) error {
	typeId, err := strconv.Atoi(context.Param("typeId"))
	if err != nil {
		fmt.Println("Could not parse typeId:", err)
		return err
	}

	if err := api.consumerTypeDAO.Delete(typeId); err != nil {
		return err
	}
	context.Response().WriteHeader(http.StatusOK)
	return nil
}
