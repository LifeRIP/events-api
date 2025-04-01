package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"events-api/internal/models"
)

func main() {
	// Conectar a MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	// Crear la colección de eventos
	collection := client.Database("events_db").Collection("events")

	// Eventos de ejemplo
	events := []interface{}{
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "Mantenimiento programado",
			Type:        models.TypeMaintenance,
			Description: "Mantenimiento programado del sistema para actualización",
			Date:        time.Now().AddDate(0, 0, 7),
			Status:      models.StatusPending,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "Alerta de seguridad",
			Type:        models.TypeAlert,
			Description: "Detección de posible intrusión en el sistema",
			Date:        time.Now().AddDate(0, 0, -2),
			Status:      models.StatusPending,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		models.Event{
			ID:              primitive.NewObjectID(),
			Name:            "Notificación de actualización",
			Type:            models.TypeNotification,
			Description:     "Nueva versión disponible",
			Date:            time.Now().AddDate(0, 0, 1),
			Status:          models.StatusReviewed,
			ManagementStatus: models.ManagementNotRequired,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		models.Event{
			ID:              primitive.NewObjectID(),
			Name:            "Emergencia de red",
			Type:            models.TypeEmergency,
			Description:     "Caída de la conexión principal",
			Date:            time.Now().AddDate(0, 0, -1),
			Status:          models.StatusReviewed,
			ManagementStatus: models.ManagementRequired,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "Información de usuario",
			Type:        models.TypeInfo,
			Description: "Actualización de información de usuarios",
			Date:        time.Now(),
			Status:      models.StatusPending,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	// Insertar eventos
	_, err = collection.InsertMany(ctx, events)
	if err != nil {
		log.Fatalf("Error al insertar eventos: %v", err)
	}

	log.Println("Eventos de ejemplo generados con éxito")
}

