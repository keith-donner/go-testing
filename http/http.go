package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)

type HttpResponse struct {
	StatusCode int                 `json:"status_code"`
	Status     string              `json:"status"`
	Headers    map[string][]string `json:"headers"`
	Body       string              `json:"body"`
	Data       interface{}
}

func Get(url string, headers map[string]string, dataStruct interface{}, contentType string) (*HttpResponse, error) {
	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the content type header
	if contentType == "" {
		contentType = "application/json; charset=utf-8"
	}
	req.Header.Set("Content-Type", contentType)

	// Add other headers to the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Extract and return the HTTPResponse
	jsonResponse, err := getHttpResponse(resp, dataStruct)

	return jsonResponse, err
}

func Post(url string, headers map[string]string, dataStruct interface{}, contentType string, body []byte) (*HttpResponse, error) {
	// Create a new POST request with the provided body
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the content type header
	req.Header.Set("Content-Type", contentType)

	// Add other headers to the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Extract and return the HTTPResponse
	jsonResponse, err := getHttpResponse(resp, dataStruct)

	// Extract and return the HTTPResponse
	return jsonResponse, err
}

func getHttpResponse(resp *http.Response, dataStruct interface{}) (*HttpResponse, error) {
	if resp == nil {
		return nil, fmt.Errorf("response is nil")
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	body := string(bodyBytes)

	// Create a map to hold the headers
	headers := make(map[string][]string)
	for key, values := range resp.Header {
		headers[key] = values
	}

	// Create a new instance of the dataStruct type
	dataValue := reflect.New(reflect.TypeOf(dataStruct)).Interface()
	err = json.Unmarshal(bodyBytes, dataValue)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Return the custom HTTPResponse struct
	return &HttpResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Headers:    headers,
		Body:       body,
		Data:       dataValue,
	}, nil
}

func JsonToStruct(url string, body []byte, responseStruct interface{}) error {
	// Create a new POST request with the provided body
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the JSON into the provided struct
	err = json.Unmarshal(bodyBytes, responseStruct)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}
