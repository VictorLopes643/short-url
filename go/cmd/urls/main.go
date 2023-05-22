package main

import (
	"fmt"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
    "encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"	
	"github.com/google/uuid"
	"crypto/rand"
	"encoding/hex"



)




type InputEvent  struct {
    Url *string `json:"url"`
}

type Item struct {
    ID        string `json:"id"`
    HASH      string `json:"hash"`
    AVAILABLE bool   `json:"available"`
	DEFAULTURL string `json:"defaultUrl"`
}



func main() {
	lambda.Start(Handler)
}


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if request.HTTPMethod == "OPTIONS" {
		fmt.Println("Chego no OPTIONS")

		response := events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "http://localhost:3000", 
				"Access-Control-Allow-Methods": "POST",
				"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
				"Access-Control-Max-Age":       "3600",
			},
			StatusCode: 200,
			Body:       "",
		}
		return response, nil
	} else if request.HTTPMethod == "POST" {
	

		var event InputEvent
		errs := json.Unmarshal([]byte(request.Body), &event)
		if errs != nil {
			fmt.Println(errs)
			

			return events.APIGatewayProxyResponse{Body: "Error Input", StatusCode: 400}, nil
		}
	// Imprime o título e o preço
	fmt.Println(*event.Url)  

	length := 10 // comprimento do ID em bytes

    randomBytes := make([]byte, length)
    _, errs = rand.Read(randomBytes)
    if errs != nil {
        log.Println("Erro ao gerar ID:", errs)
		

        return events.APIGatewayProxyResponse{Body: "Erro ao gerar ID", StatusCode: 500}, errs
    }

    id := hex.EncodeToString(randomBytes)
	// Cria um novo item sem definir manualmente o ID
	item := Item{
		HASH:        id,
		AVAILABLE:   true,
		DEFAULTURL:  *event.Url,
		ID: uuid.New().String(),
	}
	// Configura uma sessão do AWS SDK
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    })
	// Cria um cliente DynamoDB
	svc := dynamodb.New(sess)
	// Define os parâmetros para adicionar o item à tabela do DynamoDB
	tableName := "URL-flv3xtxnhzgt3m7s5ut34eqj3q-dev"
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Println("Erro ao converter item para DynamoDB AttributeValue:", err)
		return events.APIGatewayProxyResponse{Body: "Erro ao converter item para DynamoDB AttributeValue", StatusCode: 500}, err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	// Chama o método PutItem para adicionar o item à tabela do DynamoDB
	_, err = svc.PutItem(input)
	if err != nil {
		log.Println("Erro ao adicionar o item à tabela:", err)
		

		return events.APIGatewayProxyResponse{Body: "Erro ao adicionar o item à tabela", StatusCode: 500}, err
	}
		response := events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "http://localhost:3000",  // Substitua pelo URL do seu front-end
				"Content-Type":                "application/json",
			},
			StatusCode: 200,
			Body:       "It worked!",
		}
		

		return response, nil
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 404,
		Body:       "Endpoint não encontrado",
	}
	return response, nil
	}


