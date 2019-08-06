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
