package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tov_tools/pkg/character"
	"tov_tools/pkg/static_data"
)

const (
	ErrMissingParameters = "Missing required parameters"
	ErrUnsupportedType   = "Unsupported type"
)

func GetTable(c *gin.Context) {
	dataType := c.Query("type")

	if dataType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrMissingParameters})
		return
	}

	table, err := getTableDataByType(dataType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, table)

}

// getTableDataByType delegates the table generation based on type.
func getTableDataByType(dataType string) (interface{}, error) {
	switch dataType {
	case "class":
		return character.Classes, nil
	case "damageModifier":
		return static_data.DamageModifiers(), nil
	case "damageType":
		return static_data.DamageType(), nil
	default:
		return nil, fmt.Errorf("%s: '%s'", ErrUnsupportedType, dataType)
	}
}
