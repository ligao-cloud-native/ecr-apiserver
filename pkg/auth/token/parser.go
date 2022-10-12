package token

import (
	"fmt"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db/models"
	"k8s.io/klog/v2"
	"strings"
)

type image struct {
	// harbor project
	namespace string
	// image name, also is repo name which need to create
	name string
	// image tag
	tag string
}

type parser interface {
	parse(img string) (*image, error)
}

// baseParser parse base image: <namespace>/<name>:<tag>
type baseParser struct {
}

func (p *baseParser) parse(img string) (*image, error) {
	return parseImage(img)
}

// endpointParser parse image with repo url: <repo>/<namespace>/<name>:<tag>
type endpointParser struct {
}

func (p *endpointParser) parse(img string) (*image, error) {
	repoAndNsImage := strings.SplitN(img, "/", 2)
	if len(repoAndNsImage) < 2 {
		return nil, fmt.Errorf("cannot parse image from %s. please use '<repo>/<namespace>/<name>:<tag>'", img)
	}
	ep := "https://" + repoAndNsImage[0]
	hb, _ := models.Harbor{}.Get(ep)
	klog.Info(hb.Region)

	return parseImage(repoAndNsImage[1])
}

func parseImage(img string) (*image, error) {
	nsAndImage := strings.SplitN(img, "/", 2)
	if len(nsAndImage) < 2 {
		return nil, fmt.Errorf("cannot parse image from %s. please use '<namespace>/<name>:<tag>'", img)
	}

	nameAndTag := strings.SplitN(nsAndImage[1], ":", 2)

	image := &image{
		namespace: nsAndImage[0],
		name:      nameAndTag[0],
	}
	if len(nameAndTag) == 2 {
		image.tag = nameAndTag[1]
	}

	return image, nil
}
