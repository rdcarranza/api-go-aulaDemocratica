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
	jornada := dominio.JornadaElectoral{
		ID:    uuid.New(),
		Fecha: time.Now(),
	}

	assert.NotNil(t, jornada)
	assert.IsType(t, uuid.UUID{}, jornada.ID)

	assert.WithinDuration(t, time.Now(), jornada.Fecha, time.Second)

}

func TestJornadaElectoralIDIsNotValidUUID(t *testing.T) {
	invalidID := "invalid-uuid"
	jornada := dominio.JornadaElectoral{
		ID:    invalidID,
		Fecha: time.Now(),
	}

	_, err := uuid.Parse(jornada.ID.(string))
	assert.Error(t, err)
}
