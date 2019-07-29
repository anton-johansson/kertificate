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

	certificateTemplateDAO := db.NewCertificateTemplateDAO(database)
	consumerTypeDAO := db.NewConsumerTypeDAO(database)
	userDAO := db.NewUserDAO(database)

	authService := auth.NewAuthService(userDAO)

	authAPI := v1.NewAuthAPI(authService)
	certificateTemplateAPI := v1.NewCertificateTemplateAPI(certificateTemplateDAO)
	consumerTypeAPI := v1.NewConsumerTypeAPI(consumerTypeDAO)
	statusAPI := v1.NewStatusAPI()
	versionAPI := v1.NewVersionAPI()

	v1 := v1.NewApiV1(authAPI, certificateTemplateAPI, consumerTypeAPI, statusAPI, versionAPI)

	apiServer := api.NewApiServer(v1)
	if err := apiServer.Start(); err != nil {
		panic(err)
	}
}
