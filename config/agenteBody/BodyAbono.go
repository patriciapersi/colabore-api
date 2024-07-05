package agentebody

import (
	"time"

	"github.com/patriciapersi/colabore-api/config/structs"
)

func PostSolicitacaoAbono(nrInsc, taxID, matricula string, statusSol structs.StatusSolicitacao) structs.PostAbonoBody {
	return structs.PostAbonoBody{
		Abonos: []structs.Abono{
			{
				NrInscEmpregador: nrInsc,
				Evento:           "3",
				CPF:              taxID,
				Matricula:        matricula,
				Nome:             "Sandra Simone Cecília Martins",
				DataAbono:        time.Now().Format("2006-01-02"),
				MotivoId:         "00101",
				StatusSol:        statusSol,
				Turnos:           []string{"1", "2", "3", "4"},
			},
		},
	}
}

func Abono(nrInsc, taxID, matricula string) structs.AbonoBody {
	return structs.AbonoBody{
		Abonos: []structs.AbonoReverteDelete{
			{
				NrInscEmpregador: nrInsc,
				Evento:           "3",
				CPF:              taxID,
				Matricula:        matricula,
				DataAbono:        time.Now().Format("2006-01-02"),
				MotivoId:         "00101",
				Turnos:           []string{"1", "2", "3", "4"},
			},
		},
	}
}
