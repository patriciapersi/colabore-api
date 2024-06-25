package structs

type FeriasConfig struct {
	AntecedenciaMinima     int  `json:"AntecedenciaMinima"`
	HabilitaFerias         bool `json:"HabilitaFerias"`
	ExigeAprovacaoDoGestor bool `json:"ExigeAprovacaoDoGestor"`
}

type DefinicoesEmpresa struct {
	NrInscEmpregador string       `json:"NrInscEmpregador"`
	Ferias           FeriasConfig `json:"Ferias"`
}

type Definicoes struct {
	DefinicoesEmpresa []DefinicoesEmpresa `json:"Definicoes"`
}
