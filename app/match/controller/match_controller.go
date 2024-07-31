package controller

import (
	"github.com/gin-gonic/gin"
	"go-to-chat/app/match/service"
	"go-to-chat/app/utility"
	"net/http"
)

var matchService = service.NewMatchService()

func FindMatch(c *gin.Context) {
	user := c.GetString("user-info")

	err := matchService.CreateNewMatchTask(user)

	if err != nil {
		utility.NotifyError(c, err)
		return
	}

	utility.SendSuccessResponse(c, http.StatusOK, "start match task", nil)
}
