package servicios

import (
	"api-go-aulaDemocratica/src/dominio"
	"api-go-aulaDemocratica/src/puertos"
	"time"

	"github.com/google/uuid"
)

type JEServicio struct {
}

func servicio_JEServicio() puertos.JornadaElectoralServicio {
	return &JEServicio{}
}

func (jes *JEServicio) AbrirJE() (*dominio.JornadaElectoral, error) {
	jornada := dominio.JornadaElectoral{
		ID:    uuid.New(),
		Fecha: time.Now(),
	}
	return &jornada, nil
}
func (jes *JEServicio) CerrarJE(jornada dominio.JornadaElectoral) (*dominio.JornadaElectoral, error) {
	return &jornada, nil
}
