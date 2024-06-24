package config

import (
	"time"
)

var nrInsc string = "10821992"
var cpf string = "60515860409"

func DefinicoesRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Definicoes": []map[string]interface{}{
			{
				"NrInscEmpregador": nrInsc,
				"Ferias": map[string]interface{}{
					"AntecedenciaMinima":     15,
					"HabilitaFerias":         true,
					"ExigeAprovacaoDoGestor": true,
				},
			},
		},
	}
}

func PostInformacoesFeriasEmpregadoRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Colaboradores": []interface{}{
			map[string]interface{}{
				"CPF":                            cpf,
				"NrInscEmpregador":               "10821992",
				"Matricula":                      "000031",
				"SolicitouAdiantamento13":        false,
				"DiasDisponiveisAbonoPecuniario": 10,
				"DiasDisponiveis":                30,
				"InicioPeriodoConcessivo":        "2022-06-29",
				"FimPeriodoConcessivo":           "2023-07-30",
			},
		},
	}
}

func PostDispositivosRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Cnpj":          "63542443000124",
		"DispositivoId": "31e18fe4151b96cb",
		"Status":        1,
		"ListaEmpresas": []string{"10821992"},
	}
}

func PostSolicitaFeriasRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Ferias": []interface{}{
			map[string]interface{}{
				"CPF":                      cpf,
				"NrInscEmpregador":         nrInsc,
				"Matricula":                "000031",
				"SolicitouAdiantamento13":  true,
				"SolicitouAbonoPecuniario": true,
				"StatusSolicitacao":        4,
				"InicioPeriodoGozo":        time.Now().Format("2006-01-02"),
				"FimPeriodoGozo":           time.Now().AddDate(0, 0, 20).Format("2006-01-02"),
			},
		},
	}
}

func PostSolicitaFeriasAPPRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"NrInscEmpregador":         nrInsc,
		"Matricula":                "000031",
		"SolicitouAdiantamento13":  true,
		"SolicitouAbonoPecuniario": true,
		"InicioPeriodoGozo":        time.Now().Format("2006-01-02"),
		"FimPeriodoGozo":           time.Now().AddDate(0, 0, 20).Format("2006-01-02"),
	}
}

func PostAssinaturaRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"NrInscEmpregador": nrInsc,
		"AnoMes":           time.Now().AddDate(0, 0, 0).Format("200601"),
		"Liberado":         true,
	}
}

func DeleteAssinaturaRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"CPF":              "12658729375",
		"NrInscEmpregador": "10821992",
		"Matricula":        "000043",
		"AnoMes":           time.Now().AddDate(0, 0, 0).Format("200601"),
	}
}

func GestoresRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Gestores": []interface{}{
			map[string]interface{}{
				"NrInscEmpregador": nrInsc,
				"CPFGestor":        "12658729375",
				"MatriculaGestor":  "000043",
				"ListaGeridos": []interface{}{
					map[string]interface{}{
						"CPF":              "12658729375",
						"Matricula":        "000043",
						"NrInscEmpregador": nrInsc,
						"NomeFantasia":     "PERSI",
					},
				},
				"Geridos": []interface{}{
					map[string]interface{}{
						"CPF":       "12658729375",
						"Matricula": "000043",
					},
				},
			},
		},
	}
}

func GestoresRhRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"Gestores": []interface{}{
			map[string]interface{}{
				"NrInscEmpregador": nrInsc,
				"CPFGestor":        "12658729375",
				"MatriculaGestor":  "000043",
				"ListaGeridos": []interface{}{
					map[string]interface{}{
						"CPF":                "12658729375",
						"Matricula":          "000043",
						"NrInscEmpregador":   nrInsc,
						"NomeFantasia":       "PERSI",
						"Cargo":              "Analista de Mídia",
						"AreaOrganizacional": "ADMINISTRAÇÕ",
					},
				},
				"Geridos": []interface{}{
					map[string]interface{}{
						"NrInscEmpregador": nrInsc,
						"CPFGerido":        "12658729375",
						"MatriculaGerido":  "000043",
						"ListaGestores": []interface{}{
							map[string]interface{}{
								"NrInscEmpregador": nrInsc,
								"Gestores": []string{
									"12658729375",
								},
							},
						},
					},
				},
			},
		},
	}
}

func PostCandidatoBody(tax_id string, cnpj string) map[string]interface{} {
	return map[string]interface{}{
		"NrInscEmpregador": cnpj,
		"CPF":              tax_id,
	}
}

func PutReenviarCandidatoBody(tax_id string, cnpj string) map[string]interface{} {
	return map[string]interface{}{
		"NrInscEmpregador": cnpj,
		"CPF":              tax_id,
	}
}
