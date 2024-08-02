package user

import (
	"github.com/gin-gonic/gin"
	"go-to-chat/app/exception"
	"go-to-chat/app/model"
	"go-to-chat/app/utility"
	"log"
	"net/http"
	"strconv"
)

var userService = NewUserService(NewUserRepository())

func GetUser(c *gin.Context) {
	if requestUser, existed := c.Get("user-info"); existed {
		log.Println("Request User: ", requestUser)
	} else {
		utility.SendErrorResponse(c, http.StatusUnauthorized, "exception", "Unauthorized")
	}

	if c.Param("id") == "" {
		email, _ := c.Get("user-info")
		user, _ := userService.GetUserByEmail(email.(string))
		utility.SendSuccessResponse(c, http.StatusOK, "success", NewUserResponse(user))
		return
	}

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utility.NotifyError(c, err)
		return
	}

	var user *model.User
	user, err = userService.GetUser(userId)

	if err != nil {
		utility.NotifyError(c, err)
		return
	}

	utility.SendSuccessResponse(c, http.StatusOK, "success", NewUserResponse(user))
}

func CreateUser(c *gin.Context) {
	var body CreateUserBody
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

	utility.SendSuccessResponse(c, http.StatusCreated, "success", NewUserResponse(newUser))
}

func UpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utility.NotifyError(c, exception.NewBadRequestError(err.Error()))
		return
	}

	var body UpdateUserBody
	err = c.BindJSON(&body)

	if err != nil {
		utility.NotifyError(c, exception.NewBadRequestError(err.Error()))
		return
	}

	updatedUser, err := userService.UpdateUser(userId, &body)

	if err != nil {
		utility.NotifyError(c, exception.NewBadRequestError(err.Error()))
		return
	}

	utility.SendSuccessResponse(c, http.StatusOK, "success", NewUserResponse(updatedUser))
}

func UploadProfileImage(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utility.NotifyError(c, exception.NewBadRequestError(err.Error()))
		return
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		utility.NotifyError(c, exception.NewBadRequestError(err.Error()))
		return
	}

	err = userService.UploadProfileImage(userId, fileHeader)
	if err != nil {
		utility.NotifyError(c, exception.NewBadRequestError(err.Error()))
		return
	}

	utility.SendSuccessResponse(c, http.StatusOK, "success", nil)
}

func GetProfileImage(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	fileName, existed := c.GetQuery("file")

	if err != nil || !existed {
		utility.NotifyError(c, exception.NewBadRequestError(err.Error()))
		return
	}

	imagePath, err := userService.GetProfileImage(userId, fileName)

	if err != nil {
		utility.NotifyError(c, err)
		return
	}

	c.File(imagePath)
}
