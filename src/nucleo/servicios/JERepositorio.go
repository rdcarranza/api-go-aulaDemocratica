package servicios

import (
	"api-go-aulaDemocratica/src/nucleo/dominio"
	"api-go-aulaDemocratica/src/nucleo/puertos"
)

type JERepositorio struct {
	Repo puertos.JornadaElectoralRepositorio
}

//var _ puertos.JornadaElectoralRepositorio = (*JERepositorio)(nil)

// Implementación de interfaz
func Repositorio_JornadaElectoral() puertos.JornadaElectoralRepositorio {
	return &JERepositorio{}
}

// constructor
func NuevoRepositorio_JornadaElectoral(repositorio puertos.JornadaElectoralRepositorio) *JERepositorio {
	return &JERepositorio{
		Repo: repositorio,
	}
}

func (jer *JERepositorio) InsertarJE(je dominio.JornadaElectoral) (id interface{}, err error) {
	id_int, err := jer.Repo.InsertarJE(je)
	if err != nil {
		return nil, err
	}
	return id_int, nil
}
