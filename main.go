package main

import (
	"context"
	"os"
	"strings"

	"github.com/JoseGaldamez/gambitBackend/awsgo"
	"github.com/JoseGaldamez/gambitBackend/db"
	"github.com/JoseGaldamez/gambitBackend/handlers"
	"github.com/JoseGaldamez/gambitBackend/tools"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(runLambda)
}

func runLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {

	awsgo.InitializeAWS()

	if !tools.ValidateEnvironment() {
		panic("Environment variables not set, SecretName, UrlPrefix")
	}

	var response *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", 1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	headers := request.Headers

	db.ReadSecret()

	status, message := handlers.Handlers(path, method, body, headers, request)

	headersResponse := map[string]string{
		"Content-Type": "application/json",
	}

	response = &events.APIGatewayProxyResponse{
		Body:       message,
		Headers:    headersResponse,
		StatusCode: status,
	}

	return response, nil
}
