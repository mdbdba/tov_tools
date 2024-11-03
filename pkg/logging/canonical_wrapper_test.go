package logging

import (
	"encoding/json"
	"go.uber.org/zap"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLogData_ToString(t *testing.T) {
	assertions := assert.New(t)

	// Prepare log data
	logData := &LogData{
		Timestamp:  time.Now().Format(time.RFC3339),
		Canonical:  true,
		UnitOfWork: "UnitTest",
		RequestID:  "test-request-id",
		DurationMs: 123,
		Message: map[string]interface{}{
			"key1": "value1",
			"key2": 2,
		},
	}

	// Convert log data to JSON string
	jsonString := logData.ToString()
	assertions.NotEmpty(jsonString)

	// Validate JSON structure
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	assertions.NoError(err)
	assertions.Equal("value1", jsonMap["message"].(map[string]interface{})["key1"])
	assertions.Equal(2, int(jsonMap["message"].(map[string]interface{})["key2"].(float64))) // JSON unmarshals numbers as float64
}

func TestLogFunctionCall(t *testing.T) {
	assertions := assert.New(t)

	// Prepare log observer
	core, observed := observer.New(zapcore.InfoLevel)
	logger = zap.New(core)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	// Prepare log data
	logData := &LogData{
		Timestamp:  time.Now().Format(time.RFC3339),
		Canonical:  true,
		UnitOfWork: "UnitTest",
		RequestID:  "test-request-id",
		DurationMs: 123,
		Message: map[string]interface{}{
			"key1": "value1",
			"key2": 2,
		},
	}

	// Log the function call
	LogFunctionCall(logData)

	// Check the logged entry
	logEntries := observed.All()
	assertions.Len(logEntries, 1)                               // Ensure one log entry
	assertions.Equal(logData.ToString(), logEntries[0].Message) // Ensure the correct log message
}
