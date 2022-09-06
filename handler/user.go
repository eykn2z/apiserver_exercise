package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	"server/controller"
)

type UserHandler struct{}

func (handler UserHandler) Create(c echo.Context) error {
	name := c.FormValue("name")
	con := controller.UserController{}
	user, _ := con.Create(name)
	return c.JSON(http.StatusOK, user)
}

func (handler UserHandler) Read(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	con := controller.UserController{}
	user, err := con.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.NoContent(http.StatusNoContent)
		}
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (handler UserHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	con := controller.UserController{}
	user, err := con.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.NoContent(http.StatusNoContent)
		}
		return err
	}
	name := c.FormValue("name")
	user, err = con.Update(user, name)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (handler UserHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	con := controller.UserController{}
	user, err := con.GetByID(id)
	if err != nil {
		return err
	}
	err = con.Delete(user)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string("deleted"))
}

func (handler UserHandler) AllRead(c echo.Context) error {
	con := controller.UsersController{}
	users, err := con.Read()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.NoContent(http.StatusNoContent)
		}
		return err
	}
	return c.JSON(http.StatusOK, users)
}
