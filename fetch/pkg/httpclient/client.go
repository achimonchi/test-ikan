package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

type HttpClient struct {
	clientHost string
	timeout    time.Duration
}

func NewHttpClient(clientHost, clientPort string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		clientHost: fmt.Sprintf("%s:%s", clientHost, clientPort),
		timeout:    timeout,
	}
}

func (c *HttpClient) dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, c.timeout)
}

func (c *HttpClient) getClient() *http.Client {
	transport := http.Transport{
		Dial: c.dialTimeout,
	}

	client := http.Client{
		Transport: &transport,
		Timeout:   c.timeout,
	}

	return &client
}

func (c *HttpClient) Get(path string, headers map[string]string) (map[string]interface{}, error) {
	client := c.getClient()

	resp, err := client.Get(fmt.Sprintf("%s/%s", c.clientHost, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("failed to access %s/%s", c.clientHost, path)
		return nil, errors.New(msg)
	}

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, err
}

func (c *HttpClient) Post(path string, body []byte, headers map[string]string) (map[string]interface{}, error) {
	client := c.getClient()
	url := fmt.Sprintf("%s/%s", c.clientHost, path)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for key, header := range headers {
		req.Header.Set(key, header)
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var errorResponse map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to deserialize error response from %s/%s - with error %s", c.clientHost, path, err.Error())
		}

		return errorResponse, fmt.Errorf("failed to post data to %s/%s - with error %v", c.clientHost, path, errorResponse)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, err
}
