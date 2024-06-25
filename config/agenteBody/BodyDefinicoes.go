package agentebody

import (
	"github.com/patriciapersi/colabore-api/config/structs"
)

func Definicoes(nrInsc string) structs.Definicoes {
	return structs.Definicoes{
		DefinicoesEmpresa: []structs.DefinicoesEmpresa{
			{
				NrInscEmpregador: nrInsc,
				Ferias: structs.FeriasConfig{
					AntecedenciaMinima:     15,
					HabilitaFerias:         true,
					ExigeAprovacaoDoGestor: true,
				},
			},
		},
	}
}
