package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Post(url string, data interface{}) (body []byte, err error) {
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
