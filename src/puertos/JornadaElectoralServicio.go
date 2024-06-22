package puertos

import "api-go-aulaDemocratica/src/dominio"

type JornadaElectoralServicio interface {
	abrirJE(je dominio.JornadaElectoral) (*dominio.JornadaElectoral, error)
	cerrarJE(je dominio.JornadaElectoral) (*dominio.JornadaElectoral, error)
}
