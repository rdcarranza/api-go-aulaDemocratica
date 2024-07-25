package repositorio_mongodb_je

import (
	"api-go-aulaDemocratica/src/dominio"
	"context"
	"log"
)

func (r Repositorio) insertar(je dominio.JornadaElectoral) (id interface{}, err error) {
	collection := r.Cliente.Database("aula_democratica").Collection("jornadas_electorales")
	insResultado, err := collection.InsertOne(context.Background(), je)
	if err != nil {
		log.Println(err.Error())
		//return nil, fmt.Errorf("Error al insertar Jornada Electoral: %w", err)
		return nil, err
	}
	return insResultado, nil
}
