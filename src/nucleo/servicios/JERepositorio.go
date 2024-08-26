package servicios

import (
	"api-go-aulaDemocratica/src/nucleo/dominio"
	"api-go-aulaDemocratica/src/nucleo/puertos"
)

//ESTA CLASE CREO QUE DEBE SER BORRADA! los repositorios implementan la interfaz, es decir.. son los diferentes adaptadores que se conectan a los puertos. Es así que el repositorio de mongodb puede ser reemplazado por el repositorio de postgresql y continuar funcionando. O dependiendo de la funcionalidad, utilizar uno u otro repositorio.

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

func (jer *JERepositorio) ActualizarEstadoJE(je *dominio.JornadaElectoral) (b bool, err error) {
	//b_act: bandera de actualización.
	b_act, err := jer.Repo.ActualizarEstadoJE(je)
	if err != nil {
		return false, err
	}
	return b_act, nil
}
