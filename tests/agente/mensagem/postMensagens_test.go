package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/patriciapersi/colabore-api/config/structs"
	"github.com/stretchr/testify/assert"
)

func TestPostMensagens(t *testing.T) {

	// Definindo uma tabela de casos de teste
	testCases := []struct {
		description  string
		setupBody    bool
		header       map[string]string
		expected     int
		expectedDesc string
	}{
		{
			description:  "Teste envio de mensagem",
			setupBody:    true,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Teste envio de mensagem",
			setupBody:    false,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusBadRequest,
			expectedDesc: "Requisicao deveria conter chave 'Mensagem'",
		},
		{
			description:  "Teste envio de mensagem sem header",
			setupBody:    false,
			header:       map[string]string{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	// Iterando sobre os casos de teste
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()

			requestBody := structs.MessageRequestBody("10821992", "60515860409")
			id := requestBody.ID
			fmt.Println(id)

			var body interface{}
			if tc.setupBody {
				body = requestBody

			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(body).
				Post(api.EndpointsAgente["Mensagem"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

			if tc.header != nil && tc.setupBody {
				deleteDataAfterTest(id)
			}
		})
	}
}

func deleteDataAfterTest(id string) {
	api := config.SetupApi()
	api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(config.DeleteMensagensRequestBody(id)).
		Delete(api.EndpointsAgente["Mensagem"])
}
