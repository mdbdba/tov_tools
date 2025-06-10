package logging

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
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
	assertions.Equal(float64(123), jsonMap["data"].(map[string]interface{})["Field2"].(float64)) // JSON unmarshals numbers as float64
}

func TestNew(t *testing.T) {
	assertions := assert.New(t)

	// Call New to create a LogData instance
	logData := New("UnitTest")

	// Check the returned LogData
	assertions.NotNil(logData)
	assertions.Equal("UnitTest", logData.UnitOfWork)
	assertions.True(logData.Canonical)

	// No need to check RequestID length as it's now generated differently
	assertions.NotEmpty(logData.RequestID) // Just ensure it's not empty

	// Validate timestamp format
	_, err := time.Parse(time.RFC3339, logData.Timestamp)
	assertions.NoError(err)
}

func TestLogUnitOfWork(t *testing.T) {
	assertions := assert.New(t)

	// Prepare log observer
	core, observed := observer.New(zapcore.InfoLevel)
	originalLogger := logger
	logger = zap.New(core)
	defer func() {
		logger = originalLogger // Restore original logger
	}()

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

	// Access fields directly from the entry's Context instead of parsing the message
	// as our logger doesn't serialize the entire log object into the message field
	unitOfWork := findField(logEntries[0].Context, "unitOfWork")
	assertions.NotNil(unitOfWork)
	assertions.Equal("UnitTest", unitOfWork.String)

	message := findField(logEntries[0].Context, "message")
	assertions.NotNil(message)

	// Convert the message field (which is an interface{}) to a map
	messageMap, ok := message.Interface.(map[string]interface{})
	assertions.True(ok, "message field should be a map")

	assertions.Equal("Test Message", messageMap["text"])
	dataMap, ok := messageMap["data"].(map[string]interface{})
	assertions.True(ok, "data field should be a map")
	assertions.Equal("test-value", dataMap["Field1"])
	assertions.Equal(float64(123), dataMap["Field2"].(float64))
}

// Helper function to find a field by key in zap field slice
func findField(fields []zapcore.Field, key string) *zapcore.Field {
	for i, field := range fields {
		if field.Key == key {
			return &fields[i]
		}
	}
	return nil
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

	// Access message field and validate its contents
	message, ok := jsonMap["message"].(map[string]interface{})
	assertions.True(ok, "message should be a map")
	assertions.Equal("value1", message["key1"])
	assertions.Equal(float64(2), message["key2"].(float64)) // JSON unmarshals numbers as float64
}
