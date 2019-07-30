package app

import (
	"pkims.io/pkims/pkg/api"
	v1 "pkims.io/pkims/pkg/api/v1"
	"pkims.io/pkims/pkg/auth"
	"pkims.io/pkims/pkg/db"
	"pkims.io/pkims/pkg/pki"
)

// Run builds the entire object graph and runs the application
func Run() {
	database := db.NewDatabase()

	certificateTemplateDAO := db.NewCertificateTemplateDAO(database)
	commonAuthorityDAO := db.NewCommonAuthorityDAO(database)
	consumerTypeDAO := db.NewConsumerTypeDAO(database)
	userDAO := db.NewUserDAO(database)

	authService := auth.NewAuthService(userDAO)
	keyGenerator := pki.NewKeyGenerator()

	authAPI := v1.NewAuthAPI(authService)
	certificateTemplateAPI := v1.NewCertificateTemplateAPI(certificateTemplateDAO)
	commonAuthorityAPI := v1.NewCommonAuthorityAPI(keyGenerator, commonAuthorityDAO)
	consumerTypeAPI := v1.NewConsumerTypeAPI(consumerTypeDAO)
	statusAPI := v1.NewStatusAPI()
	versionAPI := v1.NewVersionAPI()

	v1 := v1.NewApiV1(authAPI, certificateTemplateAPI, commonAuthorityAPI, consumerTypeAPI, statusAPI, versionAPI)

	apiServer := api.NewApiServer(v1)

	if err := database.Connect(); err != nil {
		panic(err)
	}

	if err := apiServer.Start(); err != nil {
		panic(err)
	}
}
