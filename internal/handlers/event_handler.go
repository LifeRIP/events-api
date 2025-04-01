package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"events-api/internal/models"
	"events-api/internal/services"
)

// EventHandler maneja las solicitudes HTTP relacionadas con eventos
type EventHandler struct {
	service services.EventService
}

// NewEventHandler crea una nueva instancia de EventHandler
func NewEventHandler(service services.EventService) *EventHandler {
	return &EventHandler{
		service: service,
	}
}

// CreateEvent godoc
//
//	@Summary		Crear un nuevo evento
//	@Description	Crea un nuevo evento con la información proporcionada
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			event	body		models.CreateEventRequest	true	"Información del evento"
//	@Success		201		{object}	models.EventResponse
//	@Failure		400		{object}	models.ErrorResponse	"Error en los datos de entrada"
//	@Failure		500		{object}	models.ErrorResponse	"Error interno del servidor"
//	@Router			/events [post]
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var req models.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := h.service.CreateEvent(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, event)
}

// GetAllEvents godoc
//
//	@Summary		Obtener todos los eventos
//	@Description	Obtiene una lista de todos los eventos
//	@Tags			events
//	@Produce		json
//	@Success		200	{array}		models.EventResponse
//	@Failure		404	{object}	models.ErrorResponse	"No se encontraron eventos"
//	@Failure		500	{object}	models.ErrorResponse	"Error interno del servidor"
//	@Router			/events [get]
func (h *EventHandler) GetAllEvents(c *gin.Context) {
	events, err := h.service.GetAllEvents(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Si no se encuentran eventos, retorna un 404
	if len(events) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no se encontraron eventos"})
		return
	}

	c.JSON(http.StatusOK, events)
}

