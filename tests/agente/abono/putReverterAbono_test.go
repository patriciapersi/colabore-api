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
	cnpj  = "10821992"
	taxID = "60515860409"
	matID = "000031"
)

func TestPutReveterAbonoo(t *testing.T) {

	testCases := []struct {
		description  string
		before       func()
		setupHeaders map[string]string
		requestBody  structs.PutAbonoBody
		expected     int
		expectedDesc string
	}{
		{
			description:  "Reverter o status de um Abono para pendente",
			before:       func() { helper.CreateAbono(nrInscr, tax_id, matric, structs.PENDENTE) },
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.PutReverterSolicitacaoAbono(cnpj, taxID, matID),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentativa de reverter um abono sem body",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.PutAbonoBody{},
			expected:     http.StatusBadRequest,
			expectedDesc: "Chave \\\"Abonos\\\" não encontrada.",
		},
		{
			description:  "Tentativa de reverter um abono sem header - Unauthorized",
			setupHeaders: map[string]string{},
			requestBody:  structs.PutAbonoBody{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			if tc.setupHeaders != nil && len(tc.requestBody.Abonos) > 0 {
				helper.CreateAbono(cnpj, taxID, matID, structs.PENDENTE)
			}

			api := config.SetupApi()

			resp, err := api.Client.R().
				SetHeaders(tc.setupHeaders).
				SetBody(tc.requestBody).
				Put(api.EndpointsAgente["ReverterAbono"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}

}
