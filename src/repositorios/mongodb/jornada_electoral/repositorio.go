package repositorio_mongodb_je

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositorio struct {
	Cliente   *mongo.Client
	Coleccion *mongo.Collection
}
