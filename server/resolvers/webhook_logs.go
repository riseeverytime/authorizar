package resolvers

import (
	"context"
	"fmt"

	"github.com/authorizerdev/authorizer/server/db"
	"github.com/authorizerdev/authorizer/server/graph/model"
	"github.com/authorizerdev/authorizer/server/token"
	"github.com/authorizerdev/authorizer/server/utils"
	log "github.com/sirupsen/logrus"
)

// WebhookLogsResolver resolver for getting the list of webhook_logs based on pagination & webhook identifier
func WebhookLogsResolver(ctx context.Context, params *model.ListWebhookLogRequest) (*model.WebhookLogs, error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		log.Debug("Failed to get GinContext: ", err)
		return nil, err
	}

	if !token.IsSuperAdmin(gc) {
		log.Debug("Not logged in as super admin")
		return nil, fmt.Errorf("unauthorized")
	}

	var pagination model.Pagination
	var webhookID string

	if params != nil {
		pagination = utils.GetPagination(&model.PaginatedInput{
			Pagination: params.Pagination,
		})
		webhookID = utils.StringValue(params.WebhookID)
	} else {
		pagination = utils.GetPagination(nil)
		webhookID = ""
	}

	webhookLogs, err := db.Provider.ListWebhookLogs(ctx, pagination, webhookID)
	if err != nil {
		log.Debug("failed to get webhook logs: ", err)
		return nil, err
	}
	return webhookLogs, nil
}
