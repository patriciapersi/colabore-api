package main

import (
	"log"
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

func getMessageID() string {
	api := config.SetupApi()
	requestBody := structs.PostMessageRequestBody(nrInsc, cpf)
	id := requestBody.ID

	resp, _ := api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(requestBody).
		Post(api.EndpointsAgente["Mensagem"])

	if resp.StatusCode() != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode())
	}

	return id
}

func TestDeleteMensagens(t *testing.T) {

	testsCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.DeleteMensagensRequest
		expected     int
		expectedDesc string
	}{
		{
			description:  "Excluir Mensagem Com Sucesso",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.DeleteMessageRequestBody(getMessageID(), nrInsc, cpf),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Excluir Mensagem Com ID Inexistente",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.DeleteMessageRequestBody(uuid.New().String(), nrInsc, cpf),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Excluir Mensagem Com nrInsc vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.DeleteMessageRequestBody(uuid.New().String(), "", cpf),
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém chaves: NrInscEmpregador",
		},
		{
			description:  "Excluir Mensagem Com cpf vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.DeleteMessageRequestBody(uuid.New().String(), nrInsc, ""),
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém chaves: ListaCPF[0]",
		},
		{
			description:  "Excluir Mensagem Com nrInsccpf vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.DeleteMessageRequestBody(uuid.New().String(), "", ""),
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém chaves: NrInscEmpregador",
		},
		{
			description:  "Excluir Mensagem Com header vazio",
			setupHeaders: map[string]string{},
			requestBody:  structs.DeleteMessageRequestBody(getMessageID(), nrInsc, cpf),
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()

			resp, err := api.Client.R().
				SetHeaders(tc.setupHeaders).
				SetBody(tc.requestBody).
				Delete(api.EndpointsAgente["Mensagem"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

		})
	}
}
