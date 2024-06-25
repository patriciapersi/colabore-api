package structs

type Ferias struct {
	AntecedenciaMinima     int  `json:"AntecedenciaMinima"`
	HabilitaFerias         bool `json:"HabilitaFerias"`
	ExigeAprovacaoDoGestor bool `json:"ExigeAprovacaoDoGestor"`
}

type DefinicoesEmpresa struct {
	NrInscEmpregador string `json:"NrInscEmpregador"`
	Ferias           Ferias `json:"Ferias"`
}

type Definicoes struct {
	DefinicoesEmpresa []DefinicoesEmpresa `json:"Definicoes"`
}
