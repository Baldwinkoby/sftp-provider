package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
func NewClient(ctx context.Context, host, username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    *host,
	}

	data := []byte(fmt.Sprintf("%s:%s", *username, *password))
	token := base64.StdEncoding.EncodeToString(data)

	envToken := os.Getenv("SFTPGO_JWT_TOKEN")
	if envToken != "" {
		c.Token = fmt.Sprintf("Bearer %s", envToken)
		return &c, nil
	}

	// get token
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/token", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", token))
	body, err := c.doRequest(req)

	// parse response body
	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	c.Token = fmt.Sprintf("Bearer %s", ar.AccessToken)
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

/*
// **********************************USER API CREATION*********************************************
*/

func (c *Client) CreateUser(ctx context.Context, admin models.User) (*models.User, error) {
	rb, err := json.Marshal(admin)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/user", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// body, err := c.doRequest(req)
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	//Check Later
	/*
		if res.StatusCode != http.StatusCreated {
			bolB, _ := json.Marshal(admin)
			return nil, fmt.Errorf("status: %d, body: %s, payload: %s", res.StatusCode, body, bolB)
	*/

	order := models.User{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) GetUser(ctx context.Context, username string) (*models.User, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/users/%s", c.HostURL, username), nil)
	if err != nil {
		return nil, err
	}

	// body, err := c.doRequest(req)
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := models.User{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) DeleteUser(ctx context.Context, username string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v2/users/%s", c.HostURL, username), nil)
	if err != nil {
		return err
	}

	// body, err := c.doRequest(req)
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateUser(ctx context.Context, username string, admin models.User) error {
	rb, err := json.Marshal(admin)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v2/users/%s", c.HostURL, username), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	order := models.User{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil
	}

	return err
}

/*
// **********************************ADMIN API CREATION*********************************************
*/

func (c *Client) GetAdmin(ctx context.Context, username string) (*models.Admin, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/admins/%s", c.HostURL, username), nil)
	if err != nil {
		return nil, err
	}

	// body, err := c.doRequest(req)
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := models.Admin{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) CreateAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error) {
	rb, err := json.Marshal(admin)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/admin", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// body, err := c.doRequest(req)
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := models.Admin{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) UpdateAdmin(ctx context.Context, username string, admin models.Admin) error {
	rb, err := json.Marshal(admin)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v2/admin/%s", c.HostURL, username), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	order := models.Admin{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil
	}

	return err
}

func (c *Client) DeleteAdmin(ctx context.Context, username string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v2/Admin/%s", c.HostURL, username), nil)
	if err != nil {
		return err
	}

	// body, err := c.doRequest(req)
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
