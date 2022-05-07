package pkg

import (
	"os"

	"github.com/hashicorp/go-hclog"
)

func (c *Client) CreateFolder(f string) error {
	var (
		log hclog.Logger = c.Logger.Named("CreateFolder")
		err error
	)

	if err = os.Mkdir(f, 0755); err != nil {
		log.With("error", err, "folder", f).Error("failed to create folder")
		return err
	}

	return err
}
