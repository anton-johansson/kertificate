package v1

import (
	"net/http"

	"kertificate.io/kertificate/pkg/auth"
	"kertificate.io/kertificate/pkg/db"

	echo "github.com/labstack/echo/v4"
)

var ignoredPaths = [...]string{"/v1/authentication/authenticate", "/v1/status", "/v1/version"}

type AuthAPI struct {
	authService *auth.AuthService
	userDAO     *db.UserDAO
}

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthAPI(authService *auth.AuthService, userDAO *db.UserDAO) *AuthAPI {
	return &AuthAPI{authService, userDAO}
}

func (api *AuthAPI) Register(group *echo.Group) {
	group.POST("/authenticate", api.authenticate)
	group.GET("/me", api.getSelf)
}

func (api *AuthAPI) GetAuthMiddleware() echo.MiddlewareFunc {
	return api.checkRequest
}

func (api *AuthAPI) authenticate(context echo.Context) error {
	data := authRequest{
		Username: "",
		Password: "",
	}
	if err := context.Bind(&data); err != nil {
		return err
	}
	user, token, err := api.authService.Login(data.Username, data.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return nil
	}
	sendNewToken(context, token)
	context.JSON(http.StatusOK, user)
	return nil
}

func (api *AuthAPI) getSelf(context echo.Context) error {
	userId := userId(context)
	user, err := api.userDAO.GetUser(userId)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, user)
}

func (api *AuthAPI) checkRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		requestPath := context.Request().URL.RequestURI()
		if func(path string) bool {
			for _, ignoredPath := range ignoredPaths {
				if ignoredPath == path {
					return true
				}
			}
			return false
		}(requestPath) {
			return next(context)
		}

		token := context.Request().Header.Get("Authorization")
		userId, newToken, err := api.authService.CheckToken(token)
		if err != nil {
			return err
		}
		context.Set("userId", int(userId))
		if len(newToken) > 0 {
			sendNewToken(context, newToken)
		}
		return next(context)
	}
}

func sendNewToken(context echo.Context, token string) {
	context.Response().Header().Add("X-Set-Authorization", token)
}
