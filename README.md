# API de Gestión de Eventos

Este proyecto implementa una API RESTful para la gestión de eventos utilizando Go, Gin, Fx, MongoDB y Docker.

## Requisitos previos

- Docker y Docker Compose instalados en tu sistema


## Instalación y ejecución

1. Clona este repositorio:


```shellscript
git clone https://github.com/LifeRIP/events-api.git
cd events-api
```

2. Inicia la aplicación con Docker Compose:


```shellscript
docker compose up
```

Esto construirá e iniciará todos los servicios necesarios (API y MongoDB).

## Documentación y uso

Una vez que la aplicación esté en ejecución, puedes acceder a la documentación Swagger en:

```plaintext
http://localhost:8080/swagger/index.html
```

### Carga de datos de prueba

Para cargar datos de ejemplo en la base de datos:

1. Navega a [http://localhost:8080/swagger/index.html#/events/post_events_seed](http://localhost:8080/swagger/index.html#/events/post_events_seed)
2. Haz clic en el botón "Try it out"
3. Haz clic en "Execute" para enviar la petición
4. Verifica que la respuesta sea exitosa (código 201)


Después de cargar los datos de prueba, puedes utilizar todos los endpoints disponibles en la API.

## Endpoints disponibles

La API proporciona los siguientes endpoints principales:

- **GET /api/v1/events**: Obtener todos los eventos
- **POST /api/v1/events**: Crear un nuevo evento
- **GET /api/v1/events/id**: Obtener un evento por ID
- **PUT /api/v1/events/id**: Actualizar un evento
- **DELETE /api/v1/events/id**: Eliminar un evento
- **PUT /api/v1/events/id/review**: Revisar un evento
- **PUT /api/v1/events/id/unreview**: Deshacer revisión de un evento
- **GET /api/v1/events/types**: Obtener tipos de eventos
- **GET /api/v1/events/status**: Obtener estados de eventos
- **GET /api/v1/events/management-status**: Obtener estados de gestión
- **GET /api/v1/events/management-required**: Obtener eventos que requieren gestión
- **GET /api/v1/events/no-management-required**: Obtener eventos que no requieren gestión


La documentación completa de todos los endpoints, parámetros y respuestas está disponible en la interfaz Swagger.

## Estructura del proyecto

```plaintext
/events-api
  /cmd
    /api
      main.go
  /internal
    /config
    /models
    /repositories
    /services
    /handlers
    /middleware
  /pkg
    /database
  /docs
  /scripts
  Dockerfile
  docker-compose.yml
```

## Tecnologías utilizadas

- **Go**: Lenguaje de programación
- **Gin**: Framework web
- **Fx**: Inyección de dependencias
- **MongoDB**: Base de datos NoSQL
- **Swagger**: Documentación de la API
- **Docker**: Contenedorización


## Características principales

- CRUD completo de eventos
- Clasificación de eventos (requiere gestión / sin gestión)
- Documentación interactiva con Swagger
- Arquitectura modular y escalable
- Completamente dockerizado