package puertos

import (
	"api-go-aulaDemocratica/src/dominio"
)

type JornadaElectoralServicio interface {
	AbrirJE(u string, dni_r int) (*dominio.JornadaElectoral, error)
	CerrarJE(je dominio.JornadaElectoral) (*dominio.JornadaElectoral, error)
}

type JornadaElectoralRepositorio interface {
	//conexionCliente(db_url string) (cliente *mongo.Client, err error)
	InsertarJE(je dominio.JornadaElectoral) (id interface{}, err error)
}
