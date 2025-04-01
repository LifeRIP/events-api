package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EventStatus es un tipo para representar el estado del evento
type EventStatus string

// ManagementStatus es un tipo para representar el estado de gestión
type ManagementStatus string

// EventType es un tipo para representar el tipo de evento
type EventType string

const (
	// Estados del evento
	StatusPending  EventStatus = "PENDING"
	StatusReviewed EventStatus = "REVIEWED"

	// Estados de gestión
	ManagementRequired    ManagementStatus = "REQUIRES_MANAGEMENT"
	ManagementNotRequired ManagementStatus = "NO_MANAGEMENT"

	// Tipos de evento (ejemplo)
	TypeEmergency    EventType = "EMERGENCY"
	TypeMaintenance  EventType = "MAINTENANCE"
	TypeNotification EventType = "NOTIFICATION"
	TypeAlert        EventType = "ALERT"
	TypeInfo         EventType = "INFO"
)

// Event representa la estructura de un evento
type Event struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name             string             `json:"name" bson:"name" binding:"required"`
	Type             EventType          `json:"type" bson:"type" binding:"required"`
	Description      string             `json:"description" bson:"description" binding:"required"`
	Date             time.Time          `json:"date" bson:"date"`
	Status           EventStatus        `json:"status" bson:"status"`
	ManagementStatus ManagementStatus   `json:"managementStatus,omitempty" bson:"management_status,omitempty"`
	CreatedAt        time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updatedAt" bson:"updated_at"`
}

// CreateEventRequest representa la solicitud para crear un evento
type CreateEventRequest struct {
	Name        string    `json:"name" example:"Mantenimiento programado" binding:"required"`
	Type        EventType `json:"type" example:"MAINTENANCE" binding:"required"`
	Description string    `json:"description" example:"Mantenimiento programado del sistema para actualización" binding:"required"`
	Date        time.Time `json:"date" example:"2025-04-08T00:00:00Z" binding:"required"`
}

// UpdateEventRequest representa la solicitud para actualizar un evento
type UpdateEventRequest struct {
	Name        string    `json:"name" example:"Actualización de mantenimiento"`
	Type        EventType `json:"type" example:"MAINTENANCE"`
	Description string    `json:"description" example:"Actualización de la descripción del mantenimiento programado"`
	Date        time.Time `json:"date" example:"2025-04-15T00:00:00Z"`
}

// ReviewEventRequest representa la solicitud para revisar un evento
type ReviewEventRequest struct {
	ManagementStatus ManagementStatus `json:"managementStatus" binding:"required"`
}

// EventResponse representa la respuesta de un evento
type EventResponse struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	Description      string    `json:"description"`
	Date             time.Time `json:"date"`
	Status           string    `json:"status"`
	ManagementStatus string    `json:"managementStatus,omitempty"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
