package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Objeto de conexion a la base de datos */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://david:1069925@cluster0.vdi3g.mongodb.net/<dbname>?retryWrites=true&w=majority")

/*Conectar bd es la funcion que me permite conectar a la base de datos */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Printf(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Printf(err.Error())
		return client
	}
	log.Printf("Conexion exitosaa")
	return client

}

/*Cheque conection ping a base de datos */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
