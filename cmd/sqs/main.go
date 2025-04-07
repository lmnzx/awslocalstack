package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := sqs.NewFromConfig(cfg)
	queue, err := client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String("LocalStack-mycustomsqsqueue85419269-42b3a7b7"),
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String("hello, sqs!"),
		QueueUrl:    queue.QueueUrl,
	})
	if err != nil {
		log.Fatal(err)
	}
}
