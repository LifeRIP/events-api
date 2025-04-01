package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Logger es un middleware para registrar información de solicitudes HTTP
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Tiempo de inicio
		startTime := time.Now()

		// Procesar solicitud
		c.Next()

		// Tiempo de finalización
		endTime := time.Now()
		// Tiempo de ejecución
		latencyTime := endTime.Sub(startTime)

		// Obtener detalles de la solicitud
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		// Registrar detalles de la solicitud
		if statusCode >= 400 {
			// Registrar errores
			gin.DefaultErrorWriter.Write([]byte(
				"[ERROR] " + endTime.Format("2006-01-02 15:04:05") +
					" | " + latencyTime.String() +
					" | " + clientIP +
					" | " + reqMethod +
					" | " + reqURI +
					" | " + string(rune(statusCode)) + "\n",
			))
		} else {
			// Registrar solicitudes exitosas
			gin.DefaultWriter.Write([]byte(
				"[INFO] " + endTime.Format("2006-01-02 15:04:05") +
					" | " + latencyTime.String() +
					" | " + clientIP +
					" | " + reqMethod +
					" | " + reqURI +
					" | " + string(rune(statusCode)) + "\n",
			))
		}
	}
}
