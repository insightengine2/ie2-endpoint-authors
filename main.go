package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/insightengine2/ie2-endpoint-authors/lib"
)

func HandleRequest(context context.Context, ev events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

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

	var authors []lib.DBAuthor
	name, ok := ev.QueryStringParameters["name"]

	if !ok {
		authors, err = lib.GetAuthors()
	} else {
		authors, err = lib.GetAuthorsByName(name)
	}

	if err != nil {
		res.StatusCode = 500
		res.Body = err.Error()
		return res, err
	}

	jsonAuthors, err := json.Marshal(authors)
	if err != nil {
		res.StatusCode = 500
		res.Body = err.Error()
		return res, err
	}

	res.Body = string(jsonAuthors)

	return res, nil
}

// entry point to your lambda
func main() {
	lambda.Start(HandleRequest)
}
