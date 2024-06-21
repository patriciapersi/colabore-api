package structs

import (
	"time"

	"github.com/google/uuid"
)

func PostMessageRequestBody(nrInsc, cpf string) Mensagem {
	return Mensagem{
		ID:               uuid.New().String(),
		TpInscEmpregador: CNPJ,
		NrInscEmpregador: nrInsc,
		MensagemTitulo:   "Teste automatizado",
		MensagemCorpo:    "Mensagem enviada pelo teste automatizado",
		DataMensagem:     time.Now().Format("02/01/2006"),
		Colaboradores: []Colaborador{
			{
				CPF: cpf,
				Contrato: Contrato{
					Matricula: "000031",
					Cargo:     "ALMOXARIFE",
				},
			},
		},
	}
}

func DeleteAgenteMessageRequestBody(mensagemID, nrInsc, cpf string) DeleteAgenteMensagensRequest {
	return DeleteAgenteMensagensRequest{
		MensagemID:       mensagemID,
		NrInscEmpregador: nrInsc,
		ListaCPF:         []string{cpf},
	}
}

func DeleteAppMessageRequestBody(mensagemID, nrInsc, cpf string) DeleteAppMensagensRequest {
	return DeleteAppMensagensRequest{
		MensagemID:       mensagemID,
		NrInscEmpregador: nrInsc,
		CPF:              cpf,
	}
}

func PutAppMessageRequestBody(mensagemID, nrInsc string) PutAppMensagensRequest {
	return PutAppMensagensRequest{
		NrInscEmpregador: nrInsc,
		MensagemID:       mensagemID,
	}
}
