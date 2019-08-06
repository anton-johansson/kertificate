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

package main

import (
	"kertificate.io/kertificate/pkg/api"
	v1 "kertificate.io/kertificate/pkg/api/v1"
	"kertificate.io/kertificate/pkg/auth"
	"kertificate.io/kertificate/pkg/db"
	"kertificate.io/kertificate/pkg/pki"

	"github.com/spf13/cobra"
)

func init() {
	var command = &cobra.Command{
		Use:   "start",
		Short: "Starts an instance of Kertificate",
		RunE: func(command *cobra.Command, args []string) error {
			return run()
		},
	}

	rootCommand.AddCommand(command)
}

func run() error {
	database := db.NewDatabase()

	certificateTemplateDAO := db.NewCertificateTemplateDAO(database)
	certificateDAO := db.NewCertificateDAO(database)
	commonAuthorityDAO := db.NewCommonAuthorityDAO(database)
	consumerTypeDAO := db.NewConsumerTypeDAO(database)
	userDAO := db.NewUserDAO(database)

	authService := auth.NewAuthService(userDAO)
	keyGenerator := pki.NewKeyGenerator(commonAuthorityDAO)

	authAPI := v1.NewAuthAPI(authService, userDAO)
	certificateTemplateAPI := v1.NewCertificateTemplateAPI(certificateTemplateDAO)
	certificateAPI := v1.NewCertificateAPI(keyGenerator, certificateDAO)
	commonAuthorityAPI := v1.NewCommonAuthorityAPI(keyGenerator, commonAuthorityDAO)
	consumerTypeAPI := v1.NewConsumerTypeAPI(consumerTypeDAO)
	statusAPI := v1.NewStatusAPI()
	versionAPI := v1.NewVersionAPI()

	v1 := v1.NewApiV1(
		authAPI,
		certificateTemplateAPI,
		certificateAPI,
		commonAuthorityAPI,
		consumerTypeAPI,
		statusAPI,
		versionAPI,
	)

	apiServer := api.NewApiServer(v1)

	if err := database.Connect(); err != nil {
		return err
	}
	return apiServer.Start()
}
