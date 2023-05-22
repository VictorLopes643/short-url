package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session";
	// "github.com/aws/aws-sdk-go/service/dynamodb"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	// "github.com/aws/aws-xray-sdk-go/awsplugins/dynamodb"
	// "github.com/aws/aws-xray-sdk-go/xray"
)

type Item struct {
    // ID        string `json:"id"`
    HASH      string `json:"hash"`
    // AVAILABLE bool   `json:"available"`
}
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    length := 10 // comprimento do ID em bytes

    randomBytes := make([]byte, length)
    _, err := rand.Read(randomBytes)
    if err != nil {
        log.Println("Erro ao gerar ID:", err)
        return events.APIGatewayProxyResponse{Body: "Erro ao gerar ID", StatusCode: 500}, err
    }

    id := hex.EncodeToString(randomBytes)
    log.Println("ID criado: ", id)

    log.Println("Chegou Aqui")
  
    return events.APIGatewayProxyResponse{Body: id, StatusCode: 200}, nil
}

func main() {
    lambda.Start(handler)
}

