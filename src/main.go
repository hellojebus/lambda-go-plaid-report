package main

import (
	"github.com/aws/aws-lambda-go/events"
	"log"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
)

type body struct {
	 Message string `json="message"`
	 Method string `json="method"`
	 QueryParams map[string]string `json="query_params"`
}


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Lambda Request", request.RequestContext.RequestID)
	b, _ := json.Marshal(body{Message: "hello world", Method: request.HTTPMethod, QueryParams: request.QueryStringParameters})
	return events.APIGatewayProxyResponse{
		Body: string(b),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}