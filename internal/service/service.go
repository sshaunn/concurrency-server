package service

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"concurrency-server/internal/common"
	"concurrency-server/internal/models"
	"concurrency-server/internal/util"
)

var QueueService QueueServiceInterface

type QueueServiceInterface interface {
	GenerateQueue(*models.Queue) (map[string]interface{}, error)
	GetQueue(*models.Queue) (map[string]interface{}, error)
}

type queueServiceImpl struct {
	baseURL  string
	endpoint string
	client   *http.Client
}

func NewQueueService(baseURL string, endpoint string) QueueServiceInterface {
	return &queueServiceImpl{
		baseURL:  baseURL,
		endpoint: endpoint,
		client:   &http.Client{},
	}
}

func init() {
	baseURL := os.Getenv(common.SERVICE_EVENTING_TOOLS_BASE_URL)
	endpoint := os.Getenv(common.SERVICE_EVENTING_TOOLS_ENDPOINT)
	QueueService = NewQueueService(baseURL, endpoint)
}

func (qsi *queueServiceImpl) GenerateQueue(queue *models.Queue) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(queue)
	if err != nil {
		log.Panicf("Failed marshalling queue with error=%v", err.Error())
		return nil, err
	}

	res, err := qsi.client.Post(qsi.baseURL, common.ContentType, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Panicf("Error making request with error=%v", err.Error())
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var resbody map[string]interface{}

	err = json.Unmarshal(body, &resbody)
	if err != nil {
		log.Panicf("Response body marshalling with error=%v", err.Error())
		return nil, err
	}
	return resbody, err
}

func (qsi *queueServiceImpl) GetQueue(queue *models.Queue) (map[string]interface{}, error) {
	return util.StructToMap(queue)
}
