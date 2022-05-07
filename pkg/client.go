package pkg

import "github.com/hashicorp/go-hclog"

func NewClient(logger hclog.Logger) *Client {
	return &Client{
		Logger: logger.Named("pkg"),
	}
}
