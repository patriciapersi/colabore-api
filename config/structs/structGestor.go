package structs

type Gerido struct {
	CPF       string `json:"CPF"`
	Matricula string `json:"Matricula"`
}

type ListaGerido struct {
	CPF              string `json:"CPF"`
	Matricula        string `json:"Matricula"`
	NrInscEmpregador string `json:"NrInscEmpregador"`
	NomeFantasia     string `json:"NomeFantasia"`
}

type Gestor struct {
	NrInscEmpregador string        `json:"NrInscEmpregador"`
	CPFGestor        string        `json:"CPFGestor"`
	MatriculaGestor  string        `json:"MatriculaGestor"`
	ListaGeridos     []ListaGerido `json:"ListaGeridos"`
	Geridos          []Gerido      `json:"Geridos"`
}

type Gestores struct {
	Gestores []Gestor `json:"Gestores"`
}

//-----------------------
//GESTOR RH

type ListaGestorRH struct {
	NrInscEmpregador string   `json:"NrInscEmpregador"`
	Gestores         []string `json:"Gestores"`
}

type GeridoRH struct {
	NrInscEmpregador string          `json:"NrInscEmpregador"`
	CPFGerido        string          `json:"CPFGerido"`
	MatriculaGerido  string          `json:"MatriculaGerido"`
	ListaGestores    []ListaGestorRH `json:"ListaGestores"`
}

type ListaGeridoRH struct {
	CPF                string `json:"CPF"`
	Matricula          string `json:"Matricula"`
	NrInscEmpregador   string `json:"NrInscEmpregador"`
	NomeFantasia       string `json:"NomeFantasia"`
	Cargo              string `json:"Cargo"`
	AreaOrganizacional string `json:"AreaOrganizacional"`
}

type GestorRH struct {
	NrInscEmpregador string          `json:"NrInscEmpregador"`
	CPFGestor        string          `json:"CPFGestor"`
	MatriculaGestor  string          `json:"MatriculaGestor"`
	ListaGeridos     []ListaGeridoRH `json:"ListaGeridos"`
	Geridos          []GeridoRH      `json:"Geridos"`
}

type GestoresRH struct {
	Gestores []GestorRH `json:"Gestores"`
}
