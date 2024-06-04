package queue

import (
	"net/http"
	"ticket-system/common/response"
	"ticket-system/models/queue"
	queueService "ticket-system/services/queue"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleInitQueue(c *gin.Context) {

	var queue queue.Queue

	queue.QueueId = uuid.New().String()
	queue.IpAddress = c.ClientIP()
	queue.BrowserId = c.GetHeader("User-Agent")
	queue.EnteredAt = time.Now()

	if err := queueService.Create(&queue); err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK)
}
