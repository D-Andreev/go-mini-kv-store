package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func MakeRequest(method, url string, body interface{}) (*http.Response, error) {
	postBody, _ := json.Marshal(body)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
