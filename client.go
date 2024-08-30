package mepost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client represents the client for the Mepost API.
type Client struct {
	APIKey  string
	BaseURL string
}

// NewClient creates a new instance of MepostClient.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.mepost.io/v1",
	}
}

// AddDomain adds a domain to the Mepost account.
func (c *Client) AddDomain(request AddDomainRequest) (*AddDomainResponse, error) {
	url := fmt.Sprintf("%s/company/domain/add", c.BaseURL)
	response := &AddDomainResponse{}
	err := c.makeRequest("POST", url, request, response)
	return response, err
}

// RemoveDomain removes a domain from the Mepost account.
func (c *Client) RemoveDomain(request RemoveDomainRequest) (*RemoveDomainResponse, error) {
	url := fmt.Sprintf("%s/company/domain/remove", c.BaseURL)
	response := &RemoveDomainResponse{}
	err := c.makeRequest("DELETE", url, request, response)
	return response, err
}

// ListGroups retrieves a list of email groups.
func (c *Client) ListGroups(limit, page int) (*BaseResult[EmailGroup], error) {
	url := fmt.Sprintf("%s/groups?limit=%d&page=%d", c.BaseURL, limit, page)
	response := &BaseResult[EmailGroup]{}
	err := c.makeRequest("GET", url, nil, response)
	return response, err
}

// CreateGroup creates a new email group.
func (c *Client) CreateGroup(request CreateNewGroupRequest) (*EmailGroup, error) {
	url := fmt.Sprintf("%s/groups", c.BaseURL)
	response := &EmailGroup{}
	err := c.makeRequest("POST", url, request, response)
	return response, err
}

// DeleteGroup deletes an email group.
func (c *Client) DeleteGroup(groupId string) (bool, error) {
	url := fmt.Sprintf("%s/groups/%s", c.BaseURL, groupId)
	var response bool
	err := c.makeRequest("DELETE", url, nil, &response)
	return response, err
}

// GetGroupById retrieves details of a specific email group.
func (c *Client) GetGroupById(groupId string) (*EmailGroupWithCounts, error) {
	url := fmt.Sprintf("%s/groups/%s", c.BaseURL, groupId)
	response := &EmailGroupWithCounts{}
	err := c.makeRequest("GET", url, nil, response)
	return response, err
}

// UpdateGroup updates the name of an email group.
func (c *Client) UpdateGroup(groupId string, request RenameGroupRequest) (bool, error) {
	url := fmt.Sprintf("%s/groups/%s", c.BaseURL, groupId)
	var response bool
	err := c.makeRequest("PUT", url, request, &response)
	return response, err
}

// ListSubscribers retrieves a list of subscribers in a group.
func (c *Client) ListSubscribers(groupId string, limit, page int) (*BaseResult[Subscriber], error) {
	url := fmt.Sprintf("%s/groups/%s/subscribers?limit=%d&page=%d", c.BaseURL, groupId, limit, page)
	response := &BaseResult[Subscriber]{}
	err := c.makeRequest("GET", url, nil, response)
	return response, err
}

// AddSubscriber adds a new subscriber to a group.
func (c *Client) AddSubscriber(groupId string, request CreateSubscriberRequest) (bool, error) {
	url := fmt.Sprintf("%s/groups/%s/subscribers", c.BaseURL, groupId)
	var response bool
	err := c.makeRequest("POST", url, request, &response)
	return response, err
}

// DeleteSubscriber removes a subscriber from a group.
func (c *Client) DeleteSubscriber(groupId string, request DeleteSubscriberRequest) (bool, error) {
	url := fmt.Sprintf("%s/groups/%s/subscribers", c.BaseURL, groupId)
	var response bool
	err := c.makeRequest("DELETE", url, request, &response)
	return response, err
}

// GetSubscriberByEmail retrieves a subscriber's details by email.
func (c *Client) GetSubscriberByEmail(groupId, email string) (*Subscriber, error) {
	url := fmt.Sprintf("%s/groups/%s/subscribers/%s", c.BaseURL, groupId, email)
	response := &Subscriber{}
	err := c.makeRequest("GET", url, nil, response)
	return response, err
}

// SendMarketing sends a marketing email.
func (c *Client) SendMarketing(request SendMarketingRequest) (*Schedule, error) {
	url := fmt.Sprintf("%s/messages/marketing", c.BaseURL)
	response := &Schedule{}
	err := c.makeRequest("POST", url, request, response)
	return response, err
}

