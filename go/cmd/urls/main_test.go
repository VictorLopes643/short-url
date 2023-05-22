package main

import (
	"fmt"
	"testing"
	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {
	event := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{
			"url": "https://carro.mercadolivre.com.br/MLB-3314406795-a3-sportback-20-s-line-2022-_JM#position=4&search_layout=grid&type=item&tracking_id=6c472aa7-63c6-46e0-abe9-3037ee44fe0b"
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

	if response.Body != "It worked!" {
		t.Errorf("Body esperado: 'It worked!', Body recebido: '%s'", response.Body)
	}

	fmt.Println("Teste concluído com sucesso!")
}
