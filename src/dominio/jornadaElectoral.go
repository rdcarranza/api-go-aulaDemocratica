package dominio

import "time"

type JornadaElectoral struct {
	ID             interface{}
	Fecha          time.Time
	Ubicacion      string
	Dni_resposable int
}
