package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
	"github.com/hashicorp/go-hclog"
	"os"
)

const MaxItemsPerPage = 200

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger hclog.Logger

	DoClient *godo.Client
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
	// Init your client and 3rd party clients using the user's configuration
	// passed by the SDK providerConfig
	// if token is not present, try from environ

	if providerConfig.Token == "" {
		providerConfig.Token = getTokenFromEnv()
	}

	client := Client{
		logger:   logger,
		DoClient: godo.NewFromToken(providerConfig.Token),
	}
	// Return the initialized client and it will be passed to your resources
	return &client, nil
}

func getTokenFromEnv() string {
	doToken := os.Getenv("DIGITALOCEAN_TOKEN")
	doAccessToken := os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")
	if doToken != "" {
		return doToken
	}
	if doAccessToken != "" {
		return doAccessToken
	}
	return ""
}
