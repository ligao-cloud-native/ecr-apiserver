package token

import (
	regtoken "github.com/docker/distribution/registry/auth/token"
	"strings"
)

func AccessPullPush(a *regtoken.ResourceActions) error {
	a.Actions = permToActions("RWM")
	return nil
}

func AccessOnlyPull(a *regtoken.ResourceActions) error {
	a.Actions = permToActions("R")
	return nil
}

func AccessDeny(a *regtoken.ResourceActions) error {
	a.Actions = []string{}
	return nil
}

func permToActions(p string) (a []string) {
	switch {
	case strings.Contains(p, "W"):
		a = append(a, "push")
		fallthrough
	case strings.Contains(p, "R"):
		a = append(a, "pull")
		fallthrough
	case strings.Contains(p, "M"):
		a = append(a, "*")
	}

	return
}
