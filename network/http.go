package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
)

func Post(data interface{}, url string) (body []byte, err error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status code %v", resp.StatusCode)
		return
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func Get(url string) (body []byte, err error) {
	log.Infof("HTTP GET: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status code %v", resp.StatusCode)
		return
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
