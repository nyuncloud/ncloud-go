package core

import (
	"github.com/nyuncloud/ncloud-go/core/net/consts"
	"time"
)

type Config struct {
	Scheme   string
	Endpoint string
	Timeout  time.Duration
	Retry    int
}

func NewConfig() *Config {
	return &Config{
		Scheme:   consts.SchemeHttps,
		Endpoint: "www.cloud-api.com",
		Timeout:  5 * time.Second,
		Retry:    3,
	}
}

func (c *Config) SetScheme(scheme string) {
	c.Scheme = scheme
}

func (c *Config) SetEndpoint(endpoint string) {
	c.Endpoint = endpoint
}

func (c *Config) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}
