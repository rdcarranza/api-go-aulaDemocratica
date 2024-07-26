package dominio

import "time"

type JornadaElectoral struct {
	ID             interface{}
	Fecha          time.Time
	Ubicacion      string
	Dni_resposable int
}

func New() JornadaElectoral {
	return JornadaElectoral{}
}

func guardar() bool {

	return true
}

func actualizar() bool {

	return true
}
