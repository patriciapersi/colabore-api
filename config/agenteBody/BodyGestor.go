package agentebody

import "github.com/patriciapersi/colabore-api/config/structs"

func Gestores(nrInsc, cpf, matricula string) structs.Gestores {
	return structs.Gestores{
		Gestores: []structs.Gestor{
			{
				NrInscEmpregador: nrInsc,
				CPFGestor:        cpf,
				MatriculaGestor:  matricula,
				ListaGeridos: []structs.ListaGerido{
					{
						CPF:              cpf,
						Matricula:        matricula,
						NrInscEmpregador: nrInsc,
						NomeFantasia:     "PERSI",
					},
				},
				Geridos: []structs.Gerido{
					{
						CPF:       cpf,
						Matricula: matricula,
					},
				},
			},
		},
	}
}

func GestoresRH(nrInsc, cpf, matricula string) structs.GestoresRH {
	return structs.GestoresRH{
		Gestores: []structs.GestorRH{
			{
				NrInscEmpregador: nrInsc,
				CPFGestor:        cpf,
				MatriculaGestor:  matricula,
				ListaGeridos: []structs.ListaGeridoRH{
					{
						CPF:                cpf,
						Matricula:          matricula,
						NrInscEmpregador:   nrInsc,
						NomeFantasia:       "PERSI",
						Cargo:              "Analista de Mídia",
						AreaOrganizacional: "ADMINISTRAÇÕ",
					},
				},
				Geridos: []structs.GeridoRH{
					{
						NrInscEmpregador: nrInsc,
						CPFGerido:        cpf,
						MatriculaGerido:  matricula,
						ListaGestores: []structs.ListaGestorRH{
							{
								NrInscEmpregador: nrInsc,
								Gestores: []string{
									cpf,
								},
							},
						},
					},
				},
			},
		},
	}
}
