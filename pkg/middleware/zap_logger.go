package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		respBody := bytes.NewBufferString("")
		responseWriter := &responseBodyWriter{body: respBody, ResponseWriter: c.Writer}
		c.Writer = responseWriter

		c.Next()

		duration := time.Since(start)
		statusCode := responseWriter.Status()

		var responseData interface{}
		if respBody.Len() > 0 {
			if err := json.Unmarshal(respBody.Bytes(), &responseData); err != nil {
				zap.L().Error("Failed to unmarshal response body", zap.Error(err))
			}
		}

		zap.L().Info("Request",
			zap.Int("status", statusCode),
			zap.Duration("duration", duration),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("uri", c.Request.RequestURI),
			zap.Any("response", responseData),
		)
	}
}