// SendMessageByTemplate sends a message using a specified template.
func (c *Client) SendMessageByTemplate(request SendMessageByTemplateRequest) (*Schedule, error) {
	url := fmt.Sprintf("%s/messages/marketing-by-template", c.BaseURL)
	response := &Schedule{}
	err := c.makeRequest("POST", url, request, response)
	return response, err
}

// SendTransactional sends a transactional email.
func (c *Client) SendTransactional(request SendTransactionalRequest) (*Schedule, error) {
	url := fmt.Sprintf("%s/messages/transactional", c.BaseURL)
	response := &Schedule{}
	err := c.makeRequest("POST", url, request, response)
	return response, err
}

// SendTransactionalByTemplate sends a transactional email using a template.
func (c *Client) SendTransactionalByTemplate(request SendMessageByTemplateRequest) (*Schedule, error) {
	url := fmt.Sprintf("%s/messages/transactional-by-template", c.BaseURL)
	response := &Schedule{}
	err := c.makeRequest("POST", url, request, response)
	return response, err
}

// CreateIpGroup creates a new IP group.
func (c *Client) CreateIpGroup(request CreateIpGroupRequest) (*IPGroup, error) {
	url := fmt.Sprintf("%s/outbound/ip-group/create", c.BaseURL)
	response := &IPGroup{}
	err := c.makeRequest("POST", url, request, response)
	return response, err
}

// GetIpGroupInfo retrieves information about a specific IP group.
func (c *Client) GetIpGroupInfo(name string) (*IPGroup, error) {
	url := fmt.Sprintf("%s/outbound/ip-group/info/%s", c.BaseURL, name)
	response := &IPGroup{}
	err := c.makeRequest("GET", url, nil, response)
	return response, err
}

// ListIpGroups retrieves a list of IP groups.
func (c *Client) ListIpGroups() ([]IPGroup, error) {
	url := fmt.Sprintf("%s/outbound/ip-group/list", c.BaseURL)
	response := []IPGroup{}
	err := c.makeRequest("GET", url, nil, &response)
	return response, err
}

// CancelWarmup cancels an IP warm-up process.
func (c *Client) CancelWarmup(request CancelWarmUpRequest) (*CancelWarmUpResponse, error) {
	url := fmt.Sprintf("%s/outbound/ip/cancel-warmup", c.BaseURL)
	response := &CancelWarmUpResponse{}
	err := c.makeRequest("POST", url, request, response)
	return response, err
}

// GetIpInfo retrieves information about an IP address.
func (c *Client) GetIpInfo(ip string) (*IpAddress, error) {
	url := fmt.Sprintf("%s/outbound/ip/info/%s", c.BaseURL, ip)
	response := &IpAddress{}
	err := c.makeRequest("GET", url, nil, response)
	return response, err
}

// ListIps retrieves a list of IP addresses.
func (c *Client) ListIps() ([]IpAddress, error) {
	url := fmt.Sprintf("%s/outbound/ip/list", c.BaseURL)
	response := []IpAddress{}
	err := c.makeRequest("GET", url, nil, &response)
	return response, err
}

// SetIpGroup assigns an IP address to a group.
func (c *Client) SetIpGroup(request SetIpGroupRequest) (*SetIpGroupResponse, error) {
	url := fmt.Sprintf("%s/outbound/ip/set-ip-group", c.BaseURL)
	response := &SetIpGroupResponse{}
	err := c.makeRequest("POST", url, request, response)
	return response, err
}

// StartWarmup starts the IP warm-up process.
func (c *Client) StartWarmup(request StartWarmUpRequest) (*StartWarmUpResponse, error) {
	url := fmt.Sprintf("%s/outbound/ip/start-warmup", c.BaseURL)
	response := &StartWarmUpResponse{}
	err := c.makeRequest("POST", url, request, response)
	return response, err
}

// makeRequest handles the HTTP requests to the Mepost API.
func (c *Client) makeRequest(method, url string, requestData interface{}, response interface{}) error {
	var jsonData []byte
	var err error

	if requestData != nil {
		jsonData, err = json.Marshal(requestData)
		if err != nil {
			return fmt.Errorf("error marshalling request data: %v", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("error unmarshalling response: %v", err)
	}

	return nil
}
