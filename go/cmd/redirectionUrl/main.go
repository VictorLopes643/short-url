package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

)



type InputEvent  struct {
    Hash *string `json:"hash"`
}

type Item struct {
	DefaultURL 	string `json:"defaultUrl"`
	Hash 		string `json:"hash"`
	Id 			string `json:"id"`
}

func main() {
	lambda.Start(Handler)
}


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if request.HTTPMethod == "OPTIONS" {
		response := events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "http://localhost:3000",
				"Access-Control-Allow-Methods": "POST, OPTIONS",
				"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
				"Access-Control-Max-Age":       "3600",
			},
			StatusCode: 200,
			Body:       "",
		}
		return response, nil
	} else if request.HTTPMethod == "POST" {
		var event InputEvent
		err := json.Unmarshal([]byte(request.Body), &event)
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
		}
		// Configurar a sessão do AWS SDK
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"), // substitua pela sua região
		})
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		// Criar um cliente DynamoDB
		svc := dynamodb.New(sess)
		// Configurar o input da consulta
		filt := expression.Name("hash").Equal(expression.Value(*event.Hash))
		// proj := expression.NamesList(expression.Name("hash"), expression.Name("defaultUrl"))
		proj := expression.NamesList(
			expression.Name("hash"),
			expression.Name("defaultUrl"),
			expression.Name("id"),
		)
		expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

		if err != nil {
			fmt.Println("Got error building expression:", err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		params := &dynamodb.ScanInput{
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			FilterExpression:          expr.Filter(),
			ProjectionExpression:      expr.Projection(),
			TableName:                 aws.String("URL-flv3xtxnhzgt3m7s5ut34eqj3q-dev"),
		}
		result, err := svc.Scan(params)
		if err != nil {
			fmt.Println("Got error scan:", err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}
		var items []Item

		for _, i := range result.Items{
			item := Item{}

			err = dynamodbattribute.UnmarshalMap(i, &item)
			if err != nil {
				fmt.Println("Got error unmarshalling:")
				fmt.Println(err)
				return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
			}
			items = append(items, item) 
		}
		// jsonItems, err := json.Marshal(items)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		var defaultUrl string

		defaultUrl = items[0].DefaultURL

		fmt.Println("defaultUrl:", defaultUrl)

		response := events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "http://localhost:3000",  // Substitua pelo URL do seu front-end
				"Content-Type":                "application/json",
			},
			StatusCode: 200,
			Body:       defaultUrl,
		}
		return response, nil
	}  
	response := events.APIGatewayProxyResponse{
		StatusCode: 404,
		Body:       "Endpoint não encontrado",
	}
	return response, nil
}