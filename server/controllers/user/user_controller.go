package user

import (
	"net/http"
	"ticket-system/common/response"
	"ticket-system/models/user"
	userService "ticket-system/services/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleInitUser(c *gin.Context) {

	var user user.User

	user.Username = "guess-" + uuid.New().String()
	user.IpAddress = c.ClientIP()
	user.BrowserId = c.GetHeader("User-Agent")

	if err := userService.Create(&user); err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK)
}
