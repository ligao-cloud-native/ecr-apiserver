package v1

import (
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/rest"
)

type ProjectGetter interface {
	Projects() ProjectInterface
}

type ProjectInterface interface {
	Get()
	List()
	Create()
	Delete()
}

type projects struct {
	client rest.Interface
}

func newProjects(c *CoreV1Client) *projects {
	return &projects{
		client: c.RestClient(),
	}
}

func (c *projects) Get() {
	//var project v1.Project
	//resp, _, errs := c.client.
	//	Request(gorequest.GET, "/projects").
	//	EndStruct(&project)
}

func (c *projects) List() {}

func (c *projects) Create() {}

func (c *projects) Delete() {}
