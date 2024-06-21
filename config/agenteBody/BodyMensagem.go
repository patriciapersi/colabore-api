package agentebody

import (
	"time"

	"github.com/google/uuid"
	"github.com/patriciapersi/colabore-api/config/structs"
)

func PostMessageRequestBody(nrInsc, cpf string) structs.Mensagem {
	return structs.Mensagem{
		ID:               uuid.New().String(),
		TpInscEmpregador: structs.CNPJ,
		NrInscEmpregador: nrInsc,
		MensagemTitulo:   "Teste automatizado",
		MensagemCorpo:    "Mensagem enviada pelo teste automatizado",
		DataMensagem:     time.Now().Format("02/01/2006"),
		Colaboradores: []structs.Colaborador{
			{
				CPF: cpf,
				Contrato: structs.Contrato{
					Matricula: "000031",
					Cargo:     "ALMOXARIFE",
				},
			},
		},
	}
}

func DeleteAgenteMessageRequestBody(mensagemID, nrInsc, cpf string) structs.DeleteAgenteMensagensRequest {
	return structs.DeleteAgenteMensagensRequest{
		MensagemID:       mensagemID,
		NrInscEmpregador: nrInsc,
		ListaCPF:         []string{cpf},
	}
}

func DeleteAppMessageRequestBody(mensagemID, nrInsc, cpf string) structs.DeleteAppMensagensRequest {
	return structs.DeleteAppMensagensRequest{
		MensagemID:       mensagemID,
		NrInscEmpregador: nrInsc,
		CPF:              cpf,
	}
}

func PutAppMessageRequestBody(mensagemID, nrInsc string) structs.PutAppMensagensRequest {
	return structs.PutAppMensagensRequest{
		NrInscEmpregador: nrInsc,
		MensagemID:       mensagemID,
	}
}
