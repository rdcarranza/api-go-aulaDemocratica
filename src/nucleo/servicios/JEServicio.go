package servicios

import (
	"api-go-aulaDemocratica/src/nucleo/dominio"
	"api-go-aulaDemocratica/src/nucleo/puertos"
	"time"

	"github.com/google/uuid"
)

type JEServicio struct {
}

// implementación de la interfaz
func Servicio_JornadaElectoral() puertos.JornadaElectoralServicio {
	return &JEServicio{}
}

// constructor
func NuevoServicio_JornadaElectora() *JEServicio {
	return &JEServicio{}
}

func (jes *JEServicio) AbrirJE(u string, dni_r int) (*dominio.JornadaElectoral, error) {
	jornada := dominio.JornadaElectoral{
		ID:    uuid.New(),
		Fecha: time.Now(),
		//Fecha: time.Now().UTC(),
		Ubicacion:      u,
		Dni_resposable: dni_r,
	}
	return &jornada, nil
}
func (jes *JEServicio) CerrarJE(jornada dominio.JornadaElectoral) (*dominio.JornadaElectoral, error) {
	return &jornada, nil
}
