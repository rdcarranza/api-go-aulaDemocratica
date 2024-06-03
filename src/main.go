package main

import (
	"api-go-aulaDemocratica/src/dominio"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	jornada := dominio.JornadaElectoral{
		ID:    uuid.New(),
		Fecha: time.Now(),
	}
	fmt.Print("Jornada creada! - ID: " + jornada.ID.(uuid.UUID).String() + " fecha: " + jornada.Fecha.String())
}
