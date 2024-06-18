package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/patriciapersi/colabore-api/config"
	"github.com/stretchr/testify/assert"
)

func TestGetPesquisaLista(t *testing.T) {

	testCases := []struct {
		description      string
		id               string
		NrInscEmpregador string
		header           map[string]string
		expected         int
		expectedDesc     string
	}{
		{
			description:      "Buscar Lista de Pesquisas com sucesso",
			id:               precondition(),
			NrInscEmpregador: "10821992",
			header:           config.SetupHeadersApp(),
			expected:         http.StatusOK,
			expectedDesc:     "Sucesso", //valida
		},
		{
			description:      "Buscar Lista de Pesquisas com nrInsc Vazio",
			id:               "",
			NrInscEmpregador: "",
			header:           config.SetupHeadersApp(),
			expected:         http.StatusBadRequest,
			expectedDesc:     "Parâmetro nrInscEmpregador obrigatório.",
		},
		{
			description:      "Buscar Lista de Pesquisas com header Vazio",
			NrInscEmpregador: "",
			header:           map[string]string{},
			expected:         http.StatusUnauthorized,
			expectedDesc:     "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			fmt.Println(tc.id)

			api := config.SetupApi()
			queryParams := map[string]string{
				"NrInscEmpregador": tc.NrInscEmpregador,
			}

			resp, err := api.Client.R().
				SetHeaders(tc.header).
				SetQueryParams(queryParams).
				Get(api.EndpointsApp["PesquisaLista"])

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
			assert.Contains(t, string(resp.Body()), tc.expectedDesc, "Descrição de resposta inesperada")

			//Valida se a Pesquisa Existe
			if tc.header != nil && tc.NrInscEmpregador != "" {
				queryParams := map[string]string{
					"NrInscEmpregador": tc.NrInscEmpregador,
				}

				getResp, _ := validaGet(queryParams)

				var responseData map[string]interface{}
				err := json.Unmarshal(getResp.Body(), &responseData)
				assert.NoError(t, err, "Erro ao fazer o parse da resposta JSON")

				matriculas, _ := responseData["Matriculas"].(map[string]interface{})
				encontrado := false
				for _, matricula := range matriculas {
					if matriculaMap, ok := matricula.(map[string]interface{}); ok {
						if _, ok := matriculaMap[tc.id]; ok {
							encontrado = true
							break
						}
					}
				}

				assert.True(t, encontrado, "O ID não foi encontrado")
			}

			//DELETA A PESQUISA
			if tc.header != nil && tc.NrInscEmpregador != "" {
				deleteDataAfterTest(tc.id)
			}
		})
	}
}

func precondition() string {
	api := config.SetupApi()
	requestBody := config.PostPesquisaRequestBody()
	id := requestBody["id"].(string)
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

func validaGet(queryParams map[string]string) (*resty.Response, error) {
	api := config.SetupApi()
	resp, _ := api.Client.R().
		SetHeaders(config.SetupHeadersApp()).
		SetQueryParams(queryParams).
		Get(api.EndpointsApp["PesquisaLista"])

	return resp, nil
}

func deleteDataAfterTest(id string) {
	api := config.SetupApi()
	api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(config.DeletePesquisaBody(id)).
		Delete(api.EndpointsAgente["Pesquisa"])
}
