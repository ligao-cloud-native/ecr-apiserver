package harborapi

import (
	v1 "github.com/ligao-cloud-native/ecr-apiserver/harborapi/v1"
	v2 "github.com/ligao-cloud-native/ecr-apiserver/harborapi/v2"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/rest"
	"strings"
)

type ClientSet struct {
	v1 *v1.CoreV1Client
	v2 *v2.CoreV2Client
}

func NewClient(c *rest.Config) (*ClientSet, error) {
	configShallowCopy := *c

	var cs ClientSet
	var err error
	switch {
	case strings.Contains(c.APIPath, "v2"):
		cs.v2, err = v2.NewV2Client(&configShallowCopy)
		if err != nil {
			return nil, err
		}
	default:
		cs.v1, err = v1.NewV1Client(&configShallowCopy)
		if err != nil {
			return nil, err
		}
	}

	return &cs, nil
}

func (c *ClientSet) V1() v1.CoreV1Interface {
	return c.v1
}

func (c *ClientSet) V2() v2.CoreV2Interface {
	return c.v2
}
