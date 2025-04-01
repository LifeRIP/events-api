package repositories

import (
	"context"
	"errors"
	"time"

	"events-api/internal/config"
	"events-api/internal/models"
	"events-api/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// EventRepository define las operaciones del repositorio de eventos
type EventRepository interface {
	FindAll(ctx context.Context) ([]models.Event, error)
	FindByID(ctx context.Context, id string) (models.Event, error)
	Create(ctx context.Context, event models.Event) (models.Event, error)
	Update(ctx context.Context, id string, event models.Event) (models.Event, error)
	Delete(ctx context.Context, id string) error
	FindByStatus(ctx context.Context, status models.EventStatus) ([]models.Event, error)
	FindByManagementStatus(ctx context.Context, managementStatus models.ManagementStatus) ([]models.Event, error)
	BulkInsert(ctx context.Context, events []models.Event) error
}

// eventRepository implementa EventRepository
type eventRepository struct {
	collection *mongo.Collection
}

// NewEventRepository crea una nueva instancia de EventRepository
func NewEventRepository(client *mongo.Client, cfg *config.Config) EventRepository {
	collection := database.GetCollection(client, cfg, cfg.EventsCollection)
	return &eventRepository{
		collection: collection,
	}
}

// FindAll recupera todos los eventos
func (r *eventRepository) FindAll(ctx context.Context) ([]models.Event, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []models.Event
	if err := cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}

// FindByID recupera un evento por su ID
func (r *eventRepository) FindByID(ctx context.Context, id string) (models.Event, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Event{}, err
	}

	var event models.Event
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&event)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Event{}, errors.New("evento no encontrado")
		}
		return models.Event{}, err
	}

	return event, nil
}

// Create crea un nuevo evento
func (r *eventRepository) Create(ctx context.Context, event models.Event) (models.Event, error) {
	now := time.Now()

	event.CreatedAt = now
	event.UpdatedAt = now
	event.Status = models.StatusPending

	if event.ID.IsZero() {
		event.ID = primitive.NewObjectID()
	}

	_, err := r.collection.InsertOne(ctx, event)
	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}

// Update actualiza un evento existente
func (r *eventRepository) Update(ctx context.Context, id string, event models.Event) (models.Event, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Event{}, err
	}

	event.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"name":              event.Name,
			"type":              event.Type,
			"description":       event.Description,
			"date":              event.Date,
			"status":            event.Status,
			"management_status": event.ManagementStatus,
			"updated_at":        event.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return models.Event{}, err
	}

	if result.MatchedCount == 0 {
		return models.Event{}, errors.New("evento no encontrado")
	}

	return r.FindByID(ctx, id)
}

// Delete elimina un evento
func (r *eventRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("evento no encontrado")
	}

	return nil
}

// FindByStatus recupera eventos por su estado
func (r *eventRepository) FindByStatus(ctx context.Context, status models.EventStatus) ([]models.Event, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{"status": status}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []models.Event
	if err := cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}

// FindByManagementStatus recupera eventos por su estado de gestión
func (r *eventRepository) FindByManagementStatus(ctx context.Context, managementStatus models.ManagementStatus) ([]models.Event, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	filter := bson.M{
		"status":            models.StatusReviewed,
		"management_status": managementStatus,
	}

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []models.Event
	if err := cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}

// BulkInsert inserta múltiples eventos
func (r *eventRepository) BulkInsert(ctx context.Context, events []models.Event) error {
	if len(events) == 0 {
		return nil
	}

	now := time.Now()
	var documents []interface{}

	for i := range events {
		events[i].CreatedAt = now
		events[i].UpdatedAt = now
		events[i].Status = models.StatusPending

		if events[i].ID.IsZero() {
			events[i].ID = primitive.NewObjectID()
		}

		documents = append(documents, events[i])
	}

	_, err := r.collection.InsertMany(ctx, documents)
	return err
}
