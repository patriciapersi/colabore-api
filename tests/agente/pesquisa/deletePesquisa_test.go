package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/patriciapersi/colabore-api/config/structs"
	"github.com/stretchr/testify/assert"
)

const (
	nrInsc = "10821992"
	cpf    = "60515860409"
)

func getId() string {
	api := config.SetupApi()
	requestBody := structs.PostPesquisaRequestBody(nrInsc, cpf)
	id := requestBody.ID
	resp, _ := api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(requestBody).
		Post(api.EndpointsAgente["Pesquisa"])

	if resp.StatusCode() != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode())
		panic("Falha na requisição")
	}

	return id
}

func TestDeletePesquisa(t *testing.T) {

	testsCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.DeletePesquisa
		expected     int
		expectedDesc string
	}{
		{
			description:  "Deletar Pesquisas com sucesso",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.DeletePesquisaBody(getId(), nrInsc, cpf),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentar Deletar Pesquisas com ID Vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.DeletePesquisaBody(nrInsc, cpf, ""),
			expected:     http.StatusBadRequest,
			expectedDesc: "Quantidade de Registros não processados: 1",
		},
		{
			description:  "Deletar Pesquisas com nrInsc Vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.DeletePesquisaBody(getId(), "", cpf),
			expected:     http.StatusBadRequest,
			expectedDesc: "Quantidade de Registros não processados: 1",
		},
		{
			description:  "Tentar Deletar Pesquisas com body vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.DeletePesquisaBody("", "", ""),
			expected:     http.StatusBadRequest,
			expectedDesc: "Quantidade de Registros não processados: 1",
		},
		{
			description:  "Tentar Deletar Pesquisas com header vazio",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  structs.DeletePesquisa{},
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém \\\"Pesquisas\\\"",
		},
		{
			description:  "Deletar Pesquisas com sucesso",
			setupHeaders: map[string]string{},
			requestBody:  structs.DeletePesquisaBody(getId(), nrInsc, cpf),
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()

			resp, err := api.Client.R().
				SetHeaders(tc.setupHeaders).
				SetBody(tc.requestBody).
				Delete(api.EndpointsAgente["Pesquisa"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

		})
	}
}
