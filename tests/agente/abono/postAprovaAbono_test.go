package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func TestPostAprovaAbono(t *testing.T) {

	testCases := []struct {
		description  string
		header       map[string]string
		setupBody    bool
		expected     int
		expectedDesc string
		nrInsc       string
		cpf          string
		matricula    string
		precondition bool
	}{
		{
			description:  "Aprova Solicitação de Abono com Sucesso",
			header:       config.SetupHeadersAgente(),
			setupBody:    true,
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
			nrInsc:       "10821992",
			cpf:          "60515860409",
			matricula:    "000031",
			precondition: true,
		},
		{
			description:  "Tentativa de Aprovar de solicitação de abono sem body",
			header:       config.SetupHeadersAgente(),
			setupBody:    false,
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém nenhum abono",
		},
		{
			description:  "Tentativa de Envio de solicitação de abono sem header - Unauthorized",
			header:       map[string]string{},
			setupBody:    false,
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			if tc.header != nil && tc.precondition {
				preconditionAprova(tc.cpf, tc.nrInsc, tc.matricula)
			}
			api := config.SetupApi()

			// Configura os parâmetros do corpo da requisição se necessário
			var body interface{}
			if tc.setupBody {
				body = config.PostAprovaAbonoBody(tc.cpf, tc.nrInsc, tc.matricula)
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(body).
				Post(api.EndpointsAgente["Abono"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}

}

func preconditionAprova(tax_id string, cnpj string, matricula string) {
	api := config.SetupApi()
	api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(config.PostSolicitaAbono2Body(tax_id, cnpj, matricula)).
		Post(api.EndpointsAgente["Abono"])

}
