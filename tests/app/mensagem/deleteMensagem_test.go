package main

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/patriciapersi/colabore-api/config"
	"github.com/patriciapersi/colabore-api/config/structs"
	"github.com/stretchr/testify/assert"
)

const (
	nrInsc = "10821992"
	cpf    = "60515860409"
)

func TestDeleteMensagens(t *testing.T) {

	testsCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.DeleteAppMensagensRequest
		expected     int
		expectedDesc string
	}{
		{
			description:  "Excluir Mensagem Com Sucesso",
			setupHeaders: config.SetupHeadersApp(),
			requestBody:  structs.DeleteAppMessageRequestBody(GetMessageID(nrInsc, cpf), nrInsc, cpf),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentativa de deletar mensagem com ID inexistente",
			setupHeaders: config.SetupHeadersApp(),
			requestBody:  structs.DeleteAppMessageRequestBody(uuid.New().String(), nrInsc, cpf),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentativa de deletar mensagem com ID vazio",
			setupHeaders: config.SetupHeadersApp(),
			requestBody:  structs.DeleteAppMessageRequestBody("", nrInsc, cpf),
			expected:     http.StatusBadRequest,
			expectedDesc: "MensagemId",
		},
		{
			description:  "Excluir Mensagem Com nrInsccpf vazio",
			setupHeaders: config.SetupHeadersApp(),
			requestBody:  structs.DeleteAppMessageRequestBody(uuid.New().String(), "", ""),
			expected:     http.StatusBadRequest,
			expectedDesc: "NrInscEmpregador",
		},
		{
			description:  "Excluir Mensagem Com header vazio",
			setupHeaders: map[string]string{},
			requestBody:  structs.DeleteAppMessageRequestBody(GetMessageID(nrInsc, cpf), nrInsc, cpf),
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
		{
			description:  "Excluir Mensagem sem body",
			setupHeaders: config.SetupHeadersApp(),
			requestBody:  structs.DeleteAppMensagensRequest{},
			expected:     http.StatusBadRequest,
			expectedDesc: "NrInscEmpregador,MensagemId",
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()

			resp, err := api.Client.R().
				SetHeaders(tc.setupHeaders).
				SetBody(tc.requestBody).
				Delete(api.EndpointsApp["Mensagem"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

		})
	}
}
