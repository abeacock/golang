package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func HandleButtonPress(ctx context.Context, event events.IoTButtonEvent) error {
	sess := session.Must(session.NewSession())
	svc := sns.New(sess)

	eventAsJson, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
	}

	message := fmt.Sprintf("Hello from your IoT Button, here is the full event: %s", string(eventAsJson))

	params := &sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(os.Getenv("PhoneNumber")),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and Message from an error.
		fmt.Println(err.Error())
	}

	fmt.Println(resp)

	return nil
}

func main() {
	lambda.Start(HandleButtonPress)
}
