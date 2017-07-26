package k8sGoRestful

import (
	"k8sGoRestful/models"
	"github.com/tedsuo/rata"
	"net/url"
	"github.com/gogo/protobuf/proto"
	"net/http"
	"bytes"
	"time"
	"io/ioutil"
)

const (
	ContentTypeHeader    = "Content-Type"
	ProtoContentType     = "application/x-protobuf"
)

type Client interface {
	CreatePods(namespace string, instance *models.Pods) (bool, error)

}

type client struct {
	reqGen    *rata.RequestGenerator
}

func newClient(url string) *client {
	return &client{
		reqGen:    rata.NewRequestGenerator(url, Routes),
	}
}

func NewClient(url string) Client {
	return newClient(url)
}

func (c *client) CreatePods(namespace string, instance *models.Pods) (bool, error) {
	request := instance
	response := models.PodsResponse{}
	err := c.doRequest(CreatePods, rata.Params{"namespace": namespace},nil, request, &response)
	if err != nil {
		return false, err
	}
	return response.ShouldStart, nil
}

func (c *client) doRequest(requestName string, params rata.Params, queryParams url.Values, requestBody, responseBody proto.Message) error {
	var err error
	var request *http.Request

	for attempts := 0; attempts < 3; attempts++ {
		request, err = c.createRequest(requestName, params, queryParams, requestBody)
		if err != nil {
			return err
		}
		err = c.do(request, responseBody)

		if err != nil {
			time.Sleep(500 * time.Millisecond)
		} else {
			break
		}
	}
	return err
}

func (c *client) do(request *http.Request, responseObject proto.Message) error {
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		// don't worry about errors when closing the body
		_ = response.Body.Close()
	}()
	return handleProtoResponse(response, responseObject)
}

func (c *client) createRequest(requestName string, params rata.Params, queryParams url.Values, message proto.Message) (*http.Request, error) {
	var messageBody []byte
	var err error
	if message != nil {
		messageBody, err = proto.Marshal(message)
		if err != nil {
			return nil, err
		}
	}

	request, err := c.reqGen.CreateRequest(requestName, params, bytes.NewReader(messageBody))
	if err != nil {
		return nil, err
	}

	request.URL.RawQuery = queryParams.Encode()
	request.ContentLength = int64(len(messageBody))
	request.Header.Set("Content-Type", ProtoContentType)
	return request, nil
}

func handleProtoResponse(response *http.Response, responseObject proto.Message) error {
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(buf, responseObject)
	if err != nil {
		return err
	}
	return nil
}