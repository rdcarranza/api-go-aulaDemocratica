package servicios

import (
	controlador_error "api-go-aulaDemocratica/src/controladores/error"
	"api-go-aulaDemocratica/src/nucleo/dominio"
	"api-go-aulaDemocratica/src/nucleo/puertos"
	"fmt"
	"time"
)

type JEServicio struct {
	jornada dominio.JornadaElectoral
}

// implementación de la interfaz
func servicio_JornadaElectoral() puertos.JornadaElectoralServicio {
	return &JEServicio{}
}

// constructor
func NuevoServicio_JornadaElectora() *JEServicio {
	je := servicio_JornadaElectoral().(*JEServicio) //casting, forzando el tipo de dato
	je.jornada.NuevaJornadaElectoral()
	return je
}

func (jes *JEServicio) AbrirJE(u string, dni_r int) (*dominio.JornadaElectoral, error) {

	jes.jornada.SetFecha_apertura(time.Now())
	jes.jornada.SetUbicacion(u)
	jes.jornada.SetDni_responsable(dni_r)
	jes.jornada.SetEstado_abrir()
	return &jes.jornada, nil
}
func (jes *JEServicio) CerrarJE(jornada dominio.JornadaElectoral) (*dominio.JornadaElectoral, error) {
	if time.Time.Compare(jes.jornada.GetFecha_cierre(), time.Time{}) == 0 && jes.jornada.GetEstado() {
		jes.jornada.SetFecha_cierre(time.Now())
		jes.jornada.SetEstado_cerrar()
	} else {
		ce := controlador_error.NuevoControladorError()
		ce.Archivo = "JEServicios"
		ce.Modulo = "CerrarJE: la jornada que intenta cerrarse, ya se encuentra cerrada."
		me := ce.Error()
		fmt.Println(me)
		return nil, ce
	}
	return &jornada, nil
}

func (jes *JEServicio) VerificarJE() bool {
	return true
}

func (jes *JEServicio) GetJornada() *dominio.JornadaElectoral {
	return &jes.jornada
}

func (jes *JEServicio) SetJornada(j dominio.JornadaElectoral) {
	jes.jornada = j
}
