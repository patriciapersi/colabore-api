package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	agentebody "github.com/patriciapersi/colabore-api/config/agenteBody"
	"github.com/patriciapersi/colabore-api/config/structs"
	"github.com/stretchr/testify/assert"

	testutil "github.com/patriciapersi/colabore-api/util"
)

func TestPostCandidato(t *testing.T) {
	var tax_id = testutil.GenerateCPF()
	testCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.Candidato
		expected     int
		expectedDesc string
	}{
		{
			description:  "Tentar Reenviar um candidato com com sucesso",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.PostCandidato(nrInsc, tax_id),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Validar um candidato ja enviado",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.PostCandidato(nrInsc, tax_id),
			expected:     http.StatusBadRequest,
			expectedDesc: "Candidato já foi enviado anteriormente",
		},
		{
			description:  "Tentativa de incluir novo candidato sem body",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.Candidato{},
			expected:     http.StatusBadRequest,
			expectedDesc: "ERRO",
		},
		{
			description:  "Tentativa de Envio de novo candidato sem header - Unauthorized",
			setupHeaders: map[string]string{},
			requestBody:  structs.Candidato{},
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
				Post(api.EndpointsAgente["Candidato"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

		})
	}

}
