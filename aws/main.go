package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
)

func HandleRequest(ctx context.Context, event events.IoTButtonEvent) (string, error) {
	fmt.Println("creating session")
	sess := session.Must(session.NewSession())
	fmt.Println("session created")

	svc := sns.New(sess)
	fmt.Println("service created")

	params := &sns.PublishInput{
		Message:     aws.String("The user did a " + event.ClickType + " press"),
		PhoneNumber: aws.String(os.Getenv("PhoneNumber")),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(resp)

	return fmt.Sprintf("Hello %s!", event.SerialNumber), nil
}

func main() {
	lambda.Start(HandleRequest)
}