// GetEventByID godoc
//
//	@Summary		Obtener un evento por ID
//	@Description	Obtiene un evento por su ID
//	@Tags			events
//	@Produce		json
//	@Param			id	path		string	true	"ID del evento"
//	@Success		200	{object}	models.EventResponse
//	@Failure		404	{object}	models.ErrorResponse	"Evento no encontrado"
//	@Failure		500	{object}	models.ErrorResponse	"Error interno del servidor"
//	@Router			/events/{id} [get]
func (h *EventHandler) GetEventByID(c *gin.Context) {
	id := c.Param("id")
	event, err := h.service.GetEventByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

// UpdateEvent godoc
//
//	@Summary		Actualizar un evento
//	@Description	Actualiza un evento existente
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"ID del evento"
//	@Param			event	body		models.UpdateEventRequest	true	"Información del evento"
//	@Success		200		{object}	models.EventResponse
//	@Failure		400		{object}	models.ErrorResponse	"Error en los datos de entrada"
//	@Failure		404		{object}	models.ErrorResponse	"Evento no encontrado"
//	@Failure		500		{object}	models.ErrorResponse	"Error interno del servidor"
//	@Router			/events/{id} [put]
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var req models.UpdateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := h.service.UpdateEvent(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

// DeleteEvent godoc
//
//	@Summary		Eliminar un evento
//	@Description	Elimina un evento existente
//	@Tags			events
//	@Param			id	path		string	true	"ID del evento"
//	@Success		204	{object}	nil
//	@Failure		404	{object}	models.ErrorResponse	"Evento no encontrado"
//	@Failure		500	{object}	models.ErrorResponse	"Error interno del servidor"
//	@Router			/events/{id} [delete]
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteEvent(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ReviewEvent godoc
//
//	@Summary		Revisar un evento
//	@Description	Marca un evento como revisado y asigna automáticamente un estado de gestión según su tipo
//	@Tags			events
//	@Produce		json
//	@Param			id	path		string	true	"ID del evento"
//	@Success		200	{object}	models.EventResponse
//	@Failure		404	{object}	models.ErrorResponse	"Evento no encontrado"
//	@Failure		500	{object}	models.ErrorResponse	"Error interno del servidor"
//	@Router			/events/{id}/review [put]
func (h *EventHandler) ReviewEvent(c *gin.Context) {
	id := c.Param("id")
	var req models.ReviewEventRequest

	event, err := h.service.ReviewEvent(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

// UnreviewEvent godoc
//
//	@Summary		Deshacer revisión de un evento
//	@Description	Devuelve un evento del estado revisado al estado pendiente
//	@Tags			events
//	@Produce		json
//	@Param			id	path		string	true	"ID del evento"
//	@Success		200	{object}	models.EventResponse
//	@Failure		400	{object}	models.ErrorResponse	"El evento no está en estado revisado"
//	@Failure		404	{object}	models.ErrorResponse	"Evento no encontrado"
//	@Failure		500	{object}	models.ErrorResponse	"Error interno del servidor"
//	@Router			/events/{id}/unreview [put]
func (h *EventHandler) UnreviewEvent(c *gin.Context) {
	id := c.Param("id")

	event, err := h.service.UnreviewEvent(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

// GetEventTypes godoc
//
//	@Summary		Obtener tipos de eventos
//	@Description	Obtiene una lista de los tipos de eventos disponibles
//	@Tags			events
//	@Produce		json
//	@Success		200	{array}	string
//	@Router			/events/types [get]
func (h *EventHandler) GetEventTypes(c *gin.Context) {
	types := h.service.GetEventTypes(c.Request.Context())
	c.JSON(http.StatusOK, types)
}

// GetEventStatus godoc
//
//	@Summary		Obtener estados de eventos
//	@Description	Obtiene una lista de los estados de eventos disponibles
//	@Tags			events
//	@Produce		json
//	@Success		200	{array}	string
//	@Router			/events/status [get]
func (h *EventHandler) GetEventStatus(c *gin.Context) {
	status := h.service.GetEventStatus(c.Request.Context())
	c.JSON(http.StatusOK, status)
}

// GetEventManagementStatus godoc
//
//	@Summary		Obtener estados de gestión de eventos
//	@Description	Obtiene una lista de los estados de gestión de eventos disponibles
//	@Tags			events
//	@Produce		json
//	@Success		200	{array}	string
//	@Router			/events/management-status [get]
func (h *EventHandler) GetEventManagementStatus(c *gin.Context) {
	managementStatus := h.service.GetEventManagementStatus(c.Request.Context())
	c.JSON(http.StatusOK, managementStatus)
}

// SeedEvents godoc
//
//	@Summary		Generar eventos de ejemplo
//	@Description	Genera eventos de ejemplo para pruebas
//	@Tags			events
//	@Produce		json
//	@Success		201	{object}	models.SuccessResponse	"Eventos generados correctamente"
//	@Failure		500	{object}	models.ErrorResponse	"Error interno del servidor"
//	@Router			/events/seed [post]
func (h *EventHandler) SeedEvents(c *gin.Context) {
	err := h.service.SeedEvents(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "eventos de ejemplo generados con éxito"})
}

// GetEventsRequiringManagement godoc
//
//	@Summary		Obtener eventos que requieren gestión
//	@Description	Obtiene una lista de eventos revisados que requieren gestión
//	@Tags			events
//	@Produce		json
//	@Success		200	{array}		models.EventResponse
//	@Failure		404	{object}	models.ErrorResponse	"No se encontraron eventos"
//	@Failure		500	{object}	models.ErrorResponse	"Error interno del servidor"
//	@Router			/events/management-required [get]
func (h *EventHandler) GetEventsRequiringManagement(c *gin.Context) {
	events, err := h.service.GetEventsRequiringManagement(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Si no se encuentran eventos, retorna un 404
	if len(events) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no se encontraron eventos"})
		return
	}

	c.JSON(http.StatusOK, events)
}

// GetEventsNotRequiringManagement godoc
//
//	@Summary		Obtener eventos que no requieren gestión
//	@Description	Obtiene una lista de eventos revisados que no requieren gestión
//	@Tags			events
//	@Produce		json
//	@Success		200	{array}		models.EventResponse
//	@Failure		404	{object}	models.ErrorResponse	"No se encontraron eventos"
//	@Failure		500	{object}	models.ErrorResponse	"Error interno del servidor"
//	@Router			/events/no-management-required [get]
func (h *EventHandler) GetEventsNotRequiringManagement(c *gin.Context) {
	events, err := h.service.GetEventsNotRequiringManagement(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Si no se encuentran eventos, retorna un 404
	if len(events) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no se encontraron eventos"})
		return
	}

	c.JSON(http.StatusOK, events)
}
