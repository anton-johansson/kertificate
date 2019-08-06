package auth

import (
	"strings"

	"kertificate.io/kertificate/pkg/model"
)

type dummy struct {
}

const constantName = "anton3"

func (dummy *dummy) Login(username string, password string) *model.User {
	if strings.EqualFold(username, constantName) && password == "s3cr3t" {
		return &model.User{
			Username:     constantName,
			FirstName:    "Anton",
			LastName:     "Johansson",
			EmailAddress: "antoon.johansson@gmail.com",
		}
	}
	return nil
}
