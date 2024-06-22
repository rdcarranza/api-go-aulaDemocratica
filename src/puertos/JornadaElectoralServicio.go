package puertos

import "api-go-aulaDemocratica/src/dominio"

type JornadaElectoralServicio interface {
	AbrirJE() (*dominio.JornadaElectoral, error)
	CerrarJE(je dominio.JornadaElectoral) (*dominio.JornadaElectoral, error)
}
