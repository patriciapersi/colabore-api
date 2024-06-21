package main

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/patriciapersi/colabore-api/config"
	agentebody "github.com/patriciapersi/colabore-api/config/agenteBody"
	"github.com/patriciapersi/colabore-api/config/structs"
	"github.com/stretchr/testify/assert"
)

func TestAtualizacaoMensagens(t *testing.T) {

	testsCases := []struct {
		description  string
		setupHeaders map[string]string
		requestBody  structs.PutAppMensagensRequest
		expected     int
		expectedDesc string
	}{
		{
			description:  "Atualizar Mensagem com Sucesso - Marcar como Lida",
			setupHeaders: config.SetupHeadersApp(),
			requestBody:  agentebody.PutAppMessageRequestBody(GetMessageID(nrInsc, cpf), nrInsc),
			expected:     http.StatusOK,
			expectedDesc: "Sucesso",
		},
		{
			description:  "Tentar Atualizar Mensagem com NrInscEMpregado Vazio",
			setupHeaders: config.SetupHeadersApp(),
			requestBody:  agentebody.PutAppMessageRequestBody(uuid.New().String(), ""),
			expected:     http.StatusBadRequest,
			expectedDesc: "ERRO",
		},
		{
			description:  "Tentar Atualizar Mensagem com header Vazio",
			setupHeaders: map[string]string{},
			requestBody:  agentebody.PutAppMessageRequestBody(uuid.New().String(), nrInsc),
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
				Put(api.EndpointsApp["MensagemLida"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

			//DELETA A MENSAGEM
			if tc.setupHeaders != nil {
				id := tc.requestBody.MensagemID
				DeleteDataAfterTest(id, nrInsc, cpf)
			}

		})
	}
}
