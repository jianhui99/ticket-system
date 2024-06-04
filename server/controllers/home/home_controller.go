package home

import (
	"ticket-system/common/response"

	"github.com/gin-gonic/gin"
)

func HandleGetHome(c *gin.Context) {
	response.Success(c, "Welcome to the ticket system")
}
