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
// PRE CONDITION
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

// AFTERCONDITION
func DeleteDataAfterTest(id, nrInsc, cpf string) {
	api, headers := setupAPIAndHeaders()
	api.Client.R().
		SetHeaders(headers).
		SetBody(agentebody.DeleteAgenteMessageRequestBody(id, nrInsc, cpf)).
		Delete(api.EndpointsAgente["Mensagem"])
}

//----------------------------------------------------------

// PRE CONDITION PARA ABONOS
func CreateAbonoStatusPendente(nrInsc, taxID, matricula string, statusSol structs.StatusSolicitacao) {
	api := config.SetupApi()
	api.Client.R().
		SetHeaders(config.SetupHeadersAgente()).
		SetBody(agentebody.PostSolicitacaoAbono(nrInsc, taxID, matricula, statusSol)).
		Post(api.EndpointsAgente["Abono"])

}
