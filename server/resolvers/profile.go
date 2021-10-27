package resolvers

import (
	"context"
	"fmt"
	"strings"

	"github.com/authorizerdev/authorizer/server/db"
	"github.com/authorizerdev/authorizer/server/graph/model"
	"github.com/authorizerdev/authorizer/server/session"
	"github.com/authorizerdev/authorizer/server/utils"
)

func Profile(ctx context.Context) (*model.User, error) {
	gc, err := utils.GinContextFromContext(ctx)
	var res *model.User
	if err != nil {
		return res, err
	}

	token, err := utils.GetAuthToken(gc)
	if err != nil {
		return res, err
	}

	claim, err := utils.VerifyAuthToken(token)
	if err != nil {
		return res, err
	}

	userID := fmt.Sprintf("%v", claim["id"])
	email := fmt.Sprintf("%v", claim["email"])
	sessionToken := session.GetToken(userID, token)

	if sessionToken == "" {
		return res, fmt.Errorf(`unauthorized`)
	}

	user, err := db.Mgr.GetUserByEmail(email)
	if err != nil {
		return res, err
	}

	userIdStr := fmt.Sprintf("%v", user.ID)

	res = &model.User{
		ID:              userIdStr,
		Email:           user.Email,
		Image:           &user.Image,
		FirstName:       &user.FirstName,
		LastName:        &user.LastName,
		SignupMethod:    user.SignupMethod,
		EmailVerifiedAt: &user.EmailVerifiedAt,
		Roles:           strings.Split(user.Roles, ","),
		CreatedAt:       &user.CreatedAt,
		UpdatedAt:       &user.UpdatedAt,
	}

	return res, nil
}
