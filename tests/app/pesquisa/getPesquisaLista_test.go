package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/patriciapersi/colabore-api/helper"
	"github.com/stretchr/testify/assert"
)

func TestGetPesquisaLista(t *testing.T) {
	const (
		nrInsc = "10821992"
		cpf    = "60515860409"
	)

	testsCases := []struct {
		description      string
		id               string
		NrInscEmpregador string
		setupHeaders     map[string]string
		expected         int
		expectedDesc     string
	}{
		{
			description:      "Buscar Lista de Pesquisas com sucesso",
			id:               helper.GetPesquisaID(nrInsc, cpf),
			NrInscEmpregador: nrInsc,
			setupHeaders:     config.SetupHeadersApp(),
			expected:         http.StatusOK,
			expectedDesc:     "Sucesso",
		},
		{
			description:      "Buscar Lista de Pesquisas com nrInsc Vazio",
			NrInscEmpregador: "",
			setupHeaders:     config.SetupHeadersApp(),
			expected:         http.StatusBadRequest,
			expectedDesc:     "Parâmetro nrInscEmpregador obrigatório.",
		},
		{
			description:  "Buscar Lista de Pesquisas com header Vazio",
			setupHeaders: map[string]string{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()
			queryParams := map[string]string{
				"NrInscEmpregador": tc.NrInscEmpregador,
			}

			resp, err := api.Client.R().
				SetHeaders(tc.setupHeaders).
				SetQueryParams(queryParams).
				Get(api.EndpointsApp["PesquisaLista"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
			assert.Contains(t, string(resp.Body()), tc.id, "ID específico não encontrado na resposta")

			//DELETA A PESQUISA
			if tc.setupHeaders != nil && tc.NrInscEmpregador != "" {
				helper.DeletePesquisaAfterTest(tc.id, nrInsc, cpf)
			}
		})
	}
}
