package agentebody

import (
	"time"

	"github.com/patriciapersi/colabore-api/config/structs"
)

func InformacoesFerias(cpf, nrInsc, matricula string) structs.ColaboradorRequestBody {
	return structs.ColaboradorRequestBody{
		Colaboradores: []structs.ColaboradoresInfo{
			{
				CPF:                            cpf,
				NrInscEmpregador:               "10821992",
				Matricula:                      "000031",
				SolicitouAdiantamento13:        false,
				DiasDisponiveisAbonoPecuniario: 10,
				DiasDisponiveis:                30,
				InicioPeriodoConcessivo:        "2022-06-29",
				FimPeriodoConcessivo:           "2023-07-30",
			},
		},
	}
}

func Ferias(cpf, nrInsc, matricula string) structs.Ferias {
	return structs.Ferias{
		Ferias: []structs.FeriasInfo{
			{
				CPF:                      cpf,
				NrInscEmpregador:         nrInsc,
				Matricula:                matricula,
				SolicitouAdiantamento13:  true,
				SolicitouAbonoPecuniario: false,
				StatusSolicitacao:        1,
				InicioPeriodoGozo:        time.Now().Format("2006-01-02"),
				FimPeriodoGozo:           time.Now().AddDate(0, 0, 20).Format("2006-01-02"),
			},
		},
	}
}
