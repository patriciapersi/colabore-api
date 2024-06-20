package main

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"

	testutil "github.com/patriciapersi/colabore-api/util"
)

func main() {
	cpf := testutil.GenerateCPF()

	fmt.Println("CPF gerado:", cpf)

}

func TestPostCandidato(t *testing.T) {
	var cpf = testutil.GenerateCPF()
	testCases := []struct {
		description  string
		header       map[string]string
		setupBody    bool
		expected     int
		expectedDesc string
		nrInsc       string
		cpf          string
	}{
		{
			description:  "Enviar um novo candidato com sucesso",
			header:       config.SetupHeadersAgente(),
			setupBody:    true,
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
			nrInsc:       "10821992",
			cpf:          cpf,
		},
		{
			description:  "Validar um candidato ja enviado",
			header:       config.SetupHeadersAgente(),
			setupBody:    true,
			expected:     http.StatusBadRequest,
			expectedDesc: "Candidato já foi enviado anteriormente",
			nrInsc:       "10821992",
			cpf:          cpf,
		},
		{
			description:  "Tentativa de incluir novo candidato sem body",
			header:       config.SetupHeadersAgente(),
			setupBody:    false,
			expected:     http.StatusBadRequest,
			expectedDesc: "Corpo da requisição não contém candidato",
		},
		{
			description:  "Tentativa de Envio de novo candidato sem header - Unauthorized",
			header:       map[string]string{},
			setupBody:    false,
			expected:     http.StatusUnauthorized,
			expectedDesc: "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()

			// Configura os parâmetros do corpo da requisição se necessário
			var body interface{}
			if tc.setupBody {
				body = config.PostCandidatoBody(tc.cpf, tc.nrInsc)
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(body).
				Post(api.EndpointsAgente["Candidato"])

			assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

		})
	}

}
