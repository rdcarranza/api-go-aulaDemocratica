package repositorio_mongodb_je

import (
	"api-go-aulaDemocratica/src/nucleo/puertos"

	"go.mongodb.org/mongo-driver/mongo"
)

var _ puertos.JornadaElectoralRepositorio = &Repositorio{}

type Repositorio struct {
	Cliente   *mongo.Client
	Coleccion *mongo.Collection
}
