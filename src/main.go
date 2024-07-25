package main

import (
	"api-go-aulaDemocratica/src/servicios"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

func main() {
	/*
		jornada := dominio.JornadaElectoral{
			ID:    uuid.New(),
			Fecha: time.Now(),
		}
		fmt.Print("Jornada creada! - ID: " + jornada.ID.(uuid.UUID).String() + " fecha: " + jornada.Fecha.String())
	*/

	//jornada := servicios.JEServicio(dominio.JornadaElectoral().New())

	//jes := servicios.JEServicio{}
	jes := servicios.Servicio_JornadaElectoral()

	jornada, err := jes.AbrirJE("escuela", 11111111)
	if err == nil {
		fmt.Print("Jornada creada! - ID: " + jornada.ID.(uuid.UUID).String() + " fecha: " + jornada.Fecha.String() + " - ubicación: " + jornada.Ubicacion + " - DNI del responsable: " + strconv.Itoa(jornada.Dni_resposable))
	}

}
