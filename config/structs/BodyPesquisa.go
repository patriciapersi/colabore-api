package structs

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func PostPesquisaRequestBody(nrInsc, cpf string) Pesquisa {
	perguntaID := uuid.New().String()

	return Pesquisa{
		ID:                 uuid.New().String(),
		Inicio:             time.Now().Format("02/01/2006"),
		Fim:                time.Now().Format("02/01/2006"),
		NrInscEmpregador:   nrInsc,
		Titulo:             faker.Word(),
		MonitoramentoSaude: false,
		PesquisaAnonima:    true,
		IndependeMatricula: true,
		Versao:             uuid.New().String(),
		Perguntas: []Pergunta{
			{
				ID:          perguntaID,
				Tipo:        SUBJETIVA,
				Ordem:       "1",
				Texto:       faker.Word(),
				Obrigatoria: true,
				NotaMinima:  0,
				NotaMaxima:  0,
				Respostas:   []map[string]interface{}{},
			},
		},
		Colaboradores: []Colaboradores{
			{
				Matricula:        "000031",
				CPF:              cpf,
				NrInscEmpregador: nrInsc,
			},
		},
	}
}

func DeletePesquisaBody(pesquisaId, nrInsc, cpf string) DeletePesquisa {
	return DeletePesquisa{
		Pesquisas: []PesquisaDeletada{
			{
				Matricula:        "000031",
				CPF:              cpf,
				NrInscEmpregador: nrInsc,
				PesquisaId:       pesquisaId,
			},
		},
	}
}
