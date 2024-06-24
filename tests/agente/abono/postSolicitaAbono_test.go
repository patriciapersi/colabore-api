package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	agentebody "github.com/patriciapersi/colabore-api/config/agenteBody"
	"github.com/patriciapersi/colabore-api/config/structs"
	helper "github.com/patriciapersi/colabore-api/helper"
	"github.com/stretchr/testify/assert"
)

func TestPostSolicitaAbono(t *testing.T) {

	testCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.PostAbonoBody
		expected     int
		expectedDesc string
	}{
		{
			description:  "Envia Solicitação de Abono com Sucesso",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.PostSolicitacaoAbono(nrInsc, cpf, matricula, structs.PENDENTE),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentativa de Envio de solicitação de abono sem body",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.PostAbonoBody{},
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém nenhum abono",
		},
		{
			description:  "Tentativa de Envio de solicitação de abono sem header - Unauthorized",
			setupHeaders: map[string]string{},
			requestBody:  structs.PostAbonoBody{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			if tc.setupHeaders != nil && len(tc.requestBody.Abonos) > 0 {
				helper.CreateAbonoStatusPendente(nrInsc, cpf, matricula, structs.PENDENTE)
			}
			api := config.SetupApi()

			resp, err := api.Client.R().
				SetHeaders(tc.setupHeaders).
				SetBody(tc.requestBody).
				Post(api.EndpointsAgente["Abono"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}

}
