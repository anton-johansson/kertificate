// Copyright 2019 Anton Johansson
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
