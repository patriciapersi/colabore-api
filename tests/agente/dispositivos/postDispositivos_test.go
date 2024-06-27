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

func TestPostDispositivos(t *testing.T) {

	testCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.Dispositivos
		expected     int
		expectedDesc string
	}{
		{
			description:  "Inserir Dispositivo com Sucesso",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.Dispositivo(nrInsc),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentar inserir dispositivo sem Body",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.Dispositivos{},
			expected:     http.StatusBadRequest,
			expectedDesc: "Cnpj,DispositivoId,Status",
		},
		{
			description:  "Tentar inserir dispositivo sem header",
			setupHeaders: map[string]string{},
			requestBody:  structs.Dispositivos{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	// Itera sobre os casos de teste
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()

			resp, err := api.Client.R().
				SetHeaders(tc.setupHeaders).
				SetBody(tc.requestBody).
				Post(api.EndpointsAgente["DispositivosStatus"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
