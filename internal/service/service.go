package service

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"concurrency-server/internal/common"
	"concurrency-server/internal/models"
	"concurrency-server/internal/util"
)

var QueueService QueueServiceInterface

// counter      uint64
// backends     = []string{os.Getenv(common.SERVICE_EVENTING_TOOLS_BASE_URL_1), os.Getenv(common.SERVICE_EVENTING_TOOLS_BASE_URL_2), os.Getenv(common.SERVICE_EVENTING_TOOLS_BASE_URL_3)}

type QueueServiceInterface interface {
	GenerateQueue([]byte) (map[string]interface{}, error)
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
		client:   &http.Client{Timeout: 60000 * time.Millisecond},
	}
}

func init() {
	baseURL := os.Getenv(common.SERVICE_EVENTING_TOOLS_BASE_URL_1)
	endpoint := os.Getenv(common.SERVICE_EVENTING_TOOLS_ENDPOINT)
	QueueService = NewQueueService(baseURL, endpoint)
}

// func getNextBackend() string {
// 	return backends[int(atomic.AddUint64(&counter, 1)%uint64(len(backends)))]
// }

func (qsi *queueServiceImpl) GenerateQueue(queue []byte) (map[string]interface{}, error) {
	// backend := getNextBackend()
	res, err := qsi.client.Post(os.Getenv(common.SERVICE_EVENTING_TOOLS_BASE_URL_1)+qsi.endpoint, common.ContentType, bytes.NewBuffer(queue))
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
