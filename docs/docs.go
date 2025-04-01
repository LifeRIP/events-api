// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/events": {
            "get": {
                "description": "Obtiene una lista de todos los eventos",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Obtener todos los eventos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.EventResponse"
                            }
                        }
                    },
                    "404": {
                        "description": "No se encontraron eventos",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Crea un nuevo evento con la información proporcionada",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Crear un nuevo evento",
                "parameters": [
                    {
                        "description": "Información del evento",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateEventRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.EventResponse"
                        }
                    },
                    "400": {
                        "description": "Error en los datos de entrada",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/events/management-required": {
            "get": {
                "description": "Obtiene una lista de eventos revisados que requieren gestión",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Obtener eventos que requieren gestión",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.EventResponse"
                            }
                        }
                    },
                    "404": {
                        "description": "No se encontraron eventos",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/events/management-status": {
            "get": {
                "description": "Obtiene una lista de los estados de gestión de eventos disponibles",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Obtener estados de gestión de eventos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/events/no-management-required": {
            "get": {
                "description": "Obtiene una lista de eventos revisados que no requieren gestión",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Obtener eventos que no requieren gestión",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.EventResponse"
                            }
                        }
                    },
                    "404": {
                        "description": "No se encontraron eventos",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/events/seed": {
            "post": {
                "description": "Genera eventos de ejemplo para pruebas",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Generar eventos de ejemplo",
                "responses": {
                    "201": {
                        "description": "Eventos generados correctamente",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/events/status": {
            "get": {
                "description": "Obtiene una lista de los estados de eventos disponibles",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Obtener estados de eventos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/events/types": {
            "get": {
                "description": "Obtiene una lista de los tipos de eventos disponibles",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Obtener tipos de eventos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/events/{id}": {
            "get": {
                "description": "Obtiene un evento por su ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Obtener un evento por ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del evento",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EventResponse"
                        }
                    },
                    "404": {
                        "description": "Evento no encontrado",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Actualiza un evento existente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Actualizar un evento",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del evento",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Información del evento",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateEventRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EventResponse"
                        }
                    },
                    "400": {
                        "description": "Error en los datos de entrada",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Evento no encontrado",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Elimina un evento existente",
                "tags": [
                    "events"
                ],
                "summary": "Eliminar un evento",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del evento",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Evento no encontrado",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/events/{id}/review": {
            "put": {
                "description": "Marca un evento como revisado y asigna automáticamente un estado de gestión según su tipo",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Revisar un evento",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del evento",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EventResponse"
                        }
                    },
                    "404": {
                        "description": "Evento no encontrado",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/events/{id}/unreview": {
            "put": {
                "description": "Devuelve un evento del estado revisado al estado pendiente",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Deshacer revisión de un evento",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del evento",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EventResponse"
                        }
                    },
                    "400": {
                        "description": "El evento no está en estado revisado",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Evento no encontrado",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateEventRequest": {
            "type": "object",
            "required": [
                "date",
                "description",
                "name",
                "type"
            ],
            "properties": {
                "date": {
                    "type": "string",
                    "example": "2025-04-08T00:00:00Z"
                },
                "description": {
                    "type": "string",
                    "example": "Mantenimiento programado del sistema para actualización"
                },
                "name": {
                    "type": "string",
                    "example": "Mantenimiento programado"
                },
                "type": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.EventType"
                        }
                    ],
                    "example": "MAINTENANCE"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "mensaje descriptivo del error"
                }
            }
        },
        "models.EventResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "managementStatus": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.EventType": {
            "type": "string",
            "enum": [
                "EMERGENCY",
                "MAINTENANCE",
                "NOTIFICATION",
                "ALERT",
                "INFO"
            ],
            "x-enum-varnames": [
                "TypeEmergency",
                "TypeMaintenance",
                "TypeNotification",
                "TypeAlert",
                "TypeInfo"
            ]
        },
        "models.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "operación realizada con éxito"
                }
            }
        },
        "models.UpdateEventRequest": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/models.EventType"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Events API",
	Description:      "API para gestión de eventos",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
