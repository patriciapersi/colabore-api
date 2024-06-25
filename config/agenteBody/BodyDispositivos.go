package agentebody

import (
	"github.com/google/uuid"
	"github.com/patriciapersi/colabore-api/config/structs"
)

func Dispositivo(nrInsc string) structs.Dispositivos {
	return structs.Dispositivos{
		Cnpj:          nrInsc,
		DispositivoId: uuid.New().String(),
		Status:        1,
		ListaEmpresas: nrInsc,
	}
}
