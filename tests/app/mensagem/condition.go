package main

import (
	"log"
	"net/http"

	"github.com/patriciapersi/colabore-api/config"
	"github.com/patriciapersi/colabore-api/config/structs"
)

// PRE CONDITION
func GetMessageID(nrInsc, cpf string) string {
	api := config.SetupApi()
	requestBody := structs.PostMessageRequestBody(nrInsc, cpf)
	id := requestBody.ID

	resp, _ := api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(requestBody).
		Post(api.EndpointsAgente["Mensagem"])

	if resp.StatusCode() != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode())
	}

	return id
}

// AFTERCONDITION
func DeleteDataAfterTest(id, nrInsc, cpf string) {
	api := config.SetupApi()
	api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(structs.DeleteAgenteMessageRequestBody(id, nrInsc, cpf)).
		Delete(api.EndpointsAgente["Mensagem"])
}
