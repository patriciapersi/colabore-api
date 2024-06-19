package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
)

func TestPutReveterAbono(t *testing.T) {

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
			description:  "Reverter o status de um Abono para pendente",
			header:       config.SetupHeadersAgente(),
			setupBody:    true,
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
			nrInsc:       "10821992",
			cpf:          "60515860409",
			matricula:    "000031",
			precondition: true,
		},
		// {
		// 	description:  "Tentativa de reverter um abono sem body",
		// 	header:       config.SetupHeadersAgente(),
		// 	setupBody:    false,
		// 	expected:     http.StatusBadRequest,
		// 	expectedDesc: "Corpo da requisição não contém nenhum abono",
		// },
		// {
		// 	description:  "Tentativa de reverter um abono sem header - Unauthorized",
		// 	header:       map[string]string{},
		// 	setupBody:    false,
		// 	expected:     http.StatusUnauthorized,
		// 	expectedDesc: "Unauthorized",
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.header != nil && tc.precondition {
				precondition(tc.nrInsc, tc.cpf, tc.matricula)
			}
			// api := config.SetupApi()

			// // Configura os parâmetros do corpo da requisição se necessário
			// var body interface{}
			// if tc.setupBody {
			// 	body = config.PutReveterAbonoBody(tc.nrInsc, tc.cpf, tc.matricula)
			// }

			// resp, err := api.Client.R().
			// 	SetHeaders(tc.header).
			// 	SetBody(body).
			// 	Put(api.EndpointsAgente["ReverterAbono"])

			// assert.NoError(t, err, "Erro ao fazer a requisição para %s", tc.description)
			// assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado para %s", tc.description)
			// assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")
		})
	}

}

func precondition(tax_id string, cnpj string, matricula string) {
	api := config.SetupApi()
	requestBody := config.PostSolicitaAbono2Body(tax_id, cnpj, matricula)
	resp, _ := api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(requestBody).
		Post(api.EndpointsAgente["Abono"])

	if resp.StatusCode() != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode())
		panic("Falha na requisição")
	}

}
