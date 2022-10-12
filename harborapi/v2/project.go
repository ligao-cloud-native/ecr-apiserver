package v2

import (
	"fmt"
	"github.com/ligao-cloud-native/ecr-apiserver/harborapi/types/meta"
	typesv2 "github.com/ligao-cloud-native/ecr-apiserver/harborapi/types/v2"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/rest"
	"github.com/parnurzeal/gorequest"
)

type ProjectGetter interface {
	Projects() ProjectInterface
}

type ProjectInterface interface {
	Get(name string)
	List(opts meta.ListOptions) (*[]typesv2.ProjectResponse, error)
	Create(project *typesv2.Project) error
	Delete()
}

type projects struct {
	client rest.Interface
}

func newProjects(c *CoreV2Client) *projects {
	return &projects{
		client: c.RestClient(),
	}
}

func (c *projects) Get(name string) {
	//var project v1.Project
	//resp, _, errs := c.client.
	//	Request(gorequest.GET, "/projects").
	//	EndStruct(&project)
}

func (c *projects) List(opts meta.ListOptions) (*[]typesv2.ProjectResponse, error) {
	var projectList []typesv2.ProjectResponse
	resp, _, errs := c.client.
		Request(gorequest.GET, "/projects").
		Query(opts).
		EndStruct(&projectList)
	projectRes := c.client.Response(resp, &errs)
	if projectRes != nil {
		return &projectList, nil
	}

	return nil, fmt.Errorf("%v", errs)
}

func (c *projects) Create(project *typesv2.Project) error {
	resp, _, errs := c.client.
		Request(gorequest.POST, "/projects").
		Send(project).
		End()

	projectRes := c.client.Response(resp, &errs)
	if projectRes == nil {
		return fmt.Errorf("%v", errs)
	}

	return nil
}

func (c *projects) Delete() {}
