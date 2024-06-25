package agentebody

import (
	"github.com/patriciapersi/colabore-api/config/structs"
)

func PostCandidato(nrInsc, taxID string) structs.Candidato {
	return structs.Candidato{
		NrInscEmpregador: nrInsc,
		CPF:              taxID,
	}
}
