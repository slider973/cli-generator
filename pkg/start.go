package pkg

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
)

func (c *Client) CreateStruct(p Project) error {
	var (
		log hclog.Logger = c.Logger.Named("CreateStruct")
		err error
	)

	if err = c.CreateFolder(p.Title); err != nil {
		return err
	}

	for _, v := range p.Subfolders {
		v.Title = fmt.Sprintf("%s/%s", p.Title, v.Title)
		if err = c.CreateStruct(v); err != nil {
			log.With("error", err, "project", v.Title).Error("failed to create struct")
			continue
		}
	}

	return nil
}

func (c *Client) Start(projects []Project) {
	var (
		log hclog.Logger = c.Logger.Named("Start")
		err error
	)

	for _, v := range projects {
		// Create Folder
		if err = c.CreateStruct(v); err != nil {
			log.With("error", err).Error("failed to create struct")
			continue
		}
	}
}
