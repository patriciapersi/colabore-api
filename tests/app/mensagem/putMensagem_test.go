package main

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func getMessage_ID() string {
	api := config.SetupApi()
	requestBody := config.MensagensRequestBody()
	id := requestBody["ID"].(string)
	resp, _ := api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(requestBody).
		Post(api.EndpointsAgente["Mensagem"])

	if resp.StatusCode() != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode())
		panic("Falha na requisição")
	}
	return id
}

func TestAtualizacaoMensagens(t *testing.T) {

	testsCases := []struct {
		description      string
		header           map[string]string
		setupBody        bool
		NrInscEmpregador string
		id               string
		expected         int
		expectedDesc     string
	}{
		{
			description:      "Atualizar Mensagem com Sucesso",
			header:           config.SetupHeadersApp(),
			setupBody:        true,
			NrInscEmpregador: "10821992",
			id:               getMessage_ID(),
			expected:         http.StatusOK,
			expectedDesc:     "Sucesso",
		},
		{
			description:      "Tentar Atualizar Mensagem com NrInscEMpregado Vazio",
			header:           config.SetupHeadersApp(),
			setupBody:        false,
			NrInscEmpregador: "",
			id:               uuid.New().String(),
			expected:         http.StatusBadRequest,
			expectedDesc:     "ERRO",
		},
		{
			description:      "Tentar Atualizar Mensagem com header Vazio",
			header:           map[string]string{},
			setupBody:        false,
			NrInscEmpregador: "",
			id:               uuid.New().String(),
			expected:         http.StatusUnauthorized,
			expectedDesc:     "Unauthorized",
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.description, func(t *testing.T) {

			api := config.SetupApi()

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetBody(config.PutMensagemLidaAppRequestBody(tc.id)).
				Put(api.EndpointsApp["MensagemLida"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

			//INICIA A VALIDAÇÃO SE A MENSAGEM FOI MARCADA COMO LIDA
			if tc.header != nil && tc.setupBody {
				queryParams := map[string]string{
					"Lastupdate":       "0",
					"NrInscEmpregador": tc.NrInscEmpregador,
				}

				getResp, _ := validaGet(queryParams)

				var responseData map[string]interface{}
				err = json.Unmarshal(getResp.Body(), &responseData)
				assert.NoError(t, err, "Erro ao fazer o parse da resposta JSON")
				mensagens := responseData["Mensagens"].(map[string]interface{})
				mensagem := mensagens[tc.id].(map[string]interface{})
				assert.Equal(t, true, mensagem["Lida"], "Esperava que o campo 'Lida' fosse true")
			}

			//DELETA A MENSAGEM
			if tc.header != nil && tc.setupBody {
				deleteDataAfterTest(tc.id)
			}

		})
	}
}

func deleteDataAfterTest(id string) {
	api := config.SetupApi()
	api.Client.R().
		SetHeaders(config.SetupHeadersApp()).
		SetBody(config.DeleteMensagemAppRequestBody(id)).
		Delete(api.EndpointsApp["Mensagem"])
}

func validaGet(queryParams map[string]string) (*resty.Response, error) {
	api := config.SetupApi()
	resp, _ := api.Client.R().
		SetHeaders(config.SetupHeadersApp()).
		SetQueryParams(queryParams).
		Get(api.EndpointsApp["Mensagem"])

	return resp, nil
}
