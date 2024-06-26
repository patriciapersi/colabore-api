package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func TestColaboradorPreemium(t *testing.T) {

	testCases := []struct {
		description      string
		cpf              string
		nrInscEmpregador string
		header           map[string]string
		expected         int
		expectedDesc     string
	}{
		{
			description:      "Busca de candidato - Não Encontrado",
			cpf:              "60515860409",
			nrInscEmpregador: "10821992",
			header:           config.SetupHeadersAgente(),
			expected:         http.StatusBadRequest,
			expectedDesc:     "Candidato com dados preenchidos não encontrado",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()
			queryParams := map[string]string{
				"CPF":              tc.cpf,
				"NrInscEmpregador": tc.nrInscEmpregador,
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetQueryParams(queryParams).
				Get(api.EndpointsAgente["Candidato"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

		})
	}
}
