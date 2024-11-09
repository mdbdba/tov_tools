package logging

import (
	"encoding/json"
	"go.uber.org/zap"
	"time"
)

var (
	logger *zap.Logger
)

func init() {
	logger, _ = zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)
}

// LogData represents the structure for logging function call details
type LogData struct {
	Timestamp  string                 `json:"timestamp"`
	Canonical  bool                   `json:"canonical"`
	UnitOfWork string                 `json:"unitOfWork"`
	RequestID  string                 `json:"requestID"`
	DurationMs int64                  `json:"durationMs"`
	Message    map[string]interface{} `json:"message"`
}

// PopulateLogMessage populates LogData's Message field with a struct and message string
func PopulateLogMessage(structData interface{}, message string) map[string]interface{} {
	logMessage := make(map[string]interface{})
	logMessage["text"] = message

	dataBytes, err := json.Marshal(structData)
	if err != nil {
		logger.Error("Failed to marshal structData to JSON", zap.Error(err))
		logMessage["data"] = "Failed to marshal structData"
	} else {
		var dataMap map[string]interface{}
		err := json.Unmarshal(dataBytes, &dataMap)
		if err != nil {
			logger.Error("Failed to unmarshal structData to map", zap.Error(err))
			logMessage["data"] = "Failed to unmarshal structData"
		} else {
			logMessage["data"] = dataMap
		}
	}
	return logMessage
}

// New creates and populates a new LogData struct
func New(unitOfWork string) *LogData {
	startTime := time.Now()

	logData := &LogData{
		Timestamp:  startTime.Format(time.RFC3339),
		Canonical:  true,
		UnitOfWork: unitOfWork,
	}
	return logData
}

// LogUnitOfWork finalizes and logs the LogData
func LogUnitOfWork(logData *LogData, structData interface{}, message string) {
	startTime, err := time.Parse(time.RFC3339, logData.Timestamp)
	if err != nil {
		logger.Error("Failed to parse timestamp", zap.Error(err))
		return
	}
	logData.DurationMs = time.Since(startTime).Milliseconds()
	logData.Message = PopulateLogMessage(structData, message)
	logData.Timestamp = time.Now().Format(time.RFC3339)

	logger.Info("Perform operation completed",
		zap.String("timestamp", logData.Timestamp),
		zap.Bool("canonical", logData.Canonical),
		zap.String("unitOfWork", logData.UnitOfWork),
		zap.Int64("durationMs", logData.DurationMs),
		zap.Any("message", logData.Message),
	)
}
