package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"exampl.com/goCrud/helpers"
	"exampl.com/goCrud/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type createUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type updateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB) //type assign
	// var a interface{}
	// b := a.(int)
	var users []model.User
	db.Find(&users)
	response := helpers.BuildRespnse(true, "Ok", users)
	c.JSON(http.StatusOK, response)
}

func CreateUser(c *gin.Context) {
	var input createUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helpers.BuildErrorResponse("Someting went wrong", err.Error(), helpers.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user := model.User{Name: input.Name, Email: input.Email}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)
	response := helpers.BuildRespnse(true, "Ok", user)
	c.JSON(http.StatusOK, response)
}

func FindUser(c *gin.Context) {
	var user model.User
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		response := helpers.BuildErrorResponse("User not found", err.Error(), helpers.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildRespnse(true, "OK", user)
	c.JSON(http.StatusOK, response)
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user model.User
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err := db.Where("id=?", id).First(&user).Error; err != nil {
		response := helpers.BuildErrorResponse("User not found", err.Error(), helpers.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var input updateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		response := helpers.BuildErrorResponse("Not know error", err.Error(), helpers.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user.Name = input.Name
	user.Email = input.Email
	db.Model(&user).Where("id = ?", id).Updates(user)
	response := helpers.BuildRespnse(true, "Ok", user)
	c.JSON(http.StatusOK, response)
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user model.User
	if err := db.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		response := helpers.BuildErrorResponse("User not found", err.Error(), helpers.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	db.Delete(&user)
	response := helpers.BuildRespnse(true, "Ok", user)
	c.JSON(http.StatusOK, response)
}
