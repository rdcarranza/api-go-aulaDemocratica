package repositorio_mongodb_je

import (
	"api-go-aulaDemocratica/src/nucleo/dominio"
	"context"
	"log"
)

func (r *Repositorio) ActualizarEstadoJE(je *dominio.JornadaElectoral) (b bool, err error) {
	collection := r.Cliente.Database("aula_democratica").Collection("jornadas_electorales")
	_, err = collection.UpdateOne(context.Background(), je.GetID(), je)
	if err != nil {
		log.Println(err.Error())
		//return nil, fmt.Errorf("Error al insertar Jornada Electoral: %w", err)
		return false, err
	}
	return true, nil
}
