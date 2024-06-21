package agentebody

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/patriciapersi/colabore-api/config/structs"
)

func PostPesquisaRequestBody(nrInsc, cpf string) structs.Pesquisa {
	perguntaID := uuid.New().String()

	return structs.Pesquisa{
		ID:                 uuid.New().String(),
		Inicio:             time.Now().Format("02/01/2006"),
		Fim:                time.Now().Format("02/01/2006"),
		NrInscEmpregador:   nrInsc,
		Titulo:             faker.Word(),
		MonitoramentoSaude: false,
		PesquisaAnonima:    true,
		IndependeMatricula: true,
		Versao:             uuid.New().String(),
		Perguntas: []structs.Pergunta{
			{
				ID:          perguntaID,
				Tipo:        structs.SUBJETIVA,
				Ordem:       "1",
				Texto:       faker.Word(),
				Obrigatoria: true,
				NotaMinima:  0,
				NotaMaxima:  0,
				Respostas:   []map[string]interface{}{},
			},
		},
		Colaboradores: []structs.Colaboradores{
			{
				Matricula:        "000031",
				CPF:              cpf,
				NrInscEmpregador: nrInsc,
			},
		},
	}
}

func DeletePesquisaBody(pesquisaId, nrInsc, cpf string) structs.DeletePesquisa {
	return structs.DeletePesquisa{
		Pesquisas: []structs.PesquisaDeletada{
			{
				Matricula:        "000031",
				CPF:              cpf,
				NrInscEmpregador: nrInsc,
				PesquisaId:       pesquisaId,
			},
		},
	}
}
