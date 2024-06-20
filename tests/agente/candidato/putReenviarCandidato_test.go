package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestPutReenviarCandidato(t *testing.T) {
	var cpf = testutil.GenerateCPF()
	testCases := []struct {
		description  string
		header       map[string]string
		setupBody    bool
		expected     int
		expectedDesc string
		nrInsc       string
		cpf          string
		precondition bool
	}{
		{
			description:  "Reenviar um candidato com com preenchimento pendente",
			header:       config.SetupHeadersAgente(),
			setupBody:    true,
			expected:     http.StatusBadRequest,
			expectedDesc: "Candidato com preenchimento pendente",
			nrInsc:       "10821992",
			cpf:          cpf,
			precondition: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.header != nil && tc.precondition {
				preconditionCandidato(tc.cpf, tc.nrInsc)
			}
			api := config.SetupApi()

			// Configura os parâmetros do corpo da requisição se necessário
			var body interface{}
			if tc.setupBody {
				body = config.PutReenviarCandidatoBody(tc.cpf, tc.nrInsc)
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(body).
				Put(api.EndpointsAgente["CandidatoRetificar"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

		})
	}

}

func preconditionCandidato(tax_id string, cnpj string) {
	api := config.SetupApi()
	api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(config.PostCandidatoBody(tax_id, cnpj)).
		Post(api.EndpointsAgente["Candidato"])

}
