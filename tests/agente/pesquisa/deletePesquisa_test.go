package main

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/patriciapersi/colabore-api/config"
	agentebody "github.com/patriciapersi/colabore-api/config/agenteBody"
	"github.com/patriciapersi/colabore-api/config/structs"
	"github.com/patriciapersi/colabore-api/helper"
	"github.com/stretchr/testify/assert"
)

const (
	nrInsc = "10821992"
	cpf    = "60515860409"
)

func TestDeletePesquisa(t *testing.T) {

	testsCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.DeletePesquisa
		expected     int
		expectedDesc string
	}{
		{
			description:  "Deletar Pesquisas com sucesso",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.DeletePesquisaBody(helper.GetPesquisaID(nrInsc, cpf), nrInsc, cpf),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentar Deletar Pesquisas com ID Vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.DeletePesquisaBody("", nrInsc, cpf),
			expected:     http.StatusBadRequest,
			expectedDesc: "Quantidade de Registros não processados: 1",
		},
		{
			description:  "Deletar Pesquisas com nrInsc Vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.DeletePesquisaBody(uuid.New().String(), "", cpf),
			expected:     http.StatusBadRequest,
			expectedDesc: "Quantidade de Registros não processados: 1",
		},
		{
			description:  "Tentar Deletar Pesquisas com body vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.DeletePesquisaBody("", "", ""),
			expected:     http.StatusBadRequest,
			expectedDesc: "Quantidade de Registros não processados: 1",
		},
		{
			description:  "Tentar Deletar Pesquisas com Sem body",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.DeletePesquisa{},
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém \\\"Pesquisas\\\"",
		},
		{
			description:  "Deletar Pesquisas sem header de autenticação",
			setupHeaders: map[string]string{},
			requestBody:  agentebody.DeletePesquisaBody(uuid.New().String(), nrInsc, cpf),
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
				Delete(api.EndpointsAgente["Pesquisa"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

		})
	}
}
