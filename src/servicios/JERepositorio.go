package servicios

import (
	"api-go-aulaDemocratica/src/dominio"
	"api-go-aulaDemocratica/src/puertos"
)

type JERepositorio struct {
	Repo puertos.JornadaElectoralRepositorio
}

func Repositorio_JornadaElectoral() puertos.JornadaElectoralRepositorio {
	return &JERepositorio{}
}

func (jer *JERepositorio) InsertarJE(je dominio.JornadaElectoral) (id interface{}, err error) {
	id_int, err := jer.Repo.InsertarJE(je)
	if err != nil {
		return nil, err
	}
	return id_int, nil
}
