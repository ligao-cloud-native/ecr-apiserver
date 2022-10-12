package token

import (
	"k8s.io/klog/v2"
	"net/http"
)

type CreatorType string

const (
	RegistryCreator CreatorType = "registry"
	NotaryCreator   CreatorType = "notary"
)

type CreatorInterface interface {
	CreateToken(r *http.Request) (*ECRToken, error)
}

type CreatorFactory interface {
	Create() CreatorInterface
}

type creatorFactory struct {
	service string
}

func NewCreatorFactory(service string) CreatorFactory {
	return &creatorFactory{
		service: service,
	}
}

func (f *creatorFactory) Create() CreatorInterface {
	switch CreatorType(f.service) {
	case RegistryCreator:
		return newRegistryCreator()
	case NotaryCreator:
		return newNotaryCreator()
	default:
		klog.Errorf("not supported service type [%s]", f.service)
		return nil
	}

}
