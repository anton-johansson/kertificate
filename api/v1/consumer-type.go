package v1

import (
	"fmt"
	"net/http"
	"pkims/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ConsumerTypeAPI struct {
	consumerTypeDAO *db.ConsumerTypeDAO
}

func NewConsumerTypeAPI(consumerTypeDAO *db.ConsumerTypeDAO) *ConsumerTypeAPI {
	return &ConsumerTypeAPI{consumerTypeDAO}
}

func (api *ConsumerTypeAPI) Get(context echo.Context) error {
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

func (api *ConsumerTypeAPI) List(context echo.Context) error {
	consumerTypes, err := api.consumerTypeDAO.List()
	if err != nil {
		return err
	}
	context.JSON(http.StatusOK, consumerTypes)
	return nil
}

func (api *ConsumerTypeAPI) Update(context echo.Context) error {
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

func (api *ConsumerTypeAPI) Delete(context echo.Context) error {
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

func (api *ConsumerTypeAPI) Create(context echo.Context) error {
	var data db.ConsumerTypeData
	if err := context.Bind(&data); err != nil {
		return err
	}

	typeId, err := api.consumerTypeDAO.Create(data)
	if err != nil {
		return err
	}

	location := context.Request().RequestURI + "/" + strconv.Itoa(typeId)
	context.Response().Header().Add("Location", location)
	context.Response().WriteHeader(http.StatusOK)
	return nil
}
