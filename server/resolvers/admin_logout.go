package resolvers

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/authorizerdev/authorizer/server/cookie"
	"github.com/authorizerdev/authorizer/server/graph/model"
	"github.com/authorizerdev/authorizer/server/token"
	"github.com/authorizerdev/authorizer/server/utils"
)

// AdminLogoutResolver is a resolver for admin logout mutation
func AdminLogoutResolver(ctx context.Context) (*model.Response, error) {
	var res *model.Response

	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		log.Debug("Failed to get GinContext: ", err)
		return res, err
	}

	if !token.IsSuperAdmin(gc) {
		log.Debug("Admin is not logged in")
		return res, fmt.Errorf("unauthorized")
	}

	cookie.DeleteAdminCookie(gc)

	res = &model.Response{
		Message: "admin logged out successfully",
	}
	return res, nil
}
