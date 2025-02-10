package sqs

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	cfg "github.com/pangolin-do-golang/thumb-processor-api/internal/config"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"
)

type SQSThumbQueue struct {
	client   *sqs.Client
	queueURL string
}

func NewSQSThumbQueue(c *cfg.Config) (*SQSThumbQueue, error) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	return &SQSThumbQueue{
		client:   sqs.NewFromConfig(sdkConfig),
		queueURL: c.SQS.QueueURL,
	}, nil
}

func (q *SQSThumbQueue) SendEvent(ctx context.Context, process *entity.ThumbProcess) error {
	messageBody, err := json.Marshal(process)
	if err != nil {
		return err
	}

	_, err = q.client.SendMessage(context.Background(), &sqs.SendMessageInput{
		QueueUrl:    &q.queueURL,
		MessageBody: aws.String(string(messageBody)),
	})

	return err
}
