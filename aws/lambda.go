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
	"github.com/aws/aws-sdk-go/service/ssm"
)

func handleButtonPress(ctx context.Context, event events.IoTButtonEvent) error {
	sess := session.Must(session.NewSession())

	eventAsJSON, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
	}

	// Hello from your IoT Button, here is the full event: %s
	messageTemplate := os.Getenv("MessageTemplate")
	message := fmt.Sprintf(messageTemplate, string(eventAsJSON))

	systemsManager := ssm.New(sess)
	parameterKey := "PhoneNumber"
	param, err := systemsManager.GetParameter(&ssm.GetParameterInput{
		Name: &parameterKey,
	})
	phoneNumber := *param.Parameter.Value

	notificationService := sns.New(sess)
	params := &sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(phoneNumber),
	}
	resp, err := notificationService.Publish(params)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(resp)

	return nil
}

func main() {
	lambda.Start(handleButtonPress)
}
