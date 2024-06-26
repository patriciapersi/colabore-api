package helper

import (
	"log"
	"net/http"

	"github.com/patriciapersi/colabore-api/config"
	agentebody "github.com/patriciapersi/colabore-api/config/agenteBody"
	"github.com/patriciapersi/colabore-api/config/structs"
)

// setupAPIAndHeaders configura a API e os cabeçalhos
func setupAPIAndHeaders() (*config.API, map[string]string) {
	api := config.SetupApi()
	headers := config.SetupHeadersAgente()
	return api, headers
}

// ---------------------------------------------------
// PRE CONDITION MENSAGEM
func GetMessageID(nrInsc, cpf string) string {
	api, headers := setupAPIAndHeaders()
	requestBody := agentebody.PostMessageRequestBody(nrInsc, cpf)
	id := requestBody.ID

	resp, _ := api.Client.R().
		SetHeaders(headers).
		SetBody(requestBody).
		Post(api.EndpointsAgente["Mensagem"])

	if resp.StatusCode() != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode())
	}

	return id
}

// AFTERCONDITION MENSAGEM
func DeleteDataAfterTest(id, nrInsc, cpf string) {
	api, headers := setupAPIAndHeaders()
	api.Client.R().
		SetHeaders(headers).
		SetBody(agentebody.DeleteAgenteMessageRequestBody(id, nrInsc, cpf)).
		Delete(api.EndpointsAgente["Mensagem"])
}

//----------------------------------------------------------

// PRE CONDITION PARA ABONOS
func CreateAbono(nrInsc, taxID, matricula string, statusSol structs.StatusSolicitacao) {
	api, headers := setupAPIAndHeaders()
	api.Client.R().
		SetHeaders(headers).
		SetBody(agentebody.PostSolicitacaoAbono(nrInsc, taxID, matricula, statusSol)).
		Post(api.EndpointsAgente["Abono"])

}

//----------------------------------------------------------

// PRE CONDITION PARA PESQUISA
func GetPesquisaID(nrInsc, cpf string) string {
	api, headers := setupAPIAndHeaders()
	requestBody := agentebody.PostPesquisaRequestBody(nrInsc, cpf)
	id := requestBody.ID
	resp, _ := api.Client.R().
		SetHeaders(headers).
		SetBody(requestBody).
		Post(api.EndpointsAgente["Pesquisa"])

	if resp.StatusCode() != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode())
		panic("Falha na requisição")
	}

	return id
}

func DeletePesquisaAfterTest(id, nrInsc, cpf string) {
	api, headers := setupAPIAndHeaders()
	api.Client.R().
		SetHeaders(headers).
		SetBody(agentebody.DeletePesquisaBody(id, nrInsc, cpf)).
		Delete(api.EndpointsAgente["Pesquisa"])
}

//----------------------------------------------------------

// PRE CONDITION PARA CANDIDATO
func CreateCandidato(tax_id string, cnpj string) {
	api, headers := setupAPIAndHeaders()
	api.Client.R().
		SetHeaders(headers).
		SetBody(agentebody.PostCandidato(cnpj, tax_id)).
		Post(api.EndpointsAgente["Candidato"])

}
