package dbpool

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/neutrinocorp/life-track-api/internal/infrastructure/configuration"
	"github.com/neutrinocorp/life-track-api/internal/infrastructure/remote"
)

// NewDynamoDBPool creates a new AWS DynamoDB connection pool
func NewDynamoDBPool(cfg configuration.Configuration) *dynamodb.DynamoDB {
	return dynamodb.New(remote.NewAWSSession(cfg.DynamoTable.Region))
}
