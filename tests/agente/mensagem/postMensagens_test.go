package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/patriciapersi/colabore-api/config/structs"
	"github.com/stretchr/testify/assert"
)

func TestPostMensagens(t *testing.T) {
	const (
		nrInsc = "10821992"
		cpf    = "60515860409"
	)

	testCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.Mensagem
		expected     int
		expectedDesc string
	}{
		{
			description:  "Teste envio de mensagem com sucesso",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.PostMessageRequestBody(nrInsc, cpf),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Teste envio de mensagem Sem o nrInsc e sem CPF",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.PostMessageRequestBody("", ""),
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisicao nao contem chaves: NrInscEmpregador",
		},
		{
			description:  "Teste envio de mensagem Sem o nrInsc",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.PostMessageRequestBody("", cpf),
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisicao nao contem chaves: NrInscEmpregador",
		},
		{
			description:  "Teste envio de mensagem sem o CPF",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.PostMessageRequestBody(nrInsc, ""),
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisicao nao contem chaves: Colaboradores[0].CPF",
		},
		{
			description:  "Teste envio de mensagem sem o header",
			setupHeaders: map[string]string{},
			requestBody:  structs.PostMessageRequestBody(nrInsc, cpf),
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()

			id := tc.requestBody.ID

			resp, err := api.Client.R().
				SetHeaders(tc.setupHeaders).
				SetBody(tc.requestBody).
				Post(api.EndpointsAgente["Mensagem"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

			if tc.setupHeaders != nil {
				DeleteDataAfterTest(id, nrInsc, cpf)
			}

		})
	}
}
