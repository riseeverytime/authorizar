package mongodb

import (
	"time"

	"github.com/authorizerdev/authorizer/server/db/models"
	"github.com/authorizerdev/authorizer/server/graph/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddWebhookLog to add webhook log
func (p *provider) AddWebhookLog(webhookLog models.WebhookLog) (models.WebhookLog, error) {
	if webhookLog.ID == "" {
		webhookLog.ID = uuid.New().String()
	}

	webhookLog.Key = webhookLog.ID
	webhookLog.CreatedAt = time.Now().Unix()
	webhookLog.UpdatedAt = time.Now().Unix()

	webhookLogCollection := p.db.Collection(models.Collections.WebhookLog, options.Collection())
	_, err := webhookLogCollection.InsertOne(nil, webhookLog)
	if err != nil {
		return webhookLog, err
	}
	return webhookLog, nil
}

// ListWebhookLogs to list webhook logs
func (p *provider) ListWebhookLogs(pagination model.Pagination, webhookID string) (*model.WebhookLogs, error) {
	webhookLogs := []*model.WebhookLog{}
	opts := options.Find()
	opts.SetLimit(pagination.Limit)
	opts.SetSkip(pagination.Offset)
	opts.SetSort(bson.M{"created_at": -1})

	paginationClone := pagination
	query := bson.M{}

	if webhookID != "" {
		query = bson.M{"webhook_id": webhookID}
	}

	webhookLogCollection := p.db.Collection(models.Collections.WebhookLog, options.Collection())
	count, err := webhookLogCollection.CountDocuments(nil, query, options.Count())
	if err != nil {
		return nil, err
	}

	paginationClone.Total = count

	cursor, err := webhookLogCollection.Find(nil, query, opts)
	if err != nil {
		return nil, err
	}

	for cursor.Next(nil) {
		var webhookLog models.WebhookLog
		err := cursor.Decode(&webhookLog)
		if err != nil {
			return nil, err
		}
		webhookLogs = append(webhookLogs, webhookLog.AsAPIWebhookLog())
	}

	return &model.WebhookLogs{
		Pagination:  &paginationClone,
		WebhookLogs: webhookLogs,
	}, nil
}
