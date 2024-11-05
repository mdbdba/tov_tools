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

type SampleStruct struct {
	Field1 string
	Field2 int
}

func TestPopulateLogMessage(t *testing.T) {
	assertions := assert.New(t)

	// Create a sample struct to use in the test
	sample := SampleStruct{
		Field1: "test-value",
		Field2: 123,
	}

	// Call PopulateLogMessage with the sample struct and a message
	logMessage := PopulateLogMessage(sample, "Test Message")

	// Convert logMessage to JSON string to inspect the contents
	jsonString, err := json.Marshal(logMessage)
	assertions.NoError(err)

	// Validate JSON structure
	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &jsonMap)
	assertions.NoError(err)
	assertions.Equal("Test Message", jsonMap["text"])
	assertions.Equal("test-value", jsonMap["data"].(map[string]interface{})["Field1"])
	assertions.Equal(123, int(jsonMap["data"].(map[string]interface{})["Field2"].(float64))) // JSON unmarshals numbers as float64
}

func TestNew(t *testing.T) {
	assertions := assert.New(t)

	// Call New to create a LogData instance
	logData := New("UnitTest")

	// Check the returned LogData
	assertions.NotNil(logData)
	assertions.Equal("UnitTest", logData.UnitOfWork)
	assertions.True(logData.Canonical)
	assertions.Equal(13, len(logData.RequestID)) // Ensure that request ID has the correct length

	// Validate timestamp format
	_, err := time.Parse(time.RFC3339, logData.Timestamp)
	assertions.NoError(err)
}

func TestLogUnitOfWork(t *testing.T) {
	assertions := assert.New(t)

	// Prepare log observer
	core, observed := observer.New(zapcore.InfoLevel)
	logger = zap.New(core)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
		}
	}(logger)

	// Create a sample struct to use in the test
	sample := SampleStruct{
		Field1: "test-value",
		Field2: 123,
	}

	// Create a new LogData instance
	logData := New("UnitTest")

	// Call LogUnitOfWork with the sample struct and a message
	LogUnitOfWork(logData, sample, "Test Message")

	// Check the logged entry
	logEntries := observed.All()
	assertions.Len(logEntries, 1) // Ensure one log entry

	// Convert log entry to JSON string to inspect the contents
	var loggedMap map[string]interface{}
	err := json.Unmarshal([]byte(logEntries[0].Message), &loggedMap)
	assertions.NoError(err)

	// Validate log entry contents
	assertions.Equal("UnitTest", loggedMap["unitOfWork"])
	assertions.Equal("Test Message", loggedMap["message"].(map[string]interface{})["text"])
	assertions.Equal("test-value", loggedMap["message"].(map[string]interface{})["data"].(map[string]interface{})["Field1"])
	assertions.Equal(123, int(loggedMap["message"].(map[string]interface{})["data"].(map[string]interface{})["Field2"].(float64)))
}

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
