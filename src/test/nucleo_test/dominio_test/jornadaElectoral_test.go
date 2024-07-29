// JornadaElectoral instance is created successfully
package dominio_test

import (
	"api-go-aulaDemocratica/src/nucleo/dominio"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJornadaElectoralInstanceCreatedSuccessfully(t *testing.T) {
	jornada := dominio.JornadaElectoral{}
	jornada.NuevaJornadaElectoral()

	assert.NotNil(t, jornada)
	assert.IsType(t, uuid.UUID{}, jornada.GetID())

	assert.WithinDuration(t, time.Now(), jornada.GetFecha_apertura(), time.Second)

}
