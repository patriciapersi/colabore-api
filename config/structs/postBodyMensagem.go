package structs

import (
	"time"

	"github.com/google/uuid"
)

func MessageRequestBody(nrInsc, cpf string) Mensagem {
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
