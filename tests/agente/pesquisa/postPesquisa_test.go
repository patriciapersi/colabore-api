package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/patriciapersi/colabore-api/config/structs"
	"github.com/stretchr/testify/assert"
)

func TestPostPesquisa(t *testing.T) {
	const (
		nrInsc = "10821992"
		cpf    = "60515860409"
	)

	testsCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.Pesquisa
		expected     int
		expectedDesc string
	}{
		{
			description:  "Inserir Pesquisas com sucesso",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.PostPesquisaRequestBody(nrInsc, cpf),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Inserir Pesquisa com nrinsc e cpf vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.PostPesquisaRequestBody("", ""),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Inserir Pesquisa com header vazio",
			setupHeaders: map[string]string{},
			requestBody:  structs.PostPesquisaRequestBody("", ""),
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
		{
			description:  "Inserir Pesquisa com body vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.Pesquisa{},
			expected:     http.StatusBadRequest,
			expectedDesc: "Quantidade de Registros não processados: 1",
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()

			id := tc.requestBody.ID

			resp, err := api.Client.R().
				SetHeaders(tc.setupHeaders).
				SetBody(tc.requestBody).
				Post(api.EndpointsAgente["Pesquisa"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

			if tc.setupHeaders != nil {
				deleteDataAfterTest(id, nrInsc, cpf)
			}

		})
	}
}

func deleteDataAfterTest(id, nrInsc, cpf string) {
	api := config.SetupApi()
	api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(structs.DeletePesquisaBody(id, nrInsc, cpf)).
		Delete(api.EndpointsAgente["Pesquisa"])
}
