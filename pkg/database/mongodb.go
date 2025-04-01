package database

import (
	"context"
	"time"

	"events-api/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewMongoClient crea y retorna un nuevo cliente de MongoDB
func NewMongoClient(cfg *config.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Configurar opciones del cliente
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	// Conectar al servidor de MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Verificar la conexión
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}

// GetCollection obtiene una colección de MongoDB
func GetCollection(client *mongo.Client, cfg *config.Config, collectionName string) *mongo.Collection {
	return client.Database(cfg.MongoDatabase).Collection(collectionName)
}
