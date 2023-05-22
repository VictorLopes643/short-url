package main

import (
	"fmt"
	"testing"
	"github.com/aws/aws-lambda-go/events"
	"reflect"
)

func TestHandleRequest(t *testing.T) {
	event := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{
			"hash": "e6d3802f933979910af7"
		}`,
	}

	// Chamar a função HandleRequest
	response, err := Handler(event)
	if err != nil {
		t.Errorf("Erro ao chamar HandleRequest: %v", err)
	}

	// Verificar a resposta retornada
	if response.StatusCode != 200 {
		t.Errorf("StatusCode esperado: 200, StatusCode recebido: %d", response.StatusCode)
	}

	if reflect.TypeOf(response.Body).Kind() != reflect.String {
		t.Errorf("Tipo esperado: string, Tipo recebido: %s", reflect.TypeOf(response.Body).Kind())
	}

	fmt.Println("Teste concluído com sucesso!")
}
