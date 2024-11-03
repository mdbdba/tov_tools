package logging

import (
	"encoding/json"
	"go.uber.org/zap"
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

// ToString converts the LogData struct to a JSON string
func (logData *LogData) ToString() string {
	jsonData, err := json.Marshal(logData)
	if err != nil {
		logger.Error("Failed to marshal log data to JSON", zap.Error(err))
		return ""
	}
	return string(jsonData)
}

// LogFunctionCall logs the details of a function call
func LogFunctionCall(logData *LogData) {
	logger.Info(logData.ToString())
}
