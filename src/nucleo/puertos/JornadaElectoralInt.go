package puertos

import (
	"api-go-aulaDemocratica/src/nucleo/dominio"
)

type JornadaElectoralServicio interface {
	AbrirJE(u string, dni_r int) (*dominio.JornadaElectoral, error)
	CerrarJE(je dominio.JornadaElectoral) (*dominio.JornadaElectoral, error)
}

type JornadaElectoralRepositorio interface {
	InsertarJE(je dominio.JornadaElectoral) (id interface{}, err error)
	ActualizarEstadoJE(je *dominio.JornadaElectoral) (b bool, err error)
}
