package pkg

import "github.com/hashicorp/go-hclog"

type Project struct {
	Title      string    `yaml:"title,omitempty"`
	Subfolders []Project `yaml:"sub_folders,omitempty"`
}

type Client struct {
	Logger hclog.Logger
}
