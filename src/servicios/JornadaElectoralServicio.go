package servicios

import (
	"api-go-aulaDemocratica/src/puertos"
)

type JEServicio struct {
	func abrirJE(je dominio.JornadaElectoral) (*dominio.JornadaElectoral, error){

		
	}
	func cerrarJE(je dominio.JornadaElectoral) (*dominio.JornadaElectoral, error){

	}
}

func servicio_JEServicio() puertos.JornadaElectoralServicio {
	return &JEServicio{}
}
