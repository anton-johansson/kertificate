package auth

import "strings"

type dummy struct {
}

const name = "anton3"

func (dummy *dummy) Login(username string, password string) *authResult {
	if strings.EqualFold(username, name) && password == "s3cr3t" {
		return &authResult{
			username: name,
			active:   true,
		}
	}
	return nil
}
