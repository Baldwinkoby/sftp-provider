package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// AuthResponse
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

// client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// NewClient -
func NewClient(SftpGoHost string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		HostURL:    *host,
	}
	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func (c *Client) CreateUser(user User) (*User, error) {
	rb, err := json.Marshal(UserCreate{
		Name:             user.Name,
		Email:            user.Email,
		OrganizationName: user.OrganizationName,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/user", c.PritunlWrapperHost), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// body, err := c.doRequest(req)
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	createdUser := User{}
	err = json.Unmarshal(body, &createdUser)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}
