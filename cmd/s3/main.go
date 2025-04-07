package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	output, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String("localstack-testbuckete6e05abe-57c27c2c"),
		Key:    aws.String("go.mod"),
	})
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(output.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
