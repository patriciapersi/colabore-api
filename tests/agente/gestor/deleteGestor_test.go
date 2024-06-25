//Aguardando script que faça inclusão da solicitação de férias pelo próprio teste

package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

const (
	nrInsc = "10821992"
	cpf    = "12658729375"
)

func TestDeleteGestor(t *testing.T) {

	testCases := []struct {
		description  string
		setupBody    bool
		header       map[string]string
		expected     int
		expectedDesc string
	}{
		{
			description:  "Deleção de Gestores",
			setupBody:    true,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Deleção de Gestores - Sem body",
			setupBody:    false,
			header:       config.SetupHeadersAgente(),
			expected:     http.StatusBadRequest,
			expectedDesc: "Não há informação a ser processada.",
		},
		{
			description:  "Deleção de Gestores - Unauthorized",
			setupBody:    false,
			header:       map[string]string{},
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			api := config.SetupApi()
			var body interface{}
			if tc.setupBody {
				body = config.GestoresRequestBody()

			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(body).
				Delete(api.EndpointsAgente["Gestor"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}
}
