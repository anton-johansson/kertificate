package app

import (
	"pkims/api"
	v1 "pkims/api/v1"
	"pkims/auth"
	"pkims/db"
)

// Run builds the entire object graph and runs the application
func Run() {
	database := db.NewDatabase()
	if err := database.Connect(); err != nil {
		panic(err)
	}

	userDAO := db.NewUserDAO(database)
	consumerTypeDAO := db.NewConsumerTypeDAO(database)
	certificateTemplateDAO := db.NewCertificateTemplateDAO(database)

	authService := auth.NewAuthService(userDAO)

	consumerTypeAPI := v1.NewConsumerTypeAPI(consumerTypeDAO)
	certificateTemplateAPI := v1.NewCertificateTemplateAPI(certificateTemplateDAO)

	apiServer := api.NewApiServer(authService)
	v1.InitializeV1(apiServer.V1, authService, consumerTypeAPI, certificateTemplateAPI)
	if err := apiServer.Start(); err != nil {
		panic(err)
	}
}
