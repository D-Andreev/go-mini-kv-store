package e2e_test

import (
	"encoding/json"
	"go-mini-kv-store/test/utils"
	"io"
	"net/http"
	"os"
	"testing"
)

var key = "mykey"

func TestPut(t *testing.T) {
	defer setupTest(key)()
	url := os.Getenv("URL")
	resp, err := putItem(url+"/"+key, map[string]string{"value": "value1"})
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Errorf("Error in response: %v", err.Error())
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response: %v", err.Error())
	}
	var response map[string]string
	json.Unmarshal(responseData, &response)
	if response["item"] != "value1" {
		t.Errorf("Expected value1, got %v", response["item"])
	}
}

func TestGet(t *testing.T) {
	defer setupTest(key)()
	url := os.Getenv("URL")
	_, err := putItem(url+"/"+key, map[string]string{"value": "value1"})
	if err != nil {
		t.Errorf("Error in response: %v", err.Error())
	}

	resp, err := utils.MakeRequest("GET", url+"/"+key, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Errorf("Error in response: %v", err.Error())
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response: %v", err.Error())
	}
	var response map[string]string
	json.Unmarshal(responseData, &response)
	if response["success"] != "true" {
		t.Errorf("Expected success=true, got %v", response["success"])
	}
}

func TestDelete(t *testing.T) {
	defer setupTest(key)()
	url := os.Getenv("URL")
	resp, err := utils.MakeRequest("DELETE", url+"/"+key, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Errorf("Error in response: %v", err.Error())
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response: %v", err.Error())
	}
	var response map[string]string
	json.Unmarshal(responseData, &response)
	if response["success"] != "true" {
		t.Errorf("Expected success=true, got %v", response["success"])
	}
}

func setupTest(key string) func() {
	return func() {
		url := os.Getenv("URL")
		resp, err := utils.MakeRequest("DELETE", url+"/"+key, nil)
		if err != nil || resp.StatusCode != http.StatusOK {
			panic("Error in response")
		}
	}
}

func putItem(url string, item map[string]string) (*http.Response, error) {
	resp, err := utils.MakeRequest("PUT", url, item)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return resp, nil
}
