package structs

type ColaboradoresInfo struct {
	CPF                            string `json:"CPF"`
	NrInscEmpregador               string `json:"NrInscEmpregador"`
	Matricula                      string `json:"Matricula"`
	SolicitouAdiantamento13        bool   `json:"SolicitouAdiantamento13"`
	DiasDisponiveisAbonoPecuniario int    `json:"DiasDisponiveisAbonoPecuniario"`
	DiasDisponiveis                int    `json:"DiasDisponiveis"`
	InicioPeriodoConcessivo        string `json:"InicioPeriodoConcessivo"`
	FimPeriodoConcessivo           string `json:"FimPeriodoConcessivo"`
}

type ColaboradorRequestBody struct {
	Colaboradores []ColaboradoresInfo `json:"Colaboradores"`
}

type FeriasInfo struct {
	CPF                      string `json:"CPF"`
	NrInscEmpregador         string `json:"NrInscEmpregador"`
	Matricula                string `json:"Matricula"`
	SolicitouAdiantamento13  bool   `json:"SolicitouAdiantamento13"`
	SolicitouAbonoPecuniario bool   `json:"SolicitouAbonoPecuniario"`
	StatusSolicitacao        int    `json:"StatusSolicitacao"`
	InicioPeriodoGozo        string `json:"InicioPeriodoGozo"`
	FimPeriodoGozo           string `json:"FimPeriodoGozo"`
}

type Ferias struct {
	Ferias []FeriasInfo `json:"Ferias"`
}
