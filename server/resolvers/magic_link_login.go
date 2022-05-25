package resolvers

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/authorizerdev/authorizer/server/constants"
	"github.com/authorizerdev/authorizer/server/db"
	"github.com/authorizerdev/authorizer/server/db/models"
	"github.com/authorizerdev/authorizer/server/email"
	"github.com/authorizerdev/authorizer/server/envstore"
	"github.com/authorizerdev/authorizer/server/graph/model"
	"github.com/authorizerdev/authorizer/server/token"
	"github.com/authorizerdev/authorizer/server/utils"
)

// MagicLinkLoginResolver is a resolver for magic link login mutation
func MagicLinkLoginResolver(ctx context.Context, params model.MagicLinkLoginInput) (*model.Response, error) {
	var res *model.Response

	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		log.Debug("Failed to get GinContext: ", err)
		return res, err
	}

	if envstore.EnvStoreObj.GetBoolStoreEnvVariable(constants.EnvKeyDisableMagicLinkLogin) {
		log.Debug("Magic link login is disabled.")
		return res, fmt.Errorf(`magic link login is disabled for this instance`)
	}

	params.Email = strings.ToLower(params.Email)

	if !utils.IsValidEmail(params.Email) {
		log.Debug("Invalid email")
		return res, fmt.Errorf(`invalid email address`)
	}

	log := log.WithFields(log.Fields{
		"email": params.Email,
	})

	inputRoles := []string{}

	user := models.User{
		Email: params.Email,
	}

	// find user with email
	existingUser, err := db.Provider.GetUserByEmail(params.Email)
	if err != nil {
		if envstore.EnvStoreObj.GetBoolStoreEnvVariable(constants.EnvKeyDisableSignUp) {
			log.Debug("Signup is disabled.")
			return res, fmt.Errorf(`signup is disabled for this instance`)
		}

		user.SignupMethods = constants.SignupMethodMagicLinkLogin
		// define roles for new user
		if len(params.Roles) > 0 {
			// check if roles exists
			if !utils.IsValidRoles(params.Roles, envstore.EnvStoreObj.GetSliceStoreEnvVariable(constants.EnvKeyRoles)) {
				log.Debug("Invalid roles: ", params.Roles)
				return res, fmt.Errorf(`invalid roles`)
			} else {
				inputRoles = params.Roles
			}
		} else {
			inputRoles = envstore.EnvStoreObj.GetSliceStoreEnvVariable(constants.EnvKeyDefaultRoles)
		}

		user.Roles = strings.Join(inputRoles, ",")
		user, _ = db.Provider.AddUser(user)
	} else {
		user = existingUser
		// There multiple scenarios with roles here in magic link login
		// 1. user has access to protected roles + roles and trying to login
		// 2. user has not signed up for one of the available role but trying to signup.
		// 		Need to modify roles in this case

		if user.RevokedTimestamp != nil {
			log.Debug("User access is revoked at: ", user.RevokedTimestamp)
			return res, fmt.Errorf(`user access has been revoked`)
		}

		// find the unassigned roles
		if len(params.Roles) <= 0 {
			inputRoles = envstore.EnvStoreObj.GetSliceStoreEnvVariable(constants.EnvKeyDefaultRoles)
		}
		existingRoles := strings.Split(existingUser.Roles, ",")
		unasignedRoles := []string{}
		for _, ir := range inputRoles {
			if !utils.StringSliceContains(existingRoles, ir) {
				unasignedRoles = append(unasignedRoles, ir)
			}
		}

		if len(unasignedRoles) > 0 {
			// check if it contains protected unassigned role
			hasProtectedRole := false
			for _, ur := range unasignedRoles {
				if utils.StringSliceContains(envstore.EnvStoreObj.GetSliceStoreEnvVariable(constants.EnvKeyProtectedRoles), ur) {
					hasProtectedRole = true
				}
			}

			if hasProtectedRole {
				log.Debug("User is not assigned one of the protected roles", unasignedRoles)
				return res, fmt.Errorf(`invalid roles`)
			} else {
				user.Roles = existingUser.Roles + "," + strings.Join(unasignedRoles, ",")
			}
		} else {
			user.Roles = existingUser.Roles
		}

		signupMethod := existingUser.SignupMethods
		if !strings.Contains(signupMethod, constants.SignupMethodMagicLinkLogin) {
			signupMethod = signupMethod + "," + constants.SignupMethodMagicLinkLogin
		}

		user.SignupMethods = signupMethod
		user, _ = db.Provider.UpdateUser(user)
		if err != nil {
			log.Debug("Failed to update user: ", err)
		}
	}

	hostname := utils.GetHost(gc)
	if !envstore.EnvStoreObj.GetBoolStoreEnvVariable(constants.EnvKeyDisableEmailVerification) {
		// insert verification request
		_, nonceHash, err := utils.GenerateNonce()
		if err != nil {
			log.Debug("Failed to generate nonce: ", err)
			return res, err
		}
		redirectURLParams := "&roles=" + strings.Join(inputRoles, ",")
		if params.State != nil {
			redirectURLParams = redirectURLParams + "&state=" + *params.State
		}
		if params.Scope != nil && len(params.Scope) > 0 {
			redirectURLParams = redirectURLParams + "&scope=" + strings.Join(params.Scope, " ")
		}
		redirectURL := utils.GetAppURL(gc)
		if params.RedirectURI != nil {
			redirectURL = *params.RedirectURI
		}

		if strings.Contains(redirectURL, "?") {
			redirectURL = redirectURL + "&" + redirectURLParams
		} else {
			redirectURL = redirectURL + "?" + redirectURLParams
		}

		verificationType := constants.VerificationTypeMagicLinkLogin
		verificationToken, err := token.CreateVerificationToken(params.Email, verificationType, hostname, nonceHash, redirectURL)
		if err != nil {
			log.Debug("Failed to create verification token: ", err)
		}
		_, err = db.Provider.AddVerificationRequest(models.VerificationRequest{
			Token:       verificationToken,
			Identifier:  verificationType,
			ExpiresAt:   time.Now().Add(time.Minute * 30).Unix(),
			Email:       params.Email,
			Nonce:       nonceHash,
			RedirectURI: redirectURL,
		})
		if err != nil {
			log.Debug("Failed to add verification request in db: ", err)
			return res, err
		}

		// exec it as go routing so that we can reduce the api latency
		go email.SendVerificationMail(params.Email, verificationToken, hostname)
	}

	res = &model.Response{
		Message: `Magic Link has been sent to your email. Please check your inbox!`,
	}

	return res, nil
}
