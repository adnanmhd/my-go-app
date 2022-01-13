package main

import (
	"fmt"
	"my-go-app/model"
	"my-go-app/utils"
	"net/http"
	"strconv"
	"time"

	common "my-go-app/common"
	controller "my-go-app/controller"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/tylerb/graceful"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
	db    *gorm.DB
)

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/apps_db?parseTime=True&loc=Asia%2FJakarta"
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("Connected successfully")
}

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func getMenus(c echo.Context) error {
	db.SingularTable(true)
	var u []*model.Menu
	if err := db.Preload("MenuType").Find(&u).Error; err != nil {
		// error handling here
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func getMenuById(c echo.Context) error {
	db.SingularTable(true)
	id := c.Param("id")
	menuObj := model.Menu{}
	menuObj.Id = id
	err := db.Preload("MenuType").Find(&menuObj).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"statusCode": http.StatusNotFound,
			"message":    "Data not found",
		})
	}
	return c.JSON(http.StatusOK, menuObj)
}

func addMenu(c echo.Context) error {
	db.SingularTable(true)
	menuObj := model.Menu{}
	menuObj.Id = uuid.NewString()
	menuObj.CreatedBy = "go-backend"
	time.Now().Local().Zone()
	menuObj.CreatedDate = time.Now()
	if err := c.Bind(&menuObj); err != nil {
		return err
	}
	err := db.Preload("MenuType").Create(&menuObj).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	menuType := model.MenuType{}
	menuType.MenuTypeCd = menuObj.MenuTypeCd
	db.Find(&menuType)
	menuObj.MenuType = menuType
	return c.JSON(http.StatusOK, menuObj)
}

func deleteMenu(c echo.Context) error {
	db.SingularTable(true)
	id := c.Param("id")
	menuDel := model.Menu{}
	menuDel.Id = id
	err := db.Delete(&menuDel).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"statusCode": http.StatusOK,
		"message":    "Data successfully deleted",
	})
}

func addTransaction(c echo.Context) error {
	db.SingularTable(true)
	billTrx := model.BillTransaction{}
	billTrx.Id = uuid.NewString()
	time.Now().Local().Zone()
	billTrx.CreatedDate = time.Now()
	billTrx.BillCode = utils.GenerateBillCode()
	if err := c.Bind(&billTrx); err != nil {
		return err
	}
	err := db.Create(&billTrx).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, billTrx)
}

func getTransactions(c echo.Context) error {
	db.SingularTable(true)
	var billTrx []*model.BillTransaction
	if err := db.Preload("BillTransactionDtl").Find(&billTrx).Error; err != nil {
		// error handling here
		return err
	}
	return c.JSON(http.StatusOK, billTrx)
}

func main() {
	e := echo.New()
	common.InitConfig()
	common.GetInstanceDb()
	controller.AssignRouting(e)
	e.Server.Addr = common.Config.Port
	graceful.ListenAndServe(e.Server, 5*time.Second)
	// Start server
}
