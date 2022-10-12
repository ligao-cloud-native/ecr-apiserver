package v1

import "github.com/ligao-cloud-native/ecr-apiserver/pkg/rest"

type CoreV1Interface interface {
	//Client() rest.Interface
	ProjectGetter
}

type CoreV1Client struct {
	restClient rest.Interface
}

func NewV1Client(c *rest.Config) (*CoreV1Client, error) {
	config := *c
	setConfigDefaults(&config)
	client, err := rest.NewRestClient(&config)
	if err != nil {
		return nil, err
	}

	return &CoreV1Client{client}, nil
}

func (c *CoreV1Client) Projects() ProjectInterface {
	return newProjects(c)
}

func (c *CoreV1Client) RestClient() rest.Interface {
	if c == nil {
		return nil
	}

	return c.restClient
}

func setConfigDefaults(config *rest.Config) {
	config.APIPath = "/api"
	if config.UserAgent == "" {
		config.UserAgent = "harbor/v1"
	}
}
