package main

import (
	"log"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"fmt"
)

type body struct {
	 Message string `json="message"`
	 Method string `json="method"`
}


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Lambda Request", request.RequestContext.RequestID)

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("us-west-1")},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}
	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("us-west-1"))
	keyname := "/MyService/MyApp/Dev/DATABASE_URI"
	withDecryption := false
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           &keyname,
		WithDecryption: &withDecryption,
	})

	value := *param.Parameter.Value
	fmt.Println(value)
	b, _ := json.Marshal(body{Message: "hello world", Method: request.HTTPMethod})
	return events.APIGatewayProxyResponse{
		Body: string(b),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}