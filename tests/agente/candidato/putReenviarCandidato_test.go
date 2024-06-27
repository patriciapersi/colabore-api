package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	agentebody "github.com/patriciapersi/colabore-api/config/agenteBody"
	"github.com/patriciapersi/colabore-api/config/structs"
	"github.com/patriciapersi/colabore-api/helper"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

const (
	nrInsc = "10821992"
	taxID  = "60515860409"
	matID  = "000031"
)

func TestPutReenviarCandidato(t *testing.T) {
	var cpf = testutil.GenerateCPF()
	testCases := []struct {
		description  string
		before       func()
		setupHeaders map[string]string
		requestBody  structs.Candidato
		expected     int
		expectedDesc string
	}{
		{
			description:  "Tentar Reenviar um candidato com preenchimento pendente",
			before:       func() { helper.CreateCandidato(cpf, nrInsc) },
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.PostCandidato(nrInsc, cpf),
			expected:     http.StatusBadRequest,
			expectedDesc: "Candidato com preenchimento pendente",
		},
		{
			description:  "Tentar Reenviar um candidato com inexistente",
			setupHeaders: config.SetupHeadersAgente(),
			requestBody:  agentebody.PostCandidato(nrInsc, testutil.GenerateCPF()),
			expected:     http.StatusBadRequest,
			expectedDesc: "Candidato não existe na base de dados",
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
				Put(api.EndpointsAgente["CandidatoRetificar"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

		})
	}

}
