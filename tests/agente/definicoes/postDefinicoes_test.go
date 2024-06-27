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
	nrInsc = "10821992"
)

func TestPostDefinicoes(t *testing.T) {

	testCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.Definicoes
		expected     int
		expectedDesc string
	}{
		{
			description:  "Teste envio de Definições com sucesso",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.Definicoes(nrInsc),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Teste envio de Definições sem body",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.Definicoes{},
			expected:     http.StatusBadRequest,
			expectedDesc: "ERRO",
		},
		{
			description:  "Teste envio de Definições sem Header",
			setupHeaders: map[string]string{},
			requestBody:  structs.Definicoes{},
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
				Post(api.EndpointsAgente["LicenciadoDefinicoes"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}

}
