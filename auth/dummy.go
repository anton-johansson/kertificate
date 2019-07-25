package auth

import "strings"

type dummy struct {
}

func (dummy *dummy) Login(username string, password string) *authResult {
	if strings.EqualFold(username, "anton2") && password == "s3cr3t" {
		return &authResult{
			username: "anton2",
			active:   true,
		}
	}
	return nil
}
