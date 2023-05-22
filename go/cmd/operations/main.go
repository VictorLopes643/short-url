package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	// "flag"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    // "encoding/json"
	// "github.com/aws/aws-sdk-go/aws/session"
    // "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"	
	// "github.com/google/uuid"
	// "crypto/rand"
	// "encoding/hex"
	"encoding/json"

)




func main() {
	lambda.Start(Handler)
}
type DeleteEvent  struct {
    Id *string `json:"id"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if request.HTTPMethod == "OPTIONS" {
		fmt.Println("Chego no OPTIONS")

		response := events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "http://localhost:3000", 
				"Access-Control-Allow-Methods": "GET,OPTIONS,POST,DELETE",
				"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
				"Access-Control-Max-Age":       "3600",
			},
			StatusCode: 200,
			Body:       "",
		}
		return response, nil
	} else if request.HTTPMethod == "GET" {
		
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

		// Construa o objeto de input para a operação de Scan
		input := &dynamodb.ScanInput{
			TableName: aws.String("URL-flv3xtxnhzgt3m7s5ut34eqj3q-dev"), // Substitua pelo nome da sua tabela
		}

		// Execute a operação de Scan
		result, err := svc.Scan(input)
		if err != nil {
			fmt.Println("Erro ao executar o Scan:", err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		jsonData, err := json.Marshal(result.Items)
		if err != nil {
			fmt.Println("Erro ao converter para JSON:", err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
		}
			response := events.APIGatewayProxyResponse{
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "http://localhost:3000",  // Substitua pelo URL do seu front-end
					"Content-Type":                "application/json",
				},
				StatusCode: 200,
				Body: string(jsonData),
			}
			return response, nil
	}else if request.HTTPMethod == "DELETE" {
		var event DeleteEvent
		err := json.Unmarshal([]byte(request.Body), &event)
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
		}
		fmt.Println("Chego", *event.Id)
		fmt.Println("Chego", event)

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"),
		})
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		svc := dynamodb.New(sess)

		input := &dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(*event.Id),
				},
			},
			TableName: aws.String("URL-flv3xtxnhzgt3m7s5ut34eqj3q-dev"),
		}

		_, err = svc.DeleteItem(input)
		if err != nil {
			fmt.Println("Got error deleting item:", err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		response := events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "http://localhost:3000",
				"Content-Type":                "application/json",
			},
			StatusCode: 200,
			Body:       "Item deleted successfully",
		}
		return response, nil
	}
	return events.APIGatewayProxyResponse{Body: "Error operation", StatusCode: 500}, nil
	}
