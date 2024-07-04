package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	agentebody "github.com/patriciapersi/colabore-api/config/agenteBody"
	"github.com/patriciapersi/colabore-api/config/structs"
	"github.com/stretchr/testify/assert"
)

const (
	nrInsc    = "10821992"
	cpf       = "60515860409"
	matricula = "000034"
)

func TestPostDefinicoesFerias(t *testing.T) {

	testCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.ColaboradorRequestBody
		expected     int
		expectedDesc string
	}{
		{
			description:  "Teste envio de Definições com sucesso",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.InformacoesFerias(cpf, nrInsc, matricula),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Teste envio de Definições sem CPF",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.InformacoesFerias("", nrInsc, matricula),
			expected:     http.StatusBadRequest,
			expectedDesc: "Quantidade de Registros não processados: 1",
		},
		{
			description:  "Teste envio de Definições sem CPF",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.InformacoesFerias("", "", matricula),
			expected:     http.StatusBadRequest,
			expectedDesc: "Quantidade de Registros não processados: 1",
		},
		{
			description:  "Teste envio de Definições sem CPF",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.InformacoesFerias("", "", ""),
			expected:     http.StatusBadRequest,
			expectedDesc: "Quantidade de Registros não processados: 1",
		},
		{
			description:  "Teste envio de Definições sem body",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.ColaboradorRequestBody{},
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém nenhuma informação de férias",
		},
		{
			description:  "Teste envio de Definições sem body",
			setupHeaders: map[string]string{},
			requestBody:  structs.ColaboradorRequestBody{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()

			resp, err := api.Client.R().
				SetHeaders(tc.setupHeaders).
				SetBody(tc.requestBody).
				Post(api.EndpointsAgente["FeriasInformacoes"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
