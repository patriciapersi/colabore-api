package structs

type TpInscEmpregador string

const (
	CNPJ TpInscEmpregador = "1"
	CPF  TpInscEmpregador = "2"
)

// STRUCT PARA INCLUSÃO DE MENSAGENS
type Mensagem struct {
	ID               string           `json:"ID"`
	TpInscEmpregador TpInscEmpregador `json:"TpInscEmpregador"`
	NrInscEmpregador string           `json:"NrInscEmpregador"`
	MensagemTitulo   string           `json:"MensagemTitulo"`
	MensagemCorpo    string           `json:"MensagemCorpo"`
	DataMensagem     string           `json:"DataMensagem"`
	Colaboradores    []Colaborador    `json:"Colaboradores"`
}

type Colaborador struct {
	CPF      string   `json:"CPF"`
	Contrato Contrato `json:"Contrato"`
}

type Contrato struct {
	Matricula string `json:"Matricula"`
	Cargo     string `json:"Cargo"`
}

// STRUCT PARA DELEÇÃO DE MENSAGENS
type DeleteMensagensRequest struct {
	MensagemID       string   `json:"MensagemId"`
	NrInscEmpregador string   `json:"NrInscEmpregador"`
	ListaCPF         []string `json:"ListaCPF"`
}
