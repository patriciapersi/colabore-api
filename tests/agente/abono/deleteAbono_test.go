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

func TestDeleteAbono(t *testing.T) {

	testCases := []struct {
		description  string
		before       func()
		setupHeaders map[string]string
		requestBody  structs.AbonoBody
		expected     int
		expectedDesc string
	}{
		{
			description:  "Deletar Abono pendente",
			before:       func() { helper.CreateAbono(nrInscr, tax_id, matric, structs.PENDENTE) },
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.Abono(cnpj, taxID, matID),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentativa de deletar um abono sem body",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.AbonoBody{},
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém nenhum abono",
		},
		{
			description:  "Tentativa de deletar um abono sem header - Unauthorized",
			setupHeaders: map[string]string{},
			requestBody:  structs.AbonoBody{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			if tc.before != nil {
				tc.before()
			}

			api := config.SetupApi()

			resp, err := api.Client.R().
				SetHeaders(tc.setupHeaders).
				SetBody(tc.requestBody).
				Delete(api.EndpointsAgente["Abono"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}

}
