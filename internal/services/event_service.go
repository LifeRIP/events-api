package services

import (
	"context"
	"errors"
	"time"

	"events-api/internal/models"
	"events-api/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EventService define las operaciones del servicio de eventos
type EventService interface {
	GetAllEvents(ctx context.Context) ([]models.EventResponse, error)
	GetEventByID(ctx context.Context, id string) (models.EventResponse, error)
	CreateEvent(ctx context.Context, req models.CreateEventRequest) (models.EventResponse, error)
	UpdateEvent(ctx context.Context, id string, req models.UpdateEventRequest) (models.EventResponse, error)
	DeleteEvent(ctx context.Context, id string) error
	ReviewEvent(ctx context.Context, id string, req models.ReviewEventRequest) (models.EventResponse, error)
	UnreviewEvent(ctx context.Context, id string) (models.EventResponse, error)
	GetEventTypes(ctx context.Context) []string
	GetEventStatus(ctx context.Context) []string
	GetEventManagementStatus(ctx context.Context) []string
	SeedEvents(ctx context.Context) error
	GetEventsRequiringManagement(ctx context.Context) ([]models.EventResponse, error)
	GetEventsNotRequiringManagement(ctx context.Context) ([]models.EventResponse, error)
}

// eventService implementa EventService
type eventService struct {
	repository repositories.EventRepository
}

// NewEventService crea una nueva instancia de EventService
func NewEventService(repository repositories.EventRepository) EventService {
	return &eventService{
		repository: repository,
	}
}

// GetAllEvents recupera todos los eventos
func (s *eventService) GetAllEvents(ctx context.Context) ([]models.EventResponse, error) {
	events, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []models.EventResponse
	for _, event := range events {
		responses = append(responses, mapEventToResponse(event))
	}

	return responses, nil
}

// GetEventByID recupera un evento por su ID
func (s *eventService) GetEventByID(ctx context.Context, id string) (models.EventResponse, error) {
	event, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return models.EventResponse{}, err
	}

	return mapEventToResponse(event), nil
}

// CreateEvent crea un nuevo evento
func (s *eventService) CreateEvent(ctx context.Context, req models.CreateEventRequest) (models.EventResponse, error) {
	// Validar tipo de evento
	if !isValidEventType(req.Type) {
		return models.EventResponse{}, errors.New("tipo de evento no válido")
	}

	event := models.Event{
		Name:        req.Name,
		Type:        req.Type,
		Description: req.Description,
		Date:        req.Date,
		Status:      models.StatusPending,
	}

	createdEvent, err := s.repository.Create(ctx, event)
	if err != nil {
		return models.EventResponse{}, err
	}

	return mapEventToResponse(createdEvent), nil
}

// UpdateEvent actualiza un evento existente
func (s *eventService) UpdateEvent(ctx context.Context, id string, req models.UpdateEventRequest) (models.EventResponse, error) {
	existingEvent, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return models.EventResponse{}, err
	}

	// Actualizar solo los campos proporcionados
	if req.Name != "" {
		existingEvent.Name = req.Name
	}

	if req.Type != "" {
		if !isValidEventType(req.Type) {
			return models.EventResponse{}, errors.New("tipo de evento no válido")
		}
		existingEvent.Type = req.Type
	}

	if req.Description != "" {
		existingEvent.Description = req.Description
	}

	if !req.Date.IsZero() {
		existingEvent.Date = req.Date
	}

	updatedEvent, err := s.repository.Update(ctx, id, existingEvent)
	if err != nil {
		return models.EventResponse{}, err
	}

	return mapEventToResponse(updatedEvent), nil
}

// DeleteEvent elimina un evento
func (s *eventService) DeleteEvent(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

// ReviewEvent revisa un evento y asigna un estado de gestión automáticamente
func (s *eventService) ReviewEvent(ctx context.Context, id string, req models.ReviewEventRequest) (models.EventResponse, error) {
	existingEvent, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return models.EventResponse{}, err
	}

	// Determinar automáticamente el estado de gestión basado en el tipo de evento
	var managementStatus models.ManagementStatus
	switch existingEvent.Type {
	case models.TypeEmergency, models.TypeAlert:
		managementStatus = models.ManagementRequired
	case models.TypeMaintenance, models.TypeNotification, models.TypeInfo:
		managementStatus = models.ManagementNotRequired
	default:
		return models.EventResponse{}, errors.New("tipo de evento no reconocido")
	}

	// Actualizar estado y estado de gestión
	existingEvent.Status = models.StatusReviewed
	existingEvent.ManagementStatus = managementStatus

	updatedEvent, err := s.repository.Update(ctx, id, existingEvent)
	if err != nil {
		return models.EventResponse{}, err
	}

	return mapEventToResponse(updatedEvent), nil
}

