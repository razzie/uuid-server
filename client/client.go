package client

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

// Client is a lightweight http client to request location data from uuid-server
type Client struct {
	ServerAddress string
}

// DefaultClient is the default client
var DefaultClient = *NewClient()

// NewClient returns a new client
func NewClient() *Client {
	return &Client{ServerAddress: "https://uuid.gorzsony.com"}
}

// GetUUID requests a new random UUID from uuid-server
func (c *Client) GetUUID(ctx context.Context) (uuid.UUID, error) {
	req, _ := http.NewRequest("GET", c.ServerAddress, nil)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return uuid.Nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.Parse(string(result))
}
