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

const (
	nrInscr = "10821992"
	tax_id  = "60515860409"
	matric  = "000034"
)

func TestPostAprovaAbono(t *testing.T) {

	testCases := []struct {
		description  string
		before       func()
		setupHeaders map[string]string
		requestBody  structs.PostAbonoBody
		expected     int
		expectedDesc string
	}{
		{
			description:  "Aprova Solicitação de Abono com Sucesso",
			before:       func() { helper.CreateAbono(nrInscr, tax_id, matric, structs.PENDENTE) },
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.PostSolicitacaoAbono(nrInscr, tax_id, matric, structs.ACEITO),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentativa de Aprovar de solicitação de abono sem body",
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
