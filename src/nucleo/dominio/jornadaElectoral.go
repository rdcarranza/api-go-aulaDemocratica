package dominio

import (
	"time"

	"github.com/google/uuid"
)

type JornadaElectoral struct {
	id             uuid.UUID
	fecha_apertura time.Time
	fecha_cierre   time.Time
	ubicacion      string
	dni_resposable int
	estado         bool
}

// constructor
func (je *JornadaElectoral) NuevaJornadaElectoral() {
	je.id = uuid.New()
	je.estado = false
}

func (*JornadaElectoral) Guardar() bool {
	//este método, debe estar controlado como un servicio
	return true
}

func (*JornadaElectoral) Actualizar() bool {
	//este método, debe estar controlado como un servicio
	return true
}

func (je *JornadaElectoral) GetID() uuid.UUID {
	return je.id
}

/*
	func (je *JornadaElectoral) SetID(uuid uuid.UUID) {
		je.id = uuid
	}
*/
func (je *JornadaElectoral) GetFecha_apertura() time.Time {
	return je.fecha_apertura
}

func (je *JornadaElectoral) SetFecha_apertura(t time.Time) {
	je.fecha_apertura = t
}

func (je *JornadaElectoral) GetFecha_cierre() time.Time {
	return je.fecha_cierre
}

func (je *JornadaElectoral) SetFecha_cierre(t time.Time) {
	je.fecha_cierre = t
}

func (je *JornadaElectoral) GetUbicacion() string {
	return je.ubicacion
}

func (je *JornadaElectoral) SetUbicacion(u string) {
	je.ubicacion = u
}

func (je *JornadaElectoral) GetDni_responsable() int {
	return je.dni_resposable
}

func (je *JornadaElectoral) SetDni_responsable(dni int) {
	je.dni_resposable = dni
}

func (je *JornadaElectoral) GetEstado() bool {
	return je.estado
}

func (je *JornadaElectoral) SetEstado_abrir() {
	je.estado = true
}

func (je *JornadaElectoral) SetEstado_cerrar() {
	je.estado = false
}
