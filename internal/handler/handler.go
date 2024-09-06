package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"concurrency-server/internal/common"
	"concurrency-server/internal/models"
	"concurrency-server/internal/service"

	"github.com/gin-gonic/gin"
)

var QueueHandler *queueHandler

type queueHandler struct {
	queueService service.QueueServiceInterface
}

func init() {
	QueueHandler = newQueueHandler(service.QueueService)
}

func newQueueHandler(queueService service.QueueServiceInterface) *queueHandler {
	return &queueHandler{queueService: queueService}
}

func (qh *queueHandler) GenerateQueue(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": common.INTERNAL_SERVER_ERROR_MESSAGE})
		return
	}

	res, err := qh.queueService.GenerateQueue(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": common.INTERNAL_SERVER_ERROR_MESSAGE})
		return
	}

	c.JSON(http.StatusAccepted, res)
}

func (qh *queueHandler) GetQueue(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": common.INTERNAL_SERVER_ERROR_MESSAGE})
		return
	}

	var queue models.Queue
	err = json.Unmarshal(body, &queue)
	if err != nil {
		log.Printf("Error unmarshaling JSON into Queue struct: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}

	res, err := qh.queueService.GetQueue(&queue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": common.INTERNAL_SERVER_ERROR_MESSAGE})
		return
	}

	c.JSON(http.StatusAccepted, res)
}
