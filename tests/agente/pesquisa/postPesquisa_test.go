package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func TestPostPesquisa(t *testing.T) {

	testsCases := []struct {
		description      string
		NrInscEmpregador string
		setupBody        bool
		header           map[string]string
		expected         int
		expectedDesc     string
	}{
		{
			description:  "Inserir Pesquisas com sucesso",
			setupBody:    true,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Inserir Pesquisas com sucesso",
			setupBody:    false,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém",
		},
		{
			description:  "Inserir Pesquisas com sucesso",
			setupBody:    false,
			header:       map[string]string{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()

			requestBody := config.PostPesquisaRequestBody()
			id := requestBody["id"].(string)

			var body interface{}
			if tc.setupBody {
				body = requestBody

			}
			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(body).
				Post(api.EndpointsAgente["Pesquisa"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

			deleteResp, _ := api.Client.R().
				SetHeaders(config.SetupHeadersAgente()).
				SetBody(config.DeletePesquisaBody(id)).
				Delete(api.EndpointsAgente["Pesquisa"])

			if deleteResp.StatusCode() != http.StatusOK {
				t.Fatalf("Falha ao excluir a pesquisa criada: status %d", deleteResp.StatusCode())
			}

		})
	}
}

// func precondition() string {
// 	api := config.SetupApi()
// 	requestBody := config.PostPesquisaRequestBody()
// 	id := requestBody["id"].(string)
// 	resp, _ := api.Client.R().
// 		SetHeaders(config.SetupHeadersAgente()).
// 		SetBody(requestBody).
// 		Post(api.EndpointsAgente["Pesquisa"])

// 	if resp.StatusCode() != http.StatusOK {
// 		log.Printf("Unexpected status code: %d", resp.StatusCode())
// 		panic("Falha na requisição")
// 	}

// 	return id
// }
