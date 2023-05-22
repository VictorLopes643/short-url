package main

import (
	"fmt"
	"testing"
	"github.com/aws/aws-lambda-go/events"
	"reflect"
)

func TestHandleRequestGet(t *testing.T) {
	event := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
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
func TestHandleRequestDelete(t *testing.T) {
	event := events.APIGatewayProxyRequest{
		HTTPMethod: "DELETE",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{
			"id": "92d46d9f-d584-4fdb-b353-619d9a9318fd"
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
