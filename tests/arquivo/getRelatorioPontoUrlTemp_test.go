package main

import (
	"net/http"
	"testing"

	"github.com/patriciapersi/colabore-api/config"
	testutil "github.com/patriciapersi/colabore-api/util"
	"github.com/stretchr/testify/assert"
)

func TestGetRelatorioPontoUrlTemp(t *testing.T) {
	if err := testutil.LoadEnv(); err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
		t.Fatalf("%v", err)
	}

	testCases := []struct {
		description      string
		cpf              string
		nrInscEmpregador string
		header           map[string]string
		matricula        string
		expected         int
	}{
		{
			description:      "Buscar Relatorio de Ponto com Sucesso",
			cpf:              "60515860409",
			nrInscEmpregador: "10821992",
			header:           config.SetupHeaders(),
			matricula:        "000031",
			expected:         http.StatusOK,
		},
		{
			description:      "Buscar Relatorio de Ponto nrInsc Vazio",
			cpf:              "60515860409",
			nrInscEmpregador: "",
			header:           config.SetupHeaders(),
			matricula:        "000031",
			expected:         http.StatusBadRequest,
		},
		{
			description:      "Buscar Relatorio de Ponto CPF Vazio",
			cpf:              "",
			nrInscEmpregador: "10821992",
			header:           config.SetupHeaders(),
			matricula:        "000031",
			expected:         http.StatusBadRequest,
		},
		{
			description:      "Buscar Relatorio de Ponto Matricula Vazio",
			cpf:              "60515860409",
			nrInscEmpregador: "10821992",
			header:           config.SetupHeaders(),
			matricula:        "",
			expected:         http.StatusBadRequest,
		},
		{
			description:      "Buscar Relatorio de Ponto com Sucesso",
			cpf:              "60515860409",
			nrInscEmpregador: "10821992",
			header:           map[string]string{},
			matricula:        "000031",
			expected:         http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			client := config.SetupClient()
			url := config.BaseURL + "/agente/Arquivo/RelatorioPonto/URLTemporaria"

			queryParams := map[string]string{
				"CPF":              tc.cpf,
				"NrInscEmpregador": tc.nrInscEmpregador,
				"Matricula":        tc.matricula,
				"AnoMes":           "202401",
			}

			resp, err := client.R().
				SetHeaders(tc.header).
				SetQueryParams(queryParams).
				Get(url)

			assert.NoError(t, err, "Erro ao fazer a requisição")
			assert.Equal(t, tc.expected, resp.StatusCode(), "Status de resposta inesperado")
		})
	}
}