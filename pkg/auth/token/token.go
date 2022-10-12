package token

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"github.com/docker/distribution/registry/auth/token"
	"github.com/docker/libtrust"
	jsoniter "github.com/json-iterator/go"
	"k8s.io/klog/v2"
	"strings"
	"time"
)

const (
	expiration = 1800 // second
	issuer     = "harbor-token-issuer"
)

type ECRToken struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
	IssuedAt  string `json:"issued_at"`
}

// https://goharbor.io/docs/2.3.0/install-config/configure-https/
//https://goharbor.io/docs/2.3.0/install-config/customize-token-service/
// https://github.com/distribution/distribution/blob/main/docs/spec/auth/token.md
func MakeToken(access []*token.ResourceActions) (*ECRToken, error) {
	// 该私钥文件与harbor-core服务中的/etc/core/private_key.pem内容一致。
	pk, err := libtrust.LoadKeyFile("/etc/ecr/private_key.pem")
	if err != nil {
		klog.Errorf("load private key file failed: %v", err)
		return nil, err
	}

	// jwt header
	jwtHeader := token.Header{
		Type:       "JWT",
		SigningAlg: "RS256",
		KeyID:      pk.KeyID(),
	}

	now := time.Now().UTC()
	// jwt token
	jwtToken := token.ClaimSet{
		Issuer:     issuer,
		Subject:    "username",
		Audience:   "service",
		Expiration: now.Add(time.Duration(expiration) * time.Second).Unix(),
		NotBefore:  now.Unix(),
		IssuedAt:   now.Unix(),
		JWTID:      "",
		Access:     access,
	}

	jwtHeaderJson, err := jsoniter.Marshal(jwtHeader)
	if err != nil {
		klog.Errorf("marshal jwt header failed: %s", err.Error())
		return nil, err
	}
	jwtTokenJson, err := jsoniter.Marshal(jwtToken)
	if err != nil {
		klog.Errorf("marshal jwt token failed: %s", err.Error())
		return nil, err
	}
	encodedJWTHeader := base64URLEncode(jwtHeaderJson)
	encodedJWTToke := base64URLEncode(jwtTokenJson)
	payload := fmt.Sprintf("%s.%s", encodedJWTHeader, encodedJWTToke)

	signature, _, err := pk.Sign(strings.NewReader(payload), crypto.SHA256)
	if err != nil {
		klog.Errorf("sign jwt payload failed: %s", err.Error())
		return nil, err
	}
	encodedSignature := base64URLEncode(signature)
	tokeStr := fmt.Sprintf("%s.%s", payload, encodedSignature)

	regToken, err := token.NewToken(tokeStr)
	if err != nil {
		klog.Errorf("new docker registry token error: %s", err.Error())
		return nil, err
	}

	return &ECRToken{
		Token:     fmt.Sprintf("%s.%s", regToken.Raw, base64URLEncode(regToken.Signature)),
		ExpiresIn: expiration,
		IssuedAt:  now.Format(time.RFC3339),
	}, nil

}

func base64URLEncode(data []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(data), "=")
}