// UnreviewEvent revierte la revisión de un evento
func (s *eventService) UnreviewEvent(ctx context.Context, id string) (models.EventResponse, error) {
	existingEvent, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return models.EventResponse{}, err
	}

	// Verificar que el evento esté en estado revisado
	if existingEvent.Status != models.StatusReviewed {
		return models.EventResponse{}, errors.New("el evento no está en estado revisado")
	}

	// Cambiar el estado a pendiente y eliminar el estado de gestión
	existingEvent.Status = models.StatusPending
	existingEvent.ManagementStatus = "" // Eliminar el estado de gestión

	updatedEvent, err := s.repository.Update(ctx, id, existingEvent)
	if err != nil {
		return models.EventResponse{}, err
	}

	return mapEventToResponse(updatedEvent), nil
}

// GetEventTypes devuelve los tipos de eventos disponibles
func (s *eventService) GetEventTypes(ctx context.Context) []string {
	return []string{
		string(models.TypeEmergency),
		string(models.TypeMaintenance),
		string(models.TypeNotification),
		string(models.TypeAlert),
		string(models.TypeInfo),
	}
}

// GetEventStatus devuelve los estados de eventos disponibles
func (s *eventService) GetEventStatus(ctx context.Context) []string {
	return []string{
		string(models.StatusPending),
		string(models.StatusReviewed),
	}
}

// GetEventManagementStatus devuelve los estados de gestión disponibles
func (s *eventService) GetEventManagementStatus(ctx context.Context) []string {
	return []string{
		string(models.ManagementRequired),
		string(models.ManagementNotRequired),
	}
}

// SeedEvents genera eventos de ejemplo
func (s *eventService) SeedEvents(ctx context.Context) error {
	// Crear algunos eventos de ejemplo
	events := []models.Event{
		{
			ID:          primitive.NewObjectID(),
			Name:        "Mantenimiento programado",
			Type:        models.TypeMaintenance,
			Description: "Mantenimiento programado del sistema para actualización",
			Date:        time.Now().AddDate(0, 0, 7),
			Status:      models.StatusPending,
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Alerta de seguridad",
			Type:        models.TypeAlert,
			Description: "Detección de posible intrusión en el sistema",
			Date:        time.Now().AddDate(0, 0, -2),
			Status:      models.StatusPending,
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Notificación de actualización",
			Type:        models.TypeNotification,
			Description: "Nueva versión disponible",
			Date:        time.Now().AddDate(0, 0, 1),
			Status:      models.StatusPending,
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Emergencia de red",
			Type:        models.TypeEmergency,
			Description: "Caída de la conexión principal",
			Date:        time.Now().AddDate(0, 0, -1),
			Status:      models.StatusPending,
		},
		{
			ID:               primitive.NewObjectID(),
			Name:             "Información de usuario",
			Type:             models.TypeInfo,
			Description:      "Actualización de información de usuarios",
			Date:             time.Now(),
			Status:           models.StatusReviewed,
			ManagementStatus: models.ManagementNotRequired,
		},
	}

	return s.repository.BulkInsert(ctx, events)
}

// GetEventsRequiringManagement recupera eventos que requieren gestión
func (s *eventService) GetEventsRequiringManagement(ctx context.Context) ([]models.EventResponse, error) {
	events, err := s.repository.FindByManagementStatus(ctx, models.ManagementRequired)
	if err != nil {
		return nil, err
	}

	var responses []models.EventResponse
	for _, event := range events {
		responses = append(responses, mapEventToResponse(event))
	}

	return responses, nil
}

// GetEventsNotRequiringManagement recupera eventos que no requieren gestión
func (s *eventService) GetEventsNotRequiringManagement(ctx context.Context) ([]models.EventResponse, error) {
	events, err := s.repository.FindByManagementStatus(ctx, models.ManagementNotRequired)
	if err != nil {
		return nil, err
	}

	var responses []models.EventResponse
	for _, event := range events {
		responses = append(responses, mapEventToResponse(event))
	}

	return responses, nil
}

// mapEventToResponse mapea un Event a un EventResponse
func mapEventToResponse(event models.Event) models.EventResponse {
	return models.EventResponse{
		ID:               event.ID.Hex(),
		Name:             event.Name,
		Type:             string(event.Type),
		Description:      event.Description,
		Date:             event.Date,
		Status:           string(event.Status),
		ManagementStatus: string(event.ManagementStatus),
		CreatedAt:        event.CreatedAt,
		UpdatedAt:        event.UpdatedAt,
	}
}

// isValidEventType valida si un tipo de evento es válido
func isValidEventType(eventType models.EventType) bool {
	validTypes := map[models.EventType]bool{
		models.TypeEmergency:    true,
		models.TypeMaintenance:  true,
		models.TypeNotification: true,
		models.TypeAlert:        true,
		models.TypeInfo:         true,
	}

	return validTypes[eventType]
}
