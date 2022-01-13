package controller

import (
	"my-go-app/model"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	common "my-go-app/common"
	utils "my-go-app/utils"
)

func AddMenu(c echo.Context) error {
	menuObj := model.Menu{}
	menuObj.Id = uuid.NewString()
	menuObj.CreatedBy = "go-backend"
	time.Now().Local().Zone()
	menuObj.CreatedDate = time.Now()
	err := c.Bind(&menuObj)
	db := common.GetInstanceDb().Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		}
	}()

	errCreate := db.Preload("MenuType").Create(&menuObj).Error
	if errCreate != nil {
		return utils.SendResponse(c, utils.ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	menuType := model.MenuType{}
	menuType.MenuTypeCd = menuObj.MenuTypeCd
	db.Find(&menuType)
	menuObj.MenuType = menuType

	return utils.SendResponse(c, utils.ResponseMessage{
		Code:          http.StatusOK,
		Message:       "success",
		Data:          menuObj,
		CorrelationID: "",
	})
}

func GetMenus(c echo.Context) error {
	db := common.GetInstanceDb().Begin()
	var u []*model.Menu

	if err := db.Preload("MenuType").Find(&u).Error; err != nil {
		return utils.SendResponse(c, utils.ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return utils.SendResponse(c, utils.ResponseMessage{
		Code:          http.StatusOK,
		Message:       "success",
		Data:          u,
		CorrelationID: "",
	})
}
