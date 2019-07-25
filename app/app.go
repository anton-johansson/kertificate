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
	authenticationService := auth.NewAuthService(userDAO)
	apiServer := api.NewApiServer()
	v1.InitializeV1(apiServer.V1, authenticationService)
	if err := apiServer.Start(); err != nil {
		panic(err)
	}
}
