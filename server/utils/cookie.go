package utils

import (
	"net/http"

	"github.com/authorizerdev/authorizer/server/constants"
	"github.com/authorizerdev/authorizer/server/envstore"
	"github.com/gin-gonic/gin"
)

// SetCookie sets the cookie in the response. It sets 2 cookies
// 1 COOKIE_NAME for the host (abc.com)
// 2 COOKIE_NAME-client for the domain (sub.abc.com).
// Note all sites don't allow 2nd type of cookie
func SetCookie(gc *gin.Context, token string) {
	secure := true
	httpOnly := true
	host, _ := GetHostParts(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyAuthorizerURL))
	domain := GetDomainName(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyAuthorizerURL))
	if domain != "localhost" {
		domain = "." + domain
	}

	gc.SetSameSite(http.SameSiteNoneMode)
	gc.SetCookie(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyCookieName), token, 3600, "/", host, secure, httpOnly)
	gc.SetCookie(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyCookieName)+"-client", token, 3600, "/", domain, secure, httpOnly)
}

// GetCookie gets the cookie from the request
func GetCookie(gc *gin.Context) (string, error) {
	cookie, err := gc.Request.Cookie(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyCookieName))
	if err != nil {
		cookie, err = gc.Request.Cookie(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyCookieName) + "-client")
		if err != nil {
			return "", err
		}
	}

	return cookie.Value, nil
}

// DeleteCookie sets the cookie value as empty to make it expired
func DeleteCookie(gc *gin.Context) {
	secure := true
	httpOnly := true

	host, _ := GetHostParts(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyAuthorizerURL))
	domain := GetDomainName(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyAuthorizerURL))
	if domain != "localhost" {
		domain = "." + domain
	}

	gc.SetSameSite(http.SameSiteNoneMode)
	gc.SetCookie(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyCookieName), "", -1, "/", host, secure, httpOnly)
	gc.SetCookie(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyCookieName)+"-client", "", -1, "/", domain, secure, httpOnly)
}

// SetAdminCookie sets the admin cookie in the response
func SetAdminCookie(gc *gin.Context, token string) {
	secure := true
	httpOnly := true
	host, _ := GetHostParts(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyAuthorizerURL))

	gc.SetCookie(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyAdminCookieName), token, 3600, "/", host, secure, httpOnly)
}

func GetAdminCookie(gc *gin.Context) (string, error) {
	cookie, err := gc.Request.Cookie(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyAdminCookieName))
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func DeleteAdminCookie(gc *gin.Context) {
	secure := true
	httpOnly := true
	host, _ := GetHostParts(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyAuthorizerURL))

	gc.SetCookie(envstore.EnvInMemoryStoreObj.GetStringStoreEnvVariable(constants.EnvKeyAdminCookieName), "", -1, "/", host, secure, httpOnly)
}
