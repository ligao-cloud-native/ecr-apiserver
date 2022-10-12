package v2

import "github.com/ligao-cloud-native/ecr-apiserver/pkg/rest"

type CoreV2Interface interface {
	//Client() rest.Interface
	ProjectGetter
}

type CoreV2Client struct {
	restClient rest.Interface
}

func NewV2Client(c *rest.Config) (*CoreV2Client, error) {
	config := *c
	setConfigDefaults(&config)
	client, err := rest.NewRestClient(&config)
	if err != nil {
		return nil, err
	}

	return &CoreV2Client{client}, nil
}

func (c *CoreV2Client) Projects() ProjectInterface {
	return newProjects(c)
}

func (c *CoreV2Client) RestClient() rest.Interface {
	if c == nil {
		return nil
	}

	return c.restClient
}

func setConfigDefaults(config *rest.Config) {
	config.APIPath = "/api/v2.0"
	if config.UserAgent == "" {
		config.UserAgent = "harbor/v2"
	}
}
