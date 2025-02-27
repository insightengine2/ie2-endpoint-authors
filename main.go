package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/insightengine2/ie2-endpoint-authors/lib"
)

type MyEvent struct {
	// sample properties below
	InstanceId string `json:"InstanceId"`
	Status     string `json:"Status"`
}

func HandleRequest(context context.Context, ev MyEvent) (events.APIGatewayProxyResponse, error) {

	res := events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers:         nil,
		Body:            "{}",
	}

	_, err := config.LoadDefaultConfig(context)

	if err != nil {
		res.StatusCode = 500
		res.Body = err.Error()
		return res, err
	}

	res.Body = lib.GetAuthors()

	return res, nil
}

// entry point to your lambda
func main() {
	lambda.Start(HandleRequest)
}
