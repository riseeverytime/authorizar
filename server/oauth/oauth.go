package oauth

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	facebookOAuth2 "golang.org/x/oauth2/facebook"
	githubOAuth2 "golang.org/x/oauth2/github"

	"github.com/authorizerdev/authorizer/server/constants"
	"github.com/authorizerdev/authorizer/server/memorystore"
)

// OAuthProviders is a struct that contains reference all the OAuth providers
type OAuthProvider struct {
	GoogleConfig   *oauth2.Config
	GithubConfig   *oauth2.Config
	FacebookConfig *oauth2.Config
}

// OIDCProviders is a struct that contains reference all the OpenID providers
type OIDCProvider struct {
	GoogleOIDC *oidc.Provider
}

var (
	// OAuthProviders is a global variable that contains instance for all enabled the OAuth providers
	OAuthProviders OAuthProvider
	// OIDCProviders is a global variable that contains instance for all enabled the OpenID providers
	OIDCProviders OIDCProvider
)

// InitOAuth initializes the OAuth providers based on EnvData
func InitOAuth() error {
	ctx := context.Background()
	googleClientID, err := memorystore.Provider.GetStringStoreEnvVariable(constants.EnvKeyGoogleClientID)
	if err != nil {
		googleClientID = ""
	}
	googleClientSecret, err := memorystore.Provider.GetStringStoreEnvVariable(constants.EnvKeyGoogleClientSecret)
	if err != nil {
		googleClientSecret = ""
	}
	if googleClientID != "" && googleClientSecret != "" {
		p, err := oidc.NewProvider(ctx, "https://accounts.google.com")
		if err != nil {
			return err
		}
		OIDCProviders.GoogleOIDC = p
		OAuthProviders.GoogleConfig = &oauth2.Config{
			ClientID:     googleClientID,
			ClientSecret: googleClientSecret,
			RedirectURL:  "/oauth_callback/google",
			Endpoint:     OIDCProviders.GoogleOIDC.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
		}
	}

	githubClientID, err := memorystore.Provider.GetStringStoreEnvVariable(constants.EnvKeyGithubClientID)
	if err != nil {
		githubClientID = ""
	}
	githubClientSecret, err := memorystore.Provider.GetStringStoreEnvVariable(constants.EnvKeyGithubClientSecret)
	if err != nil {
		githubClientSecret = ""
	}
	if githubClientID != "" && githubClientSecret != "" {
		OAuthProviders.GithubConfig = &oauth2.Config{
			ClientID:     githubClientID,
			ClientSecret: githubClientSecret,
			RedirectURL:  "/oauth_callback/github",
			Endpoint:     githubOAuth2.Endpoint,
		}
	}

	facebookClientID, err := memorystore.Provider.GetStringStoreEnvVariable(constants.EnvKeyFacebookClientID)
	if err != nil {
		facebookClientID = ""
	}
	facebookClientSecret, err := memorystore.Provider.GetStringStoreEnvVariable(constants.EnvKeyFacebookClientSecret)
	if err != nil {
		facebookClientSecret = ""
	}
	if facebookClientID != "" && facebookClientSecret != "" {
		OAuthProviders.FacebookConfig = &oauth2.Config{
			ClientID:     facebookClientID,
			ClientSecret: facebookClientSecret,
			RedirectURL:  "/oauth_callback/facebook",
			Endpoint:     facebookOAuth2.Endpoint,
			Scopes:       []string{"public_profile", "email"},
		}
	}

	return nil
}
