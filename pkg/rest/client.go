package rest

import (
	"crypto/tls"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"k8s.io/klog/v2"
	"net/http"
	"strings"
)

type Interface interface {
	Request(method, subPath string) *gorequest.SuperAgent
	Response(resp gorequest.Response, errs *[]error) []byte
}

type RestClient struct {
	client  *gorequest.SuperAgent
	baseUrl string
}

func NewRestClient(config *Config) (*RestClient, error) {
	c := gorequest.New()
	c.SetBasicAuth(config.Username, config.Password)

	c.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	h := c.Set("Accept", "application/json")
	if config.UserAgent != "" {
		h.Set("User-Agent", config.UserAgent)
	}
	if config.Timeout != 0 {
		c.Timeout(config.Timeout)
	}

	rc := RestClient{client: c}
	rc.setBaseURL(config.Host, config.APIPath)

	return &rc, nil

}

func (c *RestClient) Request(method, subPath string) *gorequest.SuperAgent {
	if !strings.HasPrefix(subPath, "/") {
		subPath = "/" + subPath
	}
	targetUrl := c.baseUrl + subPath
	klog.Infof("%s %s", method, targetUrl)
	switch method {
	case gorequest.GET:
		return c.client.Get(targetUrl)
	case gorequest.POST:
		return c.client.Post(targetUrl).Set("Content-Type", "application/json")
	default:
		return c.client.Get(targetUrl)
	}

}

// errs为数组结构，因为需要修改，所以传引用类型，
func (c *RestClient) Response(resp gorequest.Response, errs *[]error) []byte {
	if resp == nil {
		klog.Errorf("response error: %v", *errs)
		return nil
	}

	if resp.StatusCode >= http.StatusBadRequest {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			*errs = append(*errs, err)
		} else {
			*errs = append(*errs, fmt.Errorf("%s", string(body)))
		}
		klog.Errorf("response error: [code: %d] [error: %v]", resp.StatusCode, *errs)
		return nil
	}

	if resp.StatusCode >= http.StatusOK && resp.StatusCode <= http.StatusIMUsed {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			*errs = append(*errs, err)
			klog.Errorf("response error: [code: %d] [error: %v]", resp.StatusCode, *errs)
			return nil
		}
		return body
	}

	return nil
}

func (c *RestClient) setBaseURL(host, apiPath string) {
	baseUrl := host
	if apiPath == "" {
		c.baseUrl = strings.TrimRight(baseUrl, "/")
	}

	if !strings.HasSuffix(host, "/") {
		baseUrl += "/"
	}
	if !strings.HasPrefix(apiPath, "/") {
		baseUrl += strings.TrimRight(apiPath, "/")
	} else {
		baseUrl += strings.Trim(apiPath, "/")
	}

	c.baseUrl = baseUrl
}
