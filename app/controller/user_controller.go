package controller

import (
	"github.com/gin-gonic/gin"
	"go-to-chat/app/dto/response"
	"go-to-chat/app/model"
	userService "go-to-chat/app/service"
	"go-to-chat/app/utility"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	var user *model.User
	user, err = userService.GetUser(userId)

	if err != nil {
		c.Error(err)
		return
	}

	utility.SendSuccessResponse(c, http.StatusOK, "success", response.NewUserResponse(user))
}

func CreateUser(c *gin.Context) {
	var body userService.CreateUserBody
	err := c.BindJSON(&body)

	if err != nil {
		utility.SendErrorResponse(c, http.StatusBadRequest, "exception", err.Error())
		return
	}

	newUser, err := userService.CreateUser(&body)

	if err != nil {
		utility.SendErrorResponse(c, http.StatusBadRequest, "exception", err.Error())
		return
	}

	utility.SendSuccessResponse(c, http.StatusCreated, "success", response.NewUserResponse(newUser))
}

func UpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utility.SendErrorResponse(c, http.StatusBadRequest, "exception", err.Error())
		return
	}

	var body userService.UpdateUserBody
	err = c.BindJSON(&body)

	if err != nil {
		utility.SendErrorResponse(c, http.StatusBadRequest, "exception", err.Error())
		return
	}

	updatedUser, err := userService.UpdateUser(userId, &body)

	if err != nil {
		utility.SendErrorResponse(c, http.StatusBadRequest, "exception", err.Error())
		return
	}

	utility.SendSuccessResponse(c, http.StatusOK, "success", response.NewUserResponse(updatedUser))
}
