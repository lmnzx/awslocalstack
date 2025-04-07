package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	cfClient := cloudformation.NewFromConfig(cfg)
	stackName := "LocalStack"

	resp, err := cfClient.DescribeStacks(ctx, &cloudformation.DescribeStacksInput{
		StackName: &stackName,
	})
	if err != nil {
		log.Fatalf("Error describing stack: %v", err)
	}

	for _, output := range resp.Stacks[0].Outputs {
		if *output.OutputKey == "BucketName" {
			fmt.Printf("Bucket Name: %s\n", *output.OutputValue)
		}
		if *output.OutputKey == "QueueUrl" {
			fmt.Printf("Queue URL: %s\n", *output.OutputValue)
		}
	}

}
