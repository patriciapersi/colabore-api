package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func TestGetPontoContemBatidas(t *testing.T) {

	testCases := []struct {
		description      string
		NrInscEmpregador string
		header           map[string]string
		expected         int
		expectedDesc     string
	}{
		{
			description:      "Buscar Batidas de Ponto com sucesso",
			NrInscEmpregador: "10821992",
			header:           config.SetupHeadersApp(),
			expected:         http.StatusOK,
			expectedDesc:     "Sucesso",
		},
		{
			description:      "Buscar Mensagens sem o parametro de EMpregador",
			NrInscEmpregador: "",
			header:           config.SetupHeadersApp(),
			expected:         http.StatusBadRequest,
			expectedDesc:     "NrInscEmpregador",
		},
		{
			description:      "Buscar Mensagens sem o parametro de EMpregador",
			NrInscEmpregador: "",
			header:           map[string]string{},
			expected:         http.StatusUnauthorized,
			expectedDesc:     "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()
			queryParams := map[string]string{
				"NrInscEmpregador": tc.NrInscEmpregador,
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetQueryParams(queryParams).
				Get(api.EndpointsApp["PontoContemBatidas"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
