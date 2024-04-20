package mepost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client represents the client for the Mepost API
type Client struct {
	APIKey  string
	BaseURL string
}

// NewClient creates a new instance of MepostClient
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.mepost.io/v1",
	}
}

// SendEmail sends an email using the Mepost API
func (c *Client) SendEmail(emailData map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/messages/send", c.BaseURL)
	jsonData, err := json.Marshal(emailData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling email data: %v", err)
	}

	response, err := c.makeRequest("POST", url, jsonData)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SendEmailByTemplate sends an email using a specified template
func (c *Client) SendEmailByTemplate(emailData map[string]interface{}, templateID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/messages/send-by-template", c.BaseURL)
	data := map[string]interface{}{
		"message":    emailData,
		"templateId": templateID,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling email data: %v", err)
	}

	response, err := c.makeRequest("POST", url, jsonData)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetInfo retrieves information about a specific scheduled message
func (c *Client) GetInfo(scheduleID string, email string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/messages/%s/%s", c.BaseURL, scheduleID, email)

	response, err := c.makeRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CancelScheduledMessage cancels a scheduled message
func (c *Client) CancelScheduledMessage(scheduledMessageID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/messages/cancel-scheduled", c.BaseURL)
	data := map[string]interface{}{
		"scheduledMessageId": scheduledMessageID,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling cancellation data: %v", err)
	}

	response, err := c.makeRequest("POST", url, jsonData)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetScheduledMessage retrieves a scheduled message
func (c *Client) GetScheduledMessage(scheduleID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/messages/schedule/%s", c.BaseURL, scheduleID)

	response, err := c.makeRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// makeRequest handles the HTTP requests to the Mepost API
func (c *Client) makeRequest(method string, url string, jsonData []byte) (map[string]interface{}, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return response, nil
}
